package user_handler

import (
	"majoo/domain/repository"
	"majoo/domain/usecase"
	"majoo/internal/config"
	"majoo/internal/usecase/user"
	"majoo/pkg/service/jwt"

	"github.com/gorilla/mux"
)

type userHandler struct {
	cfgServer   config.ServerConfig
	jwtService  jwt.JWTService
	userUseCase usecase.UserUseCase
}

func UserHandler(
	r *mux.Router,
	cfgServer config.ServerConfig,
	jwtService jwt.JWTService,
	userRepository repository.UserRepository,
) {
	userUseCase := user.NewUserInteractor(cfgServer, userRepository)
	handler := &userHandler{
		cfgServer:   cfgServer,
		jwtService:  jwtService,
		userUseCase: userUseCase,
	}
	r.HandleFunc("/apis/v1/user/login", handler.Login).Methods("POST")
}
