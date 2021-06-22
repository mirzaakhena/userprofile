package repository

import (
	"context"
	"userprofile/domain/entity"
)

type SaveUserRepo interface {
	SaveUser(ctx context.Context, obj *entity.User) error
}

type FindOneUserByEmailRepo interface {
	FindOneUserByEmail(ctx context.Context, email string) (*entity.User, error)
}

type FindAllUserRepo interface {
	FindAllUser(ctx context.Context) ([]*entity.User, error)
}

type FindOneUserByIDRepo interface {
	FindOneUserByID(ctx context.Context, userID string) (*entity.User, error)
}
