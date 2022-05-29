package mapper

import (
	"gokomodo/domain/entity"
	"gokomodo/domain/valueobject"
	"gokomodo/internal/repository/psql/models"
	"gokomodo/pkg/common"
)

func ToDomainUser(m *models.User) *entity.User {
	id, _ := common.StringToID(m.Id)
	roleEnum, _ := valueobject.NewRoleFromString(m.Role)
	role, _ := valueobject.NewRole(roleEnum)
	user := &entity.User{
		Id:        id,
		Email:     m.Email,
		Name:      m.Name,
		Password:  m.Password,
		Address:   m.Address,
		Role:      role,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}

	return user
}
