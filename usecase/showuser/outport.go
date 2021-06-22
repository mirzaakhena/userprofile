package showuser

import "userprofile/domain/repository"

// Outport of ShowUser
type Outport interface {
	repository.FindOneUserByIDRepo
}
