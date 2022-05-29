package entity_test

import (
	"errors"
	"gokomodo/domain/entity"
	"gokomodo/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserDomain(t *testing.T) {
	userDTO := testdata.NewUserDTO()

	t.Run("NewUser", func(t *testing.T) {
		res, err := entity.NewUser(userDTO)

		assert.Nil(t, err)
		assert.Equal(t, userDTO.Email, res.Email)
		assert.Equal(t, userDTO.Password, res.Password)
	})

	t.Run("NewUserErrEmail", func(t *testing.T) {
		userDTO.Email = ""
		err := errors.New("email cannot be empty")
		expectedErr := []error{
			err,
		}

		res, errEntity := entity.NewUser(userDTO)

		assert.Equal(t, expectedErr, errEntity.Errors)
		assert.Nil(t, res)
	})

	t.Run("NewUserErrPassword", func(t *testing.T) {
		userDTO.Email = "test@email.com"
		userDTO.Password = ""
		err := errors.New("password cannot be empty")
		expectedErr := []error{
			err,
		}

		res, errEntity := entity.NewUser(userDTO)

		assert.Equal(t, expectedErr, errEntity.Errors)
		assert.Nil(t, res)
	})
}
