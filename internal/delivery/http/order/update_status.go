package order_handler

import (
	"context"
	"gokomodo/internal/delivery/response"
	"gokomodo/pkg/common"
	"gokomodo/pkg/exceptions"
	"gokomodo/pkg/middleware"
	"gokomodo/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-multierror"
)

func (handler *orderHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	var (
		ctx      = context.Background()
		vars     = mux.Vars(r)
		orderId  = vars["id"]
		multierr *multierror.Error
	)

	authHeader := r.Header.Get("Authorization")
	userId, role, err := handler.jwtService.GetUserIdAndRole(authHeader)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRAUTHORIZED), multierr.Errors)
		return
	}

	userUuid, _ := common.StringToID(userId)
	orderUuid, _ := common.StringToID(orderId)

	errMiddleware := middleware.CheckSellerRole(role)
	if errMiddleware != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRAUTHORIZED), errMiddleware.Errors.Errors)
		return
	}

	order, errUseCase := handler.orderUseCase.AcceptOrder(ctx, userUuid, orderUuid)
	if errUseCase != nil {
		utils.RespondWithError(w, http.StatusBadRequest, errUseCase.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, response.MapSellerOrderDomainToResponseList(order))
}
