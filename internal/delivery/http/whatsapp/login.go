package whatsapp_handler

import (
	"context"
	"go_wa_rest/internal/delivery/response"
	"go_wa_rest/pkg/exceptions"
	"go_wa_rest/pkg/utils"
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
