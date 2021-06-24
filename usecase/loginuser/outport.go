package loginuser

import (
	"userprofile/domain/repository"
	"userprofile/domain/service"
)

// Outport of LoginUser
type Outport interface {
	repository.ReadOnlyDB
	repository.FindOneUserByEmailRepo
	service.GenerateTokenService
	service.ValidatePasswordService
}
