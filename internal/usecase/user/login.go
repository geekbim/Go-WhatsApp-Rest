package user

import (
	"context"
	"errors"
	"gokomodo/domain/entity"
	"gokomodo/pkg/exceptions"
	"gokomodo/pkg/utils"

	"github.com/hashicorp/go-multierror"
)

func (interactor *userInteractor) Login(ctx context.Context, user *entity.User) (*entity.User, *exceptions.CustomerError) {
	var multierr *multierror.Error

	res, errRepo := interactor.userRepository.FindUserByEmail(ctx, user.Email)
	if errRepo != nil {
		multierr = multierror.Append(multierr, errRepo)
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	if !utils.CheckPasswordHash(user.Password, res.Password) {
		multierr = multierror.Append(multierr, errors.New("password does not match"))
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	return res, nil
}
