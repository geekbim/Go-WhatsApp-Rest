package user

import (
	"go-rest-ddd/domain/repository"
	"go-rest-ddd/domain/usecase"
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
