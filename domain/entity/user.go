package entity

import (
	"errors"
	"time"
)

type User struct {
	Id        int
	Name      string
	UserName  string
	Password  string
	CreatedAt time.Time
	CreatedBy int
	UpdatedAt time.Time
	UpdatedBy int
}

type UserDTO struct {
	UserName string
	Password string
}

func NewUser(userDTO *UserDTO) (*User, error) {
	user := &User{
		UserName: userDTO.UserName,
		Password: userDTO.Password,
	}

	if errValidate := user.Validate(); errValidate != nil {
		return nil, errValidate
	}

	return user, nil
}

func (user *User) Validate() error {
	if user.UserName == "" {
		return errors.New("username cannot be empty")
	}
	if user.Password == "" {
		return errors.New("password cannot be empty")
	}

	return nil
}
