package usecasefactory

import (
	"github.com/jfeng45/servicetmpl1/app/config"
	"github.com/jfeng45/servicetmpl1/app/container"
	"github.com/jfeng45/servicetmpl1/domain/usecase/registration"
	"github.com/pkg/errors"
)

type RegistrationTxFactory struct {
}

// Build creates concrete type for RegistrationTxUseCaseInterface
func (rtf *RegistrationTxFactory) Build(c container.Container, appConfig *config.AppConfig, key string) (UseCaseInterface, error) {
	uc := appConfig.UseCaseConfig.RegistrationTx
	udi, err := buildUserData(c, &uc.UserDataConfig)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	ruc := registration.RegistrationTxUseCase{UserDataInterface: udi}

	return &ruc, nil
}
