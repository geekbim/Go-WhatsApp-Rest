package whatsapp_handler

import (
	"context"
	"go_wa_rest/pkg/exceptions"
	"go_wa_rest/pkg/utils"
	"net/http"
)

func (handler *whatsAppHandler) LogoutV2(w http.ResponseWriter, r *http.Request) {
	errUseCase := handler.whatsAppUseCase.LogoutV2(context.Background(), r.Header.Get("Authorization"))
	if errUseCase != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), errUseCase.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusNoContent, nil)
}