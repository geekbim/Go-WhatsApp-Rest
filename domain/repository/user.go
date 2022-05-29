package repository

import (
	"context"
	"gokomodo/domain/entity"
)

type UserRepository interface {
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
}
