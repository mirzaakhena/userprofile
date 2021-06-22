package activation

import (
	"context"
)

// Inport of Activation
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase Activation
type InportRequest struct {
	Email           string
	ActivationToken string
}

// InportResponse is response payload after running the usecase Activation
type InportResponse struct {
}
