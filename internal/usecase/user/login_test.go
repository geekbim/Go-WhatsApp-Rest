package user_test

import (
	"context"
	"errors"
	"go-rest-ddd/internal/mocks"
	user_usecase "go-rest-ddd/internal/usecase/user"
	"go-rest-ddd/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	ctx := context.TODO()

	userRepo := new(mocks.UserRepositoryMock)

	userDTO := testdata.NewUserDTO()
	user := testdata.NewUser(userDTO)

	userRepo.
		On("GetUserByUserNameAndPassword", mock.Anything, user.UserName, user.Password).
		Return(user, nil)

	useCase := user_usecase.NewUserInteractor(userRepo)

	res, err := useCase.Login(ctx, user.UserName, user.Password)
	assert.Nil(t, err)
	assert.Equal(t, user.UserName, res.UserName)
	assert.Equal(t, user.Password, res.Password)
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
		On("GetUserByUserNameAndPassword", mock.Anything, user.UserName, user.Password).
		Return(user, err)

	useCase := user_usecase.NewUserInteractor(userRepo)

	res, errUseCase := useCase.Login(ctx, user.UserName, user.Password)

	assert.Nil(t, res)
	assert.Equal(t, errUseCase.Errors.Errors, expectedErr)
}
