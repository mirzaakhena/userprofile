package showalluser

import (
	"context"
)

// Inport of ShowAllUSer
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase ShowAllUSer
type InportRequest struct {
}

// InportResponse is response payload after running the usecase ShowAllUSer
type InportResponse struct {
	Users []UserResponse
}

type UserResponse struct {
	ID    string
	Email string
}
