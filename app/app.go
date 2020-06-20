package app

import (
	logConfig "github.com/jfeng45/glogger/config"
	logFactory "github.com/jfeng45/glogger/factory"
	"github.com/jfeng45/servicetmpl1/app/config"
	"github.com/jfeng45/servicetmpl1/app/container"
	"github.com/jfeng45/servicetmpl1/app/container/servicecontainer"
	"github.com/jfeng45/servicetmpl1/app/logger"
	"github.com/pkg/errors"
)

// InitApp loads the application configurations from a file and saved it in appConfig and initialize the logger
// The appConfig is cached in container, so it only loads the configuration file once.
// InitApp only needs to be called once. If the configuration changes, you can call it again to reinitialize the app.
func InitApp(filename...string) (container.Container, error) {
	config, err := config.BuildConfig(filename...)
	if err != nil {
		return nil, errors.Wrap(err, "BuildConfig")
	}
	err = initLogger(&config.LogConfig)
	if err != nil {
		return nil, err
	}
	return initContainer(config)
}

func initLogger (lc *logConfig.Logging) error{
	log, err := logFactory.Build(lc)
	if err != nil {
		return errors.Wrap(err, "loadLogger")
	}
	logger.SetLogger(log)
	return nil
}

func initContainer(config *config.AppConfig) (container.Container, error) {
	factoryMap := make(map[string]interface{})
	c := servicecontainer.ServiceContainer{factoryMap,config}
	//gdbc, err :=initGdbc(&c.AppConfig.SQLConfig)
	//if err != nil {
	//	return nil,err
	//}
	//key := config.SQLConfig.Code
	//c.Put(key, gdbc)
	return &c, nil
}
