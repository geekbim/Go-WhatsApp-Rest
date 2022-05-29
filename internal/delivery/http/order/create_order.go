package order_handler

import (
	"context"
	"encoding/json"
	"gokomodo/domain/entity"
	"gokomodo/internal/delivery/request"
	"gokomodo/internal/delivery/response"
	"gokomodo/pkg/common"
	"gokomodo/pkg/exceptions"
	"gokomodo/pkg/middleware"
	"gokomodo/pkg/utils"
	"net/http"

	"github.com/hashicorp/go-multierror"
)

func (handler *orderHandler) Store(w http.ResponseWriter, r *http.Request) {
	var (
		ctx      = context.Background()
		req      request.Order
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

	errMiddleware := middleware.CheckBuyerRole(role)
	if errMiddleware != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRAUTHORIZED), errMiddleware.Errors.Errors)
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRDOMAIN), []error{err})
		return
	}

	productId, err := common.StringToID(req.ProductId)
	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRDOMAIN), []error{err})
		return
	}

	order, errValidate := entity.NewOrder(&entity.OrderDTO{
		BuyerId:   userUuid,
		ProductId: productId,
		Qty:       req.Qty,
	})
	if errValidate != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRDOMAIN), errValidate.Errors)
		return
	}

	order, errUseCase := handler.orderUseCase.CreateOrder(ctx, order)
	if errUseCase != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), errUseCase.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, response.MapOrderDomainToResponse(order))
}
