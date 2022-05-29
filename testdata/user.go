package testdata

import (
	"gokomodo/domain/entity"
	"gokomodo/domain/valueobject"
	"gokomodo/internal/repository/psql/models"
	"gokomodo/pkg/common"
	"time"
)

func NewUserDTO() *entity.UserDTO {
	return &entity.UserDTO{
		Email:    "test@email.com",
		Password: "qweasd123",
	}
}

func NewUser(userDTO *entity.UserDTO) *entity.User {
	id, _ := common.StringToID("35da70af-aa50-44dc-ae6b-060a0f9e6933")
	roleEnum, _ := valueobject.NewRoleFromString("SELLER")
	role, _ := valueobject.NewRole(roleEnum)
	return &entity.User{
		Id:        id,
		Email:     userDTO.Email,
		Name:      "test",
		Password:  "$2a$04$RrVwWOU9AjhQ3sO1UAYkE.98pEZRnXffcl7CRWskvejdqXBuiBm2i",
		Address:   "Jakarta Barat",
		Role:      role,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func NewUserModel() *models.User {
	return &models.User{
		Id:        "35da70af-aa50-44dc-ae6b-060a0f9e6933",
		Email:     "test@email.com",
		Name:      "test",
		Password:  "$2a$04$RrVwWOU9AjhQ3sO1UAYkE.98pEZRnXffcl7CRWskvejdqXBuiBm2i",
		Address:   "Jakarta Barat",
		Role:      "SELLER",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}
