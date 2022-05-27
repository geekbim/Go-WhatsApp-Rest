package user

import (
	"context"
	"go-rest-ddd/domain/entity"
	"go-rest-ddd/pkg/exceptions"

	"github.com/hashicorp/go-multierror"
)

func (interactor *userInteractor) Login(ctx context.Context, userName, password string) (*entity.User, *exceptions.CustomerError) {
	var multierr *multierror.Error

	user, err := interactor.userRepository.GetUserByUserNameAndPassword(ctx, userName, password)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	return user, nil
}
