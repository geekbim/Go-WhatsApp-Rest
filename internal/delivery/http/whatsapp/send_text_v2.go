package whatsapp_handler

import (
	"context"
	"encoding/json"
	"go_wa_rest/domain/entity"
	"go_wa_rest/internal/delivery/request"
	"go_wa_rest/internal/delivery/response"
	"go_wa_rest/pkg/exceptions"
	"go_wa_rest/pkg/utils"
	"go_wa_rest/valueobject"
	"net/http"
)

func (handler *whatsAppHandler) SendTextV2(w http.ResponseWriter, r *http.Request) {
	var req request.WhatsApp

	id := r.Header.Get("id")

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRDOMAIN), []error{err})
		return
	}

	whatsApp, errValidate := entity.NewWhatsApp(&entity.WhatsAppDTO{
		ChatType: valueobject.ChatTypeEnum(req.ChatType),
		Msisdn:   req.Msisdn,
		Message:  req.Message,
	})
	if errValidate != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), errValidate.Errors)
		return
	}

	whatsApp, errUseCase := handler.whatsAppUseCase.SendMessageV2(context.Background(), whatsApp, id)
	if errUseCase != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), errUseCase.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response.MapWhatsAppDomainToResponse(whatsApp))
}
