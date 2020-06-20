package registration

import (
	"github.com/jfeng45/servicetmpl1/applicationservice/dataservice"
	"github.com/jfeng45/servicetmpl1/domain/model"
)

// RegistrationTxUseCase implements RegistrationTxUseCaseInterface.
// It has UserDataInterface, which can be used to access persistence layer
type RegistrationTxUseCase struct {
	UserDataInterface dataservice.UserDataInterface
}

// The use case of ModifyAndUnregister with transaction
func (rtuc *RegistrationTxUseCase) ModifyAndUnregisterWithTx(user *model.User) error {

	udi := rtuc.UserDataInterface
	return udi.EnableTx(func() error {
		// wrap the business function inside the TxEnd function
		return ModifyAndUnregister(udi, user)
	})
}
