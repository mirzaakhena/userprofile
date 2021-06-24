package registeruser

import (
  "context"
  "userprofile/application/apperror"
  "userprofile/domain/entity"
  "userprofile/domain/repository"
  "userprofile/domain/service"
)

//go:generate mockery --name Outport -output mocks/

type registerUserInteractor struct {
  outport Outport
}

// NewUsecase is constructor for create default implementation of usecase RegisterUser
func NewUsecase(outputPort Outport) Inport {
  return &registerUserInteractor{
    outport: outputPort,
  }
}

// Execute the usecase RegisterUser
func (r *registerUserInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

  res := &InportResponse{}

  err := repository.WithTransaction(ctx, r.outport, func(ctx context.Context) error {

    existingUserObj, err := r.outport.FindOneUserByEmail(ctx, req.Email)
    //if err != nil {
    //	return err
    //}
    if existingUserObj != nil {
      return apperror.EmailAlreadyUsed
    }

    generatedUUID := r.outport.GenerateUUID(ctx)

    randomString := r.outport.GenerateRandomString(ctx)

    hashedPassword, err := r.outport.HashPassword(ctx, req.Password)
    if err != nil {
      return err
    }

    userObj, err := entity.NewUser(entity.UserRequest{
      UserID:          generatedUUID,
      Address:         req.Address,
      Email:           req.Email,
      HashedPassword:  hashedPassword,
      ActivationToken: randomString,
    })
    if err != nil {
      return err
    }

    err = r.outport.SaveUser(ctx, userObj)
    if err != nil {
      return err
    }

    activationMessage, err := r.outport.ConstructStartActivationMessage(ctx, service.ConstructStartActivationMessageServiceRequest{
      Email:           req.Email,
      ActivationToken: randomString,
    })
    if err != nil {
      return err
    }

    err = r.outport.SendEmail(ctx, service.SendEmailServiceRequest{
      EmailDestination: req.Email,
      Subject:          activationMessage.Subject,
      ContentBody:      activationMessage.Body,
    })
    if err != nil {
      return err
    }

    return nil
  })

  if err != nil {
    return nil, err
  }

  return res, nil
}
