package loginuser

import (
	"context"
	"userprofile/application/apperror"
	"userprofile/domain/service"
)

//go:generate mockery --name Outport -output mocks/

type loginUserInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase LoginUser
func NewUsecase(outputPort Outport) Inport {
	return &loginUserInteractor{
		outport: outputPort,
	}
}

// Execute the usecase LoginUser
func (r *loginUserInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	userObj, err := r.outport.FindOneUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if userObj == nil {
		return nil, apperror.InvalidEmailOrPassword
	}

	err = r.outport.ValidatePassword(ctx, service.ValidatePasswordServiceRequest{
		PlainPassword:  req.Password,
		HashedPassword: userObj.HashedPassword,
	})
	if err != nil {
		return nil, err
	}

	if !userObj.UserIsActive() {
		return nil, apperror.UserIsNotActive
	}

	tokenString := userObj.GetUserToken()

	generateTokenResponse, err := r.outport.GenerateToken(ctx, service.GenerateTokenServiceRequest{
		RawContent: tokenString,
	})
	if err != nil {
		return nil, err
	}

	res.Token = generateTokenResponse

	return res, nil
}
