package activation

import (
	"userprofile/domain/repository"
	"userprofile/domain/service"
)

// Outport of Activation
type Outport interface {
	repository.FindOneUserByEmailRepo
	repository.SaveUserRepo
	service.SendEmailService
	service.ConstructSuccessActivationMessageService
}
