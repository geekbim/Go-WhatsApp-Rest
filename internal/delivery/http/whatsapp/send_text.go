package whatsapp_handler

import (
	"context"
	"encoding/json"
	"go-wa-rest/domain/entity"
	"go-wa-rest/internal/delivery/request"
	"go-wa-rest/internal/delivery/response"
	"go-wa-rest/pkg/exceptions"
	"go-wa-rest/pkg/utils"
	"net/http"
)

func (handler *whatsAppHandler) SendText(w http.ResponseWriter, r *http.Request) {
	var req request.WhatsApp

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRDOMAIN), []error{err})
		return
	}

	whatsApp := entity.NewWhatsApp(&entity.WhatsAppDTO{
		Msisdn:  req.Msisdn,
		Message: req.Message,
	})

	whatsApp, errUseCase := handler.whatsAppUseCase.SendMessage(context.Background(), whatsApp)
	if errUseCase != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), errUseCase.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response.MapWhatsAppDomainToResponse(whatsApp))
}
