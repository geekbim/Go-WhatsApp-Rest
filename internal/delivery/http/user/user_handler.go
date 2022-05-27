package user_handler

import (
	"go-rest-ddd/domain/repository"
	"go-rest-ddd/domain/usecase"
	"go-rest-ddd/internal/usecase/user"
	"go-rest-ddd/pkg/service/jwt"

	"github.com/gorilla/mux"
)

type userHandler struct {
	jwtService  jwt.JWTService
	userUseCase usecase.UserUseCase
}

func UserHandler(
	r *mux.Router,
	jwtService jwt.JWTService,
	userRepository repository.UserRepository,
) {
	userUseCase := user.NewUserInteractor(userRepository)
	handler := &userHandler{
		jwtService:  jwtService,
		userUseCase: userUseCase,
	}
	r.HandleFunc("/apis/v1/user/login", handler.Login).Methods("POST")
}
