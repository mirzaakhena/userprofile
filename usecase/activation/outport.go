package activation

import (
	"userprofile/domain/repository"
	"userprofile/domain/service"
)

// Outport of Activation
type Outport interface {
	repository.TransactionDB
	repository.FindOneUserByEmailRepo
	repository.SaveUserRepo
	service.SendEmailService
	service.ConstructSuccessActivationMessageService
}
