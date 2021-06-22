package registeruser

import (
	"context"
)

// Inport of RegisterUser
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase RegisterUser
type InportRequest struct {
	Email    string `` //
	Password string `` //
	Address  string `` //
}

// InportResponse is response payload after running the usecase RegisterUser
type InportResponse struct {
}
