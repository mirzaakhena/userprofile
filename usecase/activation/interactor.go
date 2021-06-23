package activation

import (
	"context"
	"userprofile/application/apperror"
	"userprofile/domain/service"
)

//go:generate mockery --name Outport -output mocks/

type activationInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase Activation
func NewUsecase(outputPort Outport) Inport {
	return &activationInteractor{
		outport: outputPort,
	}
}

// Execute the usecase Activation
func (r *activationInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	//!

	userObj, err := r.outport.FindOneUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if userObj == nil {
		return nil, apperror.EmailIsNotFound
	}

	err = userObj.ValidateToken(req.ActivationToken)
	if err != nil {
		return nil, err
	}

	err = userObj.ActivateUser()
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveUser(ctx, userObj)
	if err != nil {
		return nil, err
	}

	successMessage, err := r.outport.ConstructSuccessActivationMessage(ctx, service.ConstructSuccessActivationMessageServiceRequest{})
	if err != nil {
		return nil, err
	}

	err = r.outport.SendEmail(ctx, service.SendEmailServiceRequest{
		EmailDestination: req.Email,
		Subject:          successMessage.Subject,
		ContentBody:      successMessage.Body,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
