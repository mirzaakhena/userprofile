package gateway

import (
	"context"
	"userprofile/domain/entity"
	"userprofile/infrastructure/database"
	"userprofile/infrastructure/log"
	"userprofile/infrastructure/token"

	"gorm.io/gorm"
)

type indatabase2Gateway struct {
	indatabaseGateway
}

// NewIndatabase2Gateway ...
func NewIndatabase2Gateway(UserToken *token.JWTToken, db *gorm.DB) (*indatabase2Gateway, error) {

	x, err := NewIndatabaseGateway(UserToken, db)
	if err != nil {
		return nil, err
	}

	return &indatabase2Gateway{
		indatabaseGateway: *x,
	}, nil
}

func (r *indatabase2Gateway) FindAllUser(ctx context.Context) ([]*entity.User, error) {
	log.Info(ctx, "hellooo world")

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
