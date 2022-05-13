package usecase

import (
	"context"
	"majoo/domain/entity"
	"majoo/pkg/exceptions"
)

type UserUseCase interface {
	Login(ctx context.Context, userName, password string) (*entity.User, *exceptions.CustomerError)
}
