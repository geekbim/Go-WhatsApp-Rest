package whatsapp_handler

import (
	"context"
	"encoding/json"
	"go_wa_rest/internal/delivery/request"
	"go_wa_rest/pkg/exceptions"
	"go_wa_rest/pkg/utils"
	"net/http"
)

func (handler *whatsAppHandler) SaveContactV2(w http.ResponseWriter, r *http.Request) {
	var req request.WhatsAppContact

	id := r.Header.Get("id")

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRDOMAIN), []error{err})
		return
	}

	errUseCase := handler.whatsAppUseCase.SaveContactV2(context.Background(), id, req.Msisdn, req.FullName, req.FirstName)
	if errUseCase != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), errUseCase.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

