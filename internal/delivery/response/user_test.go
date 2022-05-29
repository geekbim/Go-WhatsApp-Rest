package response_test

import (
	"gokomodo/internal/delivery/response"
	"gokomodo/testdata"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestUserResponse(t *testing.T) {
	userDTO := testdata.NewUserDTO()
	user := testdata.NewUser(userDTO)

	res := response.MapUserDomainToResponse(user, "token")

	assert.Equal(t, user.Id.String(), res.Id)
	assert.Equal(t, user.Email, res.Email)
	assert.Equal(t, user.Name, res.Name)
	assert.Equal(t, user.Role.String(), res.Role)
}
