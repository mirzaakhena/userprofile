package showuser

import (
	"context"
	"userprofile/application/apperror"
)

//go:generate mockery --name Outport -output mocks/

type showUserInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase ShowUser
func NewUsecase(outputPort Outport) Inport {
	return &showUserInteractor{
		outport: outputPort,
	}
}

// Execute the usecase ShowUser
func (r *showUserInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	userObj, err := r.outport.FindOneUserByID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	if userObj == nil {
		return nil, apperror.ObjectNotFound.Var(userObj)
	}

	res.User.ID = userObj.ID
	res.User.Email = userObj.Email
	res.User.Address = userObj.Address

	return res, nil
}
