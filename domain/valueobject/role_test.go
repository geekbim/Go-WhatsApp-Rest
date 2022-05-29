package valueobject_test

import (
	"errors"
	"gokomodo/domain/valueobject"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRole(t *testing.T) {
	t.Run("NewRoleSellerFromInt", func(t *testing.T) {
		roleEnum, err := valueobject.NewRoleFromInt(0)

		assert.Nil(t, err)
		assert.Equal(t, valueobject.USER_ROLE_SELLER, roleEnum)
	})

	t.Run("NewRoleBuyerFromInt", func(t *testing.T) {
		roleEnum, err := valueobject.NewRoleFromInt(1)

		assert.Nil(t, err)
		assert.Equal(t, valueobject.USER_ROLE_BUYER, roleEnum)
	})

	t.Run("NewRoleFromIntErr", func(t *testing.T) {
		_, err := valueobject.NewRoleFromInt(2)

		assert.Equal(t, errors.New("role not found"), err)
	})

	t.Run("NewRoleSellerFromString", func(t *testing.T) {
		roleEnum, err := valueobject.NewRoleFromString("SELLER")

		assert.Nil(t, err)
		assert.Equal(t, valueobject.USER_ROLE_SELLER, roleEnum)
	})

	t.Run("NewRoleBuyerFromString", func(t *testing.T) {
		roleEnum, err := valueobject.NewRoleFromString("BUYER")

		assert.Nil(t, err)
		assert.Equal(t, valueobject.USER_ROLE_BUYER, roleEnum)
	})

	t.Run("NewRoleFromStringErr", func(t *testing.T) {
		_, err := valueobject.NewRoleFromString("GUEST")

		assert.Equal(t, errors.New("role not found"), err)
	})

	t.Run("NewRoleSeller", func(t *testing.T) {
		const USER_ROLE_SELLER valueobject.RoleEnum = 0
		role, err := valueobject.NewRole(USER_ROLE_SELLER)
		isSeller := role.IsSeller()
		roleString := role.String()

		assert.Nil(t, err)
		assert.Equal(t, valueobject.USER_ROLE_SELLER, role.GetValue())
		assert.Equal(t, true, isSeller)
		assert.Equal(t, "SELLER", roleString)
	})

	t.Run("NewRoleBuyer", func(t *testing.T) {
		const USER_ROLE_BUYER valueobject.RoleEnum = 1
		role, err := valueobject.NewRole(USER_ROLE_BUYER)

		assert.Nil(t, err)
		assert.Equal(t, valueobject.USER_ROLE_BUYER, role.GetValue())
		assert.Equal(t, true, role.IsBuyer())
		assert.Equal(t, "BUYER", role.String())
	})

	t.Run("NewRoleErr", func(t *testing.T) {
		const USER_ROLE_GUEST valueobject.RoleEnum = 2
		_, err := valueobject.NewRole(USER_ROLE_GUEST)

		assert.Equal(t, errors.New("invalid role, role not supported"), err)
	})
}
