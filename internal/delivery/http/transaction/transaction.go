package transaction_handler

import (
	"context"
	"fmt"
	"majoo/internal/delivery/request"
	"majoo/internal/delivery/response"
	"majoo/pkg/exceptions"
	"majoo/pkg/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hashicorp/go-multierror"
)

func (handler *transactionHandler) getUserIDByToken(token string) (int, error) {
	aToken, err := handler.jwtService.ValidateToken(token)
	if err != nil {
		return 0, err
	}

	claims := aToken.Claims.(jwt.MapClaims)
	userId, _ := strconv.Atoi(fmt.Sprintf("%v", claims["user_id"]))

	return userId, nil
}

func (handler *transactionHandler) GetList(w http.ResponseWriter, r *http.Request) {
	var (
		query    = r.URL.Query()
		ctx      = context.Background()
		multierr *multierror.Error
	)

	authHeader := r.Header.Get("Authorization")
	userId, err := handler.getUserIDByToken(authHeader)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRAUTHORIZED), multierr.Errors)
		return
	}

	startDate, err := time.Parse("2006-01-02", query.Get("startDate"))
	if err != nil {
		multierr = multierror.Append(multierr, err)
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRDOMAIN), multierr.Errors)
		return
	}

	endDate, err := time.Parse("2006-01-02", query.Get("endDate"))
	if err != nil {
		multierr = multierror.Append(multierr, err)
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRDOMAIN), multierr.Errors)
		return
	}

	limit := utils.StringToInt(query.Get("limit"))
	page := utils.StringToInt(query.Get("page"))

	payload := request.NewOption(limit, page)

	res, errUseCase := handler.transactionUseCase.ListTransaction(ctx, userId, startDate, endDate, payload)
	if errUseCase != nil {
		utils.RespondWithError(w, http.StatusBadRequest, errUseCase.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, response.MapTransactionListDomainToResponse(res))
}
