package user_handler

import (
	"context"
	"encoding/json"
	"majoo/internal/delivery/request"
	"majoo/internal/delivery/response"
	"majoo/pkg/utils"
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

	ctx := context.Background()
	res, err := handler.userUseCase.Login(ctx, req.UserName, req.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Errors.Errors)
		return
	}

	token := handler.jwtService.GenerateToken(strconv.Itoa(res.Id))

	utils.RespondWithJSON(w, http.StatusCreated, response.MapUserDomainToResponse(res, token))
}
