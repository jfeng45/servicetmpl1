package userdataservicefactory

import (
	"github.com/jfeng45/gtransaction/gdbc"
	"github.com/jfeng45/servicetmpl1/app/config"
	"github.com/jfeng45/servicetmpl1/app/container"
	"github.com/jfeng45/servicetmpl1/app/container/datastorefactory"
	"github.com/jfeng45/servicetmpl1/app/logger"
	"github.com/jfeng45/servicetmpl1/applicationservice/dataservice"
	"github.com/jfeng45/servicetmpl1/applicationservice/dataservice/userdata/sqldb"
	"github.com/pkg/errors"
)

// sqlUserDataServiceFactory is a empty receiver for Build method
type sqlUserDataServiceFactory struct{}

func (sudsf *sqlUserDataServiceFactory) Build(c container.Container, dataConfig *config.DataConfig) (dataservice.UserDataInterface, error) {
	logger.Log.Debug("sqlUserDataServiceFactory")
	dsc := dataConfig.DataStoreConfig
	dsi, err := datastorefactory.GetDataStoreFb(dsc.Code).Build(c, &dsc)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	ds := dsi.(gdbc.SqlGdbc)
	uds := sqldb.UserDataSql{DB: ds}
	logger.Log.Debug("uds:", uds.DB)
	return &uds, nil

}
