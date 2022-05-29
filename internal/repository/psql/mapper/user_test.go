package mapper_test

import (
	"gokomodo/internal/repository/psql/mapper"
	"gokomodo/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserMapper(t *testing.T) {
	userModel := testdata.NewUserModel()
	userDTO := testdata.NewUserDTO()
	userDomain := testdata.NewUser(userDTO)

	res := mapper.ToDomainUser(userModel)

	assert.Equal(t, res, userDomain)
}
