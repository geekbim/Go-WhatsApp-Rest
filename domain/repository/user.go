package repository

import (
	"context"
	"majoo/domain/entity"
)

type UserRepository interface {
	GetUserByUserNameAndPassword(ctx context.Context, userName, password string) (*entity.User, error)
}
