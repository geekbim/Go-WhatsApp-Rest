package entity_test

import (
	"errors"
	"go-rest-ddd/domain/entity"
	"go-rest-ddd/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	userDTO := testdata.NewUserDTO()

	t.Run("NewUser", func(t *testing.T) {
		res, err := entity.NewUser(userDTO)

		assert.Nil(t, err)
		assert.Equal(t, userDTO.UserName, res.UserName)
	})

	t.Run("NewUserErrUserName", func(t *testing.T) {
		userDTO.UserName = ""
		res, err := entity.NewUser(userDTO)

		expectedErr := errors.New("username cannot be empty")

		assert.Equal(t, expectedErr, err)
		assert.Nil(t, res)
	})

	t.Run("NewUserErrPassword", func(t *testing.T) {
		userDTO.UserName = "admin1"
		userDTO.Password = ""
		res, err := entity.NewUser(userDTO)

		expectedErr := errors.New("password cannot be empty")

		assert.Equal(t, expectedErr, err)
		assert.Nil(t, res)
	})
}
