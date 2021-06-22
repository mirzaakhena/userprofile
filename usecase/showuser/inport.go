package showuser

import (
	"context"
)

// Inport of ShowUser
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase ShowUser
type InportRequest struct {
	UserID string
}

// InportResponse is response payload after running the usecase ShowUser
type InportResponse struct {
	User UserResponse
}

type UserResponse struct {
	ID      string
	Email   string
	Address string
}
