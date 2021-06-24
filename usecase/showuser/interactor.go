package showuser

import (
	"context"
	"userprofile/application/apperror"
	"userprofile/domain/repository"
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

	err := repository.ReadOnly(ctx, r.outport, func(ctx context.Context) error {

		userObj, err := r.outport.FindOneUserByID(ctx, req.UserID)
		if err != nil {
			return err
		}
		if userObj == nil {
			return apperror.ObjectNotFound.Var(userObj)
		}

		res.User.ID = userObj.ID
		res.User.Email = userObj.Email
		res.User.Address = userObj.Address

		return nil
	})
	if err != nil {
		return nil, err
	}



	return res, nil
}
