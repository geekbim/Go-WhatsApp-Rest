package entity

import (
	"errors"
	"gokomodo/domain/valueobject"
	"gokomodo/pkg/common"
	"time"

	"github.com/hashicorp/go-multierror"
)

type User struct {
	Id        common.ID
	Email     string
	Name      string
	Password  string
	Address   string
	Role      *valueobject.Role
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserDTO struct {
	Email    string
	Password string
}

func NewUser(userDTO *UserDTO) (*User, *multierror.Error) {
	var multierr *multierror.Error

	user := &User{
		Email:    userDTO.Email,
		Password: userDTO.Password,
	}

	if errValidate := user.Validate(); errValidate != nil {
		multierr = multierror.Append(multierr, errValidate)
	}

	if multierr != nil {
		return nil, multierr
	}

	return user, nil
}

func (user *User) Validate() *multierror.Error {
	var multierr *multierror.Error

	if user.Email == "" {
		multierr = multierror.Append(multierr, errors.New("email cannot be empty"))
	}

	if user.Password == "" {
		multierr = multierror.Append(multierr, errors.New("password cannot be empty"))
	}

	return multierr
}
