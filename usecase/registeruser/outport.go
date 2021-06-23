package registeruser

import (
	"userprofile/domain/repository"
	"userprofile/domain/service"
)

// Outport of RegisterUser
type Outport interface {
	repository.SaveUserRepo
	repository.FindOneUserByEmailRepo
	service.GenerateUUIDService
	service.SendEmailService
	service.ConstructStartActivationMessageService
	service.HashPasswordService
	service.GenerateRandomStringService
}
