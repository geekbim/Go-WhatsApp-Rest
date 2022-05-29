package middleware

import (
	"errors"
	"gokomodo/domain/valueobject"
	"gokomodo/pkg/exceptions"

	"github.com/hashicorp/go-multierror"
)

func CheckSellerRole(role string) *exceptions.CustomerError {
	var multierr *multierror.Error

	roleEnum, err := valueobject.NewRoleFromString(role)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		return &exceptions.CustomerError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multierr,
		}
	}
	if roleEnum != valueobject.USER_ROLE_SELLER {
		multierr = multierror.Append(multierr, errors.New("you are not seller"))
		return &exceptions.CustomerError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multierr,
		}
	}

	return nil
}

func CheckBuyerRole(role string) *exceptions.CustomerError {
	var multierr *multierror.Error

	roleEnum, err := valueobject.NewRoleFromString(role)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		return &exceptions.CustomerError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multierr,
		}
	}
	if roleEnum != valueobject.USER_ROLE_BUYER {
		multierr = multierror.Append(multierr, errors.New("you are not buyer"))
		return &exceptions.CustomerError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multierr,
		}
	}

	return nil
}
