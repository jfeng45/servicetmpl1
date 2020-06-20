package containerhelper

import (
	"github.com/jfeng45/servicetmpl1/app/config"
	"github.com/jfeng45/servicetmpl1/app/container"
	"github.com/jfeng45/servicetmpl1/domain/usecase"
	"github.com/pkg/errors"
)

func GetListUserUseCase(c container.Container) (usecase.ListUserUseCaseInterface, error) {
	key := config.LIST_USER
	value, err := c.BuildUseCase(key)
	if err != nil {
		//logger.Log.Errorf("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}
	return value.(usecase.ListUserUseCaseInterface), nil
}

func GetRegistrationUseCase(c container.Container) (usecase.RegistrationUseCaseInterface, error) {
	key := config.REGISTRATION
	value, err := c.BuildUseCase(key)
	if err != nil {
		//logger.Log.Errorf("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}
	return value.(usecase.RegistrationUseCaseInterface), nil

}

func GetRegistrationTxUseCase(c container.Container) (usecase.RegistrationTxUseCaseInterface, error) {
	key := config.REGISTRATION_TX
	value, err := c.BuildUseCase(key)
	if err != nil {
		//logger.Log.Errorf("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}
	return value.(usecase.RegistrationTxUseCaseInterface), nil

}
