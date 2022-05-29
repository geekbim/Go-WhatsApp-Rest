package mocks

import (
	"context"
	"gokomodo/domain/entity"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (repository *UserRepositoryMock) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	args := repository.Called(ctx, email)
	return args.Get(0).(*entity.User), args.Error(1)
}
