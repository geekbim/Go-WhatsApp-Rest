package response

import "gokomodo/domain/entity"

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

func MapUserDomainToResponse(user *entity.User, token string) *User {
	return &User{
		Id:    user.Id.String(),
		Email: user.Email,
		Name:  user.Name,
		Role:  user.Role.String(),
		Token: token,
	}
}
