package response

import "go-rest-ddd/domain/entity"

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"userName"`
	Token    string `json:"token"`
}

func MapUserDomainToResponse(user *entity.User, token string) *User {
	return &User{
		Id:       user.Id,
		UserName: user.Name,
		Token:    token,
	}
}
