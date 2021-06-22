package showalluser

import "userprofile/domain/repository"

// Outport of ShowAllUSer
type Outport interface {
	repository.FindAllUserRepo
}
