package dataservicefactory

import (
	"github.com/jfeng45/servicetmpl1/app/config"
	"github.com/jfeng45/servicetmpl1/app/container"
	"github.com/jfeng45/servicetmpl1/app/container/datastorefactory"
	"github.com/jfeng45/servicetmpl1/app/logger"
	"github.com/jfeng45/servicetmpl1/applicationservice/cacheclient"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// cacheDataServiceFactory is a empty receiver for Build method
type cacheDataServiceFactory struct{}

func (cdsf *cacheDataServiceFactory) Build(c container.Container, dataConfig *config.DataConfig) (DataServiceInterface, error) {
	logger.Log.Debug("cacheDataServiceFactory")
	dsc := dataConfig.DataStoreConfig
	dsi, err := datastorefactory.GetDataStoreFb(dsc.Code).Build(c, &dsc)
	grpcConn := dsi.(*grpc.ClientConn)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	cdg := cacheclient.CacheDataGrpc{grpcConn}
	//logger.Log.Debug("udm:", udm.DB)

	return &cdg, nil
}
