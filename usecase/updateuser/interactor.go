package updateuser

import (
  "context"
  "userprofile/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type updateUserInteractor struct {
  outport Outport
}

// NewUsecase is constructor for create default implementation of usecase UpdateUser
func NewUsecase(outputPort Outport) Inport {
  return &updateUserInteractor{
    outport: outputPort,
  }
}

// Execute the usecase UpdateUser
func (r *updateUserInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

  res := &InportResponse{}

  err := repository.WithTransaction(ctx, r.outport, func(ctx context.Context) error {

    userObj, err := r.outport.FindOneUserByID(ctx, req.UserID)
    if err != nil {
      return err
    }

    err = userObj.UpdateAddress(req.Address)
    if err != nil {
      return err
    }

    err = r.outport.SaveUser(ctx, userObj)
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
