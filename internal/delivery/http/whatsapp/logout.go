package whatsapp_handler

import (
	"context"
	"go_wa_rest/pkg/exceptions"
	"go_wa_rest/pkg/utils"
	"net/http"
)

func (handler *whatsAppHandler) Logout(w http.ResponseWriter, r *http.Request) {
	errUseCase := handler.whatsAppUseCase.Logout(context.Background())
	if errUseCase != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), errUseCase.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusNoContent, nil)
}
