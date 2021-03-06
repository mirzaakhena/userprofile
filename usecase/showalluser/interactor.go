package showalluser

import (
  "context"
  "userprofile/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type showAllUSerInteractor struct {
  outport Outport
}

// NewUsecase is constructor for create default implementation of usecase ShowAllUSer
func NewUsecase(outputPort Outport) Inport {
  return &showAllUSerInteractor{
    outport: outputPort,
  }
}

// Execute the usecase ShowAllUSer
func (r *showAllUSerInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

  res := &InportResponse{}

  err := repository.ReadOnly(ctx, r.outport, func(ctx context.Context) error {

    userObjs, err := r.outport.FindAllUser(ctx)
    if err != nil {
      return err
    }

    for _, u := range userObjs {
      res.Users = append(res.Users, UserResponse{
        ID:    u.ID,
        Email: u.Email,
      })
    }

    return nil
  })

  if err != nil {
    return nil, err
  }

  return res, nil
}
