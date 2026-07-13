package whatsapp_handler

import (
	"context"
	"go_wa_rest/pkg/exceptions"
	"go_wa_rest/pkg/utils"
	"net/http"
)

func (handler *whatsAppHandler) GetContact(w http.ResponseWriter, r *http.Request) {
	contacts, errUseCase := handler.whatsAppUseCase.GetContact(context.Background())
	if errUseCase != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), errUseCase.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, contacts)
}
