package whatsapp_handler

import (
	"context"
	"encoding/json"
	"go_wa_rest/internal/delivery/request"
	"go_wa_rest/pkg/exceptions"
	"go_wa_rest/pkg/utils"
	"go_wa_rest/valueobject"
	"net/http"
)

func (handler *whatsAppHandler) StartTypingV2(w http.ResponseWriter, r *http.Request) {
	var req request.WhatsAppTyping

	id := r.Header.Get("id")

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRDOMAIN), []error{err})
		return
	}

	chatType, err := valueobject.NewChatTypeFromString(req.ChatType)
	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), []error{err})
		return
	}

	errUseCase := handler.whatsAppUseCase.StartTypingV2(context.Background(), chatType.GetValue(), req.Msisdn, id, req.IsAudio)
	if errUseCase != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), errUseCase.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "typing_started"})
}
