package user_test

import (
	"context"
	"errors"
	"gokomodo/internal/mocks"
	user_usecase "gokomodo/internal/usecase/user"
	"gokomodo/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	ctx := context.TODO()

	userRepo := new(mocks.UserRepositoryMock)

	userDTO := testdata.NewUserDTO()
	user := testdata.NewUser(userDTO)
	user.Password = "qweasd123"
	user1 := testdata.NewUser(userDTO)

	userRepo.
		On("FindUserByEmail", mock.Anything, user.Email).
		Return(user1, nil)

	useCase := user_usecase.NewUserInteractor(userRepo)

	res, err := useCase.Login(ctx, user)

	assert.Nil(t, err)
	assert.Equal(t, user.Email, res.Email)
	assert.Equal(t, user1.Password, res.Password)
}

func TestLoginErr(t *testing.T) {
	ctx := context.TODO()

	userRepo := new(mocks.UserRepositoryMock)

	userDTO := testdata.NewUserDTO()
	user := testdata.NewUser(userDTO)

	err := errors.New("account not found")
	expectedErr := []error{
		err,
	}

	userRepo.
		On("FindUserByEmail", mock.Anything, user.Email).
		Return(user, err)

	useCase := user_usecase.NewUserInteractor(userRepo)

	res, errUseCase := useCase.Login(ctx, user)

	assert.Nil(t, res)
	assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
}
