package user_handler

import (
	"context"
	"encoding/json"
	"gokomodo/domain/entity"
	"gokomodo/internal/delivery/request"
	"gokomodo/internal/delivery/response"
	"gokomodo/pkg/exceptions"
	"gokomodo/pkg/utils"
	"net/http"
)

func (handler *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req request.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRDOMAIN), []error{err})
		return
	}

	user, errValidate := entity.NewUser(&entity.UserDTO{
		Email:    req.Email,
		Password: req.Password,
	})
	if errValidate != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRDOMAIN), errValidate.Errors)
		return
	}

	ctx := context.Background()
	user, err := handler.userUseCase.Login(ctx, user)
	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), err.Errors.Errors)
		return
	}

	token := handler.jwtService.GenerateToken(user.Id.String(), user.Role.String())

	utils.RespondWithJSON(w, http.StatusOK, response.MapUserDomainToResponse(user, token))
}
