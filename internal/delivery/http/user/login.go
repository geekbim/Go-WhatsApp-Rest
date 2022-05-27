package user_handler

import (
	"context"
	"encoding/json"
	"go-rest-ddd/domain/entity"
	"go-rest-ddd/internal/delivery/request"
	"go-rest-ddd/internal/delivery/response"
	"go-rest-ddd/pkg/utils"
	"net/http"
	"strconv"
)

func (handler *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req request.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, []error{err})
		return
	}

	user, errValidate := entity.NewUser(&entity.UserDTO{
		UserName: req.UserName,
		Password: req.Password,
	})
	if errValidate != nil {
		utils.RespondWithError(w, http.StatusBadRequest, []error{errValidate})
		return
	}

	ctx := context.Background()
	res, err := handler.userUseCase.Login(ctx, user.UserName, user.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Errors.Errors)
		return
	}

	token := handler.jwtService.GenerateToken(strconv.Itoa(res.Id))

	utils.RespondWithJSON(w, http.StatusCreated, response.MapUserDomainToResponse(res, token))
}
