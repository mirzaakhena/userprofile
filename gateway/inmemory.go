package gateway

import (
	"context"
	"userprofile/application/apperror"
	"userprofile/domain/entity"
	"userprofile/infrastructure/log"
	"userprofile/infrastructure/token"
)

// const SECRET_KEY = "this secret key must stored in save place"

type inmemoryGateway struct {
	commonImplementation
	Users []*entity.User
}

// NewInmemoryGateway ...
func NewInmemoryGateway(UserToken *token.JWTToken) *inmemoryGateway {
	return &inmemoryGateway{
		commonImplementation: commonImplementation{
			UserToken: UserToken,
		},
		Users: make([]*entity.User, 0),
	}
}

func (r *inmemoryGateway) SaveUser(ctx context.Context, obj *entity.User) error {
	log.Info(ctx, "called")

	for _, v := range r.Users {
		if v.ID == obj.ID {
			v.Address = obj.Address
			return nil
		}
	}

	r.Users = append(r.Users, obj)

	return nil
}

func (r *inmemoryGateway) FindOneUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	log.Info(ctx, "called")

	for _, v := range r.Users {
		if v.Email == email {
			return v, nil
		}
	}

	return nil, apperror.UserIsNotFound
}

func (r *inmemoryGateway) FindAllUser(ctx context.Context) ([]*entity.User, error) {
	log.Info(ctx, "called")

	return r.Users, nil
}

func (r *inmemoryGateway) FindOneUserByID(ctx context.Context, userID string) (*entity.User, error) {
	log.Info(ctx, "called")

	for _, v := range r.Users {
		if v.ID == userID {
			return v, nil
		}
	}

	return nil, apperror.UserIsNotFound

}

func (r *inmemoryGateway) UpdateUserAddress(ctx context.Context, obj *entity.User) error {
	log.Info(ctx, "called")

	for _, v := range r.Users {
		if v.ID == obj.ID {
			v.Address = obj.Address
			return nil
		}
	}

	return apperror.UserIsNotFound
}
