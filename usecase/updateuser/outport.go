package updateuser

import "userprofile/domain/repository"

// Outport of UpdateUser
type Outport interface {
	repository.FindOneUserByIDRepo
	repository.SaveUserRepo
}
