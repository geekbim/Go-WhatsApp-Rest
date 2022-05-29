package usecase

import (
	"context"
	"gokomodo/domain/entity"
	"gokomodo/pkg/exceptions"
)

type UserUseCase interface {
	Login(ctx context.Context, user *entity.User) (*entity.User, *exceptions.CustomerError)
}
