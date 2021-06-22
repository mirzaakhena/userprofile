package updateuser

import (
	"context"
)

// Inport of UpdateUser
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase UpdateUser
type InportRequest struct {
	UserID  string
	Address string
}

// InportResponse is response payload after running the usecase UpdateUser
type InportResponse struct {
}
