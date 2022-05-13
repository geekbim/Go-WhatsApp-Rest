package user

import (
	"majoo/domain/repository"
	"majoo/domain/usecase"
	"majoo/internal/config"
)

type userInteractor struct {
	cfgServer      config.ServerConfig
	userRepository repository.UserRepository
}

func NewUserInteractor(
	cfgServer config.ServerConfig,
	userRepository repository.UserRepository,
) usecase.UserUseCase {
	return &userInteractor{
		cfgServer:      cfgServer,
		userRepository: userRepository,
	}
}
