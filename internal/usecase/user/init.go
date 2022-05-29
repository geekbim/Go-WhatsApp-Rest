package user

import (
	"gokomodo/domain/repository"
	"gokomodo/domain/usecase"
)

type userInteractor struct {
	userRepository repository.UserRepository
}

func NewUserInteractor(
	userRepository repository.UserRepository,
) usecase.UserUseCase {
	return &userInteractor{
		userRepository: userRepository,
	}
}
