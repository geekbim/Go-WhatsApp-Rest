package whatsapp_handler

import (
	"context"
	"go_wa_rest/domain/entity"
	"go_wa_rest/internal/delivery/response"
	"go_wa_rest/pkg/exceptions"
	"go_wa_rest/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func (handler *whatsAppHandler) GetMessageStatusV2(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("id")

	params := mux.Vars(r)

	messageId := params["id"]

	whatsAppStatus := entity.WhatsAppStatus{
		MessageId: messageId,
	}

	res, errUseCase := handler.whatsAppUseCase.GetMessageStatusV2(context.Background(), &whatsAppStatus, id)
	if errUseCase != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), errUseCase.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response.MapWhatsAppStatusDomainToResponse(res))
}
