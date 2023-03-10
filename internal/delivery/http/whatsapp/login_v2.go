package whatsapp_handler

import (
	"context"
	"go_wa_rest/internal/delivery/response"
	"go_wa_rest/pkg/exceptions"
	"go_wa_rest/pkg/utils"
	"net/http"

	"github.com/hashicorp/go-multierror"
)

func (handler *whatsAppHandler) LoginV2(w http.ResponseWriter, r *http.Request) {
	var multierr *multierror.Error

	authHeader := r.Header.Get("Authorization")
	email, err := handler.jwtService.GetEmailByToken(authHeader)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		customErr := &exceptions.CustomerError{
			Status: exceptions.ERRAUTHORIZED,
			Errors: multierr,
		}
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRAUTHORIZED), customErr.Errors.Errors)
		return
	}

	qrImage, qrTimeOut, errUseCase := handler.whatsAppUseCase.GetQrV2(context.Background(), email)
	if errUseCase != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), errUseCase.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response.MapQRToResponse(qrImage, qrTimeOut))
}
