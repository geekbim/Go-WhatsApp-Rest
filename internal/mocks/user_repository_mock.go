package mocks

import (
	"context"
	"go-rest-ddd/domain/entity"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (repository *UserRepositoryMock) GetUserByUserNameAndPassword(ctx context.Context, userName, password string) (*entity.User, error) {
	args := repository.Called(ctx, userName, password)
	return args.Get(0).(*entity.User), args.Error(1)
}
