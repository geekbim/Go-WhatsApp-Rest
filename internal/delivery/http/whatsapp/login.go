package whatsapp_handler

import (
	"context"
	"go-wa-rest/internal/delivery/response"
	"go-wa-rest/pkg/exceptions"
	"go-wa-rest/pkg/utils"
	"net/http"
)

func (handler *whatsAppHandler) Login(w http.ResponseWriter, r *http.Request) {
	qrImage, qrTimeOut, errUseCase := handler.whatsAppUseCase.GetQr(context.Background())
	if errUseCase != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), errUseCase.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response.MapQRToResponse(qrImage, qrTimeOut))
}
