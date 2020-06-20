package config

import (
	"github.com/pkg/errors"
)

// database code. Need to map to the database code (DataStoreConfig) in the configuration yaml file.
const (
	SQLDB      string = "sqldb"
	COUCHDB    string = "couch"
	CACHE_GRPC string = "cacheGrpc"
	USER_GRPC  string = "userGrpc"
)

// constant for logger code, it needs to match log code (logConfig)in configuration
const (
	LOGRUS string = "logrus"
	ZAP    string = "zap"
)

// use case code. Need to map to the use case code (UseCaseConfig) in the configuration yaml file.
// Client app use those to retrieve use case from the container
const (
	REGISTRATION string = "registration"
	REGISTRATION_TX string = "registrationTx"
	LIST_USER    string = "listUser"
)

// data service code. Need to map to the data service code (DataConfig) in the configuration yaml file.
const (
	USER_DATA   string = "userData"
	CACHE_DATA  string = "cacheData"
	TX_DATA     string = "txData"
	//COURSE_DATA string = "courseData"
)

func validateConfig(appConfig AppConfig) error {
	err := validateDataStore(appConfig)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateLogger(appConfig)
	if err != nil {
		return errors.Wrap(err, "")
	}
	useCase := appConfig.UseCaseConfig
	err = validateUseCase(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func validateLogger(appConfig AppConfig) error {
	zc :=  appConfig.ZapConfig
	key := zc.Code
	zcMsg := " in validateLogger doesn't match key = "
	if ZAP != key {
		errMsg := ZAP + zcMsg + key
		return errors.New(errMsg)
	}
	lc := appConfig.LorusConfig
	key = lc.Code
	if LOGRUS != lc.Code {
		errMsg := LOGRUS + zcMsg + key
		return errors.New(errMsg)
	}
	return nil
}

func validateDataStore(appConfig AppConfig) error {
	sc := appConfig.SQLConfig
	key := sc.Code
	scMsg := " in validateDataStore doesn't match key = "
	if SQLDB != key {
		errMsg := SQLDB + scMsg + key
		return errors.New(errMsg)
	}
	cc := appConfig.CouchdbConfig
	key = cc.Code
	if COUCHDB != key {
		errMsg := COUCHDB + scMsg + key
		return errors.New(errMsg)
	}
	cgc := appConfig.CacheGrpcConfig
	key = cgc.Code
	if CACHE_GRPC != key {
		errMsg := CACHE_GRPC + scMsg + key
		return errors.New(errMsg)
	}

	ugc := appConfig.UserGrpcConfig
	key = ugc.Code
	if USER_GRPC != key {
		errMsg := USER_GRPC + scMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateUseCase(useCase UseCaseConfig) error {
	err := validateRegistration(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateRegistrationTx(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateListUser(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func validateRegistration(useCaseConfig UseCaseConfig) error {
	rc := useCaseConfig.Registration
	key := rc.Code
	rcMsg := " in validateRegistration doesn't match key = "
	if REGISTRATION != key {
		errMsg := REGISTRATION + rcMsg + key
		return errors.New(errMsg)
	}
	key = rc.UserDataConfig.Code
	if USER_DATA != key {
		errMsg := USER_DATA + rcMsg + key
		return errors.New(errMsg)
	}
	return nil
}

func validateRegistrationTx(useCaseConfig UseCaseConfig) error {
	rc := useCaseConfig.RegistrationTx
	key := rc.Code
	rcMsg := " in validateRegistrationTx doesn't match key = "
	if REGISTRATION_TX != key {
		errMsg := REGISTRATION_TX + rcMsg + key
		return errors.New(errMsg)
	}
	key = rc.UserDataConfig.Code
	if USER_DATA != key {
		errMsg := USER_DATA + rcMsg + key
		return errors.New(errMsg)
	}
	return nil
}

func validateListUser(useCaseConfig UseCaseConfig) error {
	lc := useCaseConfig.ListUser
	key := lc.Code
	luMsg := " in validateListUser doesn't match key = "
	if LIST_USER != key {
		errMsg := LIST_USER + luMsg + key
		return errors.New(errMsg)
	}
	key = lc.CacheDataConfig.Code
	if CACHE_DATA != key {
		errMsg := CACHE_DATA + luMsg + key
		return errors.New(errMsg)
	}
	return nil
}

