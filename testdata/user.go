package testdata

import "majoo/domain/entity"

func NewUserDTO() *entity.UserDTO {
	return &entity.UserDTO{
		UserName: "admin1",
		Password: "admin1",
	}
}

func NewUser(userDTO *entity.UserDTO) *entity.User {
	return &entity.User{
		UserName: userDTO.UserName,
		Password: userDTO.Password,
	}
}
