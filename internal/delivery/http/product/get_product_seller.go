package product_handler

import (
	"context"
	"gokomodo/internal/delivery/request"
	"gokomodo/internal/delivery/response"
	"gokomodo/pkg/common"
	"gokomodo/pkg/exceptions"
	"gokomodo/pkg/middleware"
	"gokomodo/pkg/utils"
	"net/http"

	"github.com/hashicorp/go-multierror"
)

func (handler *productHandler) GetProductSeller(w http.ResponseWriter, r *http.Request) {
	var (
		ctx      = context.Background()
		query    = r.URL.Query()
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

	errMiddleware := middleware.CheckSellerRole(role)
	if errMiddleware != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRAUTHORIZED), errMiddleware.Errors.Errors)
		return
	}

	limit := utils.StringToInt(query.Get("limit"))
	page := utils.StringToInt(query.Get("page"))

	payload := request.NewOption(limit, page)

	products, count, errUseCase := handler.productUseCase.ListProductSeller(ctx, userUuid, payload)
	if errUseCase != nil {
		utils.RespondWithError(w, http.StatusBadRequest, errUseCase.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response.MapProductListDomainToResponse(products, count))
}
