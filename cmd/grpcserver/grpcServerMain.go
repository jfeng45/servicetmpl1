package main

import (
	"context"
	"fmt"
	"github.com/jfeng45/servicetmpl1/app"
	"github.com/jfeng45/servicetmpl1/app/container/containerhelper"
	"github.com/jfeng45/servicetmpl1/app/container/servicecontainer"
	"github.com/jfeng45/servicetmpl1/app/logger"
	"github.com/jfeng45/servicetmpl1/applicationservice/userclient"
	uspb "github.com/jfeng45/servicetmpl1/applicationservice/userclient/generatedclient"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"net"
)

const (
	DEV_CONFIG  string = "../../../app/config/appConfigDev.yaml"
	PROD_CONFIG string = "../../../app/config/appConfigProd.yaml"
)

type UserService struct {
	//container container.Container
	container *servicecontainer.ServiceContainer
}

func catchPanic() {
	if p := recover(); p != nil {
		logger.Log.Errorf("%+v\n", p)
	}
}

func (uss *UserService) RegisterUser(ctx context.Context, req *uspb.RegisterUserReq) (*uspb.RegisterUserResp, error) {
	defer catchPanic()
	logger.Log.Debug("RegisterUser called")

	ruci, err := containerhelper.GetRegistrationUseCase(uss.container)
	if err != nil {
		logger.Log.Errorf("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}
	mu, err := userclient.GrpcToUser(req.User)

	if err != nil {
		logger.Log.Errorf("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}
	logger.Log.Debug("mu:", mu)
	resultUser, err := ruci.RegisterUser(mu)
	if err != nil {
		logger.Log.Errorf("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}
	logger.Log.Debug("resultUser:", resultUser)
	gu, err := userclient.UserToGrpc(resultUser)
	if err != nil {
		logger.Log.Errorf("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}

	logger.Log.Debug("user registered: ", gu)

	return &uspb.RegisterUserResp{User: gu}, nil

}

func (uss *UserService) ListUser(ctx context.Context, in *uspb.ListUserReq) (*uspb.ListUserResp, error) {
	defer catchPanic()
	logger.Log.Debug("ListUser called")

	luci, err := containerhelper.GetListUserUseCase(uss.container)
	if err != nil {
		logger.Log.Errorf("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}

	lu, err := luci.ListUser()
	if err != nil {
		logger.Log.Errorf("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}
	gu, err := userclient.UserListToGrpc(lu)
	if err != nil {
		logger.Log.Errorf("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}

	logger.Log.Debug("user list: ", gu)

	return &uspb.ListUserResp{User: gu}, nil

}
func runServer(sc *servicecontainer.ServiceContainer) error {
	logger.Log.Debug("start runserver")

	srv := grpc.NewServer()

	cs := &UserService{sc}
	uspb.RegisterUserServiceServer(srv, cs)
	//l, err:=net.Listen(GRPC_NETWORK, GRPC_ADDRESS)
	ugc := sc.AppConfig.UserGrpcConfig
	logger.Log.Debugf("userGrpcConfig: %+v\n", ugc)
	l, err := net.Listen(ugc.DriverName, ugc.UrlAddress)
	if err != nil {
		return errors.Wrap(err, "")
	} else {
		logger.Log.Debug("server listening")
	}
	return srv.Serve(l)
}

func main() {
	filename := DEV_CONFIG
	//filename := PROD_CONFIG
	container, err := buildContainer(filename)
	if err != nil {
		fmt.Printf("%+v\n", err)
		//logger.Log.Errorf("%+v\n", err)
		panic(err)
	}
	if err := runServer(container); err != nil {
		logger.Log.Errorf("Failed to run user server: %+v\n", err)
		panic(err)
	} else {
		logger.Log.Info("server started")
	}
}

func buildContainer(filename string) (*servicecontainer.ServiceContainer, error) {

	container, err := app.InitApp(filename)
	sc := container.(*servicecontainer.ServiceContainer)
	if err != nil {
		//logger.Log.Errorf("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}
	return sc, nil
}

