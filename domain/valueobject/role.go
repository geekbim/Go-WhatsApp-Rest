package valueobject

import (
	"errors"
)

type RoleEnum int

const (
	USER_ROLE_SELLER RoleEnum = 0
	USER_ROLE_BUYER  RoleEnum = 1
)

type Role struct {
	value RoleEnum
}

func (r *Role) String() string {
	roleString := ""

	switch r.value {
	case USER_ROLE_SELLER:
		roleString = "SELLER"
	case USER_ROLE_BUYER:
		roleString = "BUYER"
	}

	return roleString
}

func (r *Role) GetValue() RoleEnum {
	return r.value
}

func (r *Role) IsSeller() bool {
	return r.value == USER_ROLE_SELLER
}

func (r *Role) IsBuyer() bool {
	return r.value == USER_ROLE_BUYER
}

func NewRole(value RoleEnum) (*Role, error) {
	if value < 0 || value > 1 {
		return nil, errors.New("invalid role, role not supported")
	}

	return &Role{value: value}, nil
}

func NewRoleFromString(str string) (RoleEnum, error) {
	var value RoleEnum

	switch str {
	case "SELLER":
		value = USER_ROLE_SELLER
	case "BUYER":
		value = USER_ROLE_BUYER
	default:
		return value, errors.New("role not found")
	}

	return value, nil
}

func NewRoleFromInt(num int) (RoleEnum, error) {
	var value RoleEnum

	switch num {
	case 0:
		value = USER_ROLE_SELLER
	case 1:
		value = USER_ROLE_BUYER
	default:
		return value, errors.New("role not found")
	}

	return value, nil
}
