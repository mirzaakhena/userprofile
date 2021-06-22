package entity

import (
	"fmt"
	"strings"
	"userprofile/application/apperror"
	"userprofile/domain/vo"
)

type User struct {
	ID              string        `` //
	Address         string        `` //
	Email           string        `` //
	HashedPassword  string        `` //
	Status          vo.UserStatus `` //
	ActivationToken string        `` //
}

type UserRequest struct {
	UserID          string `` //
	Address         string `` //
	Email           string `` //
	HashedPassword  string `` //
	ActivationToken string `` //
}

func NewUser(req UserRequest) (*User, error) {

	// validation
	{
		if strings.TrimSpace(req.Address) == "" {
			return nil, apperror.AddressMustNotEmpty
		}

		if strings.TrimSpace(req.UserID) == "" {
			return nil, apperror.IDMustNotEmpty
		}

		if strings.TrimSpace(req.Email) == "" {
			return nil, apperror.EmailMustNotEmpty
		}

		if strings.TrimSpace(req.HashedPassword) == "" {
			return nil, apperror.PasswordMustNotEmpty
		}
	}

	var obj User
	obj.Address = req.Address
	obj.Email = req.Email
	obj.ID = req.UserID
	obj.HashedPassword = req.HashedPassword
	obj.Status = vo.NotActiveUserStatusEnum

	return &obj, nil
}

func (r *User) ValidateToken(token string) error {

	if r.Status != vo.NotActiveUserStatusEnum {
		return apperror.UserAlreadyActiveOrSuspended
	}

	if r.ActivationToken != token {
		return apperror.InvalidActivationToken
	}

	return nil
}

func (r *User) ActivateUser() error {

	if r.Status == vo.ActiveUserStatusEnum {
		return apperror.UserAlreadyActive
	}

	r.ActivationToken = ""
	r.Status = vo.ActiveUserStatusEnum

	return nil
}

func (r *User) GetUserToken() string {
	return fmt.Sprintf("%s_%s", r.ID, r.Email)
}

func (r *User) UserIsActive() bool {
	return r.Status == vo.ActiveUserStatusEnum
}

func (r *User) UpdateAddress(newAddress string) error {

	if strings.TrimSpace(newAddress) == "" {
		return apperror.AddressMustNotEmpty
	}

	r.Address = newAddress

	return nil
}
