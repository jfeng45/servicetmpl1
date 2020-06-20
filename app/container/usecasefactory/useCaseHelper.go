package usecasefactory

import (
	"github.com/jfeng45/servicetmpl1/app/config"
	"github.com/jfeng45/servicetmpl1/app/container"
	"github.com/jfeng45/servicetmpl1/app/container/dataservicefactory"
	"github.com/jfeng45/servicetmpl1/applicationservice/dataservice"
	"github.com/pkg/errors"
)

func buildUserData(c container.Container, dc *config.DataConfig) (dataservice.UserDataInterface, error) {
	dsi, err := dataservicefactory.GetDataServiceFb(dc.Code).Build(c, dc)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	udi := dsi.(dataservice.UserDataInterface)
	return udi, nil
}

func buildCacheData(c container.Container, dc *config.DataConfig) (dataservice.CacheDataInterface, error) {
	//logger.Log.Debug("uc:", cdc)
	dsi, err := dataservicefactory.GetDataServiceFb(dc.Code).Build(c, dc)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	cdi := dsi.(dataservice.CacheDataInterface)
	return cdi, nil
}

