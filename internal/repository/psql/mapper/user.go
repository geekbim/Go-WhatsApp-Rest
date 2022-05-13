package mapper

import (
	"majoo/domain/entity"
	"majoo/internal/repository/psql/models"
)

func ToDomainUser(m *models.User) *entity.User {
	return &entity.User{
		Id:        m.Id,
		Name:      m.Name,
		UserName:  m.UserName,
		Password:  m.Password,
		CreatedAt: m.CreatedAt,
		CreatedBy: m.CreatedBy,
		UpdatedAt: m.UpdatedAt,
		UpdatedBy: m.UpdatedBy,
	}
}
