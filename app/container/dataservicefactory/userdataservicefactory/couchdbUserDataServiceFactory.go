package userdataservicefactory

import (
	"github.com/go-kivik/kivik"
	"github.com/jfeng45/servicetmpl1/app/config"
	"github.com/jfeng45/servicetmpl1/app/container"
	"github.com/jfeng45/servicetmpl1/app/container/datastorefactory"
	"github.com/jfeng45/servicetmpl1/app/logger"
	"github.com/jfeng45/servicetmpl1/applicationservice/dataservice"
	"github.com/jfeng45/servicetmpl1/applicationservice/dataservice/userdata/couchdb"
	"github.com/pkg/errors"
)

// couchdbUserDataServiceFactory is a empty receiver for Build method
type couchdbUserDataServiceFactory struct{}

func (cudsf *couchdbUserDataServiceFactory) Build(c container.Container, dataConfig *config.DataConfig) (dataservice.UserDataInterface, error) {
	logger.Log.Debug("couchdbUserDataServiceFactory")
	dsc := dataConfig.DataStoreConfig
	dsi, err := datastorefactory.GetDataStoreFb(dsc.Code).Build(c, &dsc)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	ds := dsi.(*kivik.DB)
	udc := couchdb.UserDataCouchdb{DB: ds}
	return &udc, nil

}
