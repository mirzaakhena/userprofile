package gateway

import (
  "context"
  "userprofile/domain/entity"
  "userprofile/infrastructure/database"
  "userprofile/infrastructure/log"
  "userprofile/infrastructure/token"

  "gorm.io/gorm"
)

type indatabaseGateway struct {
  database.GormReadOnlyImpl
  database.GormTransactionImpl
  commonImplementation
}

// NewIndatabaseGateway ...
func NewIndatabaseGateway(UserToken *token.JWTToken, db *gorm.DB) *indatabaseGateway {

  _ = db.AutoMigrate(&entity.User{})

  return &indatabaseGateway{
    commonImplementation: commonImplementation{
      UserToken: UserToken,
    },
  }
}

func (r *indatabaseGateway) SaveUser(ctx context.Context, obj *entity.User) error {
  log.Info(ctx, "called")

  db, err := database.ExtractDB(ctx)
  if err != nil {
    return err
  }

  err = db.Save(obj).Error
  if err != nil {
    log.Error(ctx, err.Error())
    return err
  }

  return nil
}

func (r *indatabaseGateway) FindOneUserByEmail(ctx context.Context, email string) (*entity.User, error) {
  log.Info(ctx, "called")

  db, err := database.ExtractDB(ctx)
  if err != nil {
    return nil, err
  }

  var obj entity.User
  err = db.
    Where("email = ?", email).
    First(&obj).Error

  if err != nil {
    log.Error(ctx, err.Error())
    return nil, err
  }

  return &obj, nil

}

func (r *indatabaseGateway) FindAllUser(ctx context.Context) ([]*entity.User, error) {
  log.Info(ctx, "called")

  db, err := database.ExtractDB(ctx)
  if err != nil {
    return nil, err
  }

  var objs []*entity.User
  err = db.
    Find(&objs).Error

  if err != nil {
    log.Error(ctx, err.Error())
    return nil, err
  }

  return objs, nil
}

func (r *indatabaseGateway) FindOneUserByID(ctx context.Context, userID string) (*entity.User, error) {
  log.Info(ctx, "called")

  db, err := database.ExtractDB(ctx)
  if err != nil {
    return nil, err
  }

  var obj entity.User
  err = db.
    Where("id = ?", userID).
    First(&obj).Error

  if err != nil {
    log.Error(ctx, err.Error())
    return nil, err
  }

  return &obj, nil

}

func (r *indatabaseGateway) UpdateUserAddress(ctx context.Context, obj *entity.User) error {
  log.Info(ctx, "called")

  db, err := database.ExtractDB(ctx)
  if err != nil {
    return err
  }

  err = db.Save(obj).Error
  if err != nil {
    log.Error(ctx, err.Error())
    return err
  }

  return nil
}

