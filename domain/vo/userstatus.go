package vo

import (
	"strings"
	"userprofile/application/apperror"
)

type UserStatus string

const (
	NotActiveUserStatusEnum UserStatus = "NOTACTIVE"
	ActiveUserStatusEnum    UserStatus = "ACTIVE"
)

var enumUserStatus = map[UserStatus]UserStatusDetail{
	NotActiveUserStatusEnum: {},
	ActiveUserStatusEnum:    {},
}

type UserStatusDetail struct { //
}

func NewUserStatus(name string) (UserStatus, error) {
	name = strings.ToUpper(name)

	if _, exist := enumUserStatus[UserStatus(name)]; !exist {
		return "", apperror.UnrecognizedEnum.Var(name, "UserStatus")
	}

	return UserStatus(name), nil
}

func (r UserStatus) GetDetail() UserStatusDetail {
	return enumUserStatus[r]
}

func (r UserStatus) PossibleValues() []UserStatus {
	res := make([]UserStatus, len(enumUserStatus))
	for key := range enumUserStatus {
		res = append(res, key)
	}
	return res
}
