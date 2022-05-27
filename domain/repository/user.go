package repository

import (
	"context"
	"go-rest-ddd/domain/entity"
)

type UserRepository interface {
	GetUserByUserNameAndPassword(ctx context.Context, userName, password string) (*entity.User, error)
}
