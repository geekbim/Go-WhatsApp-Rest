package usecase

import (
	"context"
	"go-rest-ddd/domain/entity"
	"go-rest-ddd/pkg/exceptions"
)

type UserUseCase interface {
	Login(ctx context.Context, userName, password string) (*entity.User, *exceptions.CustomerError)
}
