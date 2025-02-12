package whatsapp_handler

import (
	"context"
	"go_wa_rest/pkg/exceptions"
	"go_wa_rest/pkg/utils"
	"net/http"
)

func (handler *whatsAppHandler) GetGroup(w http.ResponseWriter, r *http.Request) {
	groups, errUseCase := handler.whatsAppUseCase.GetGroup(context.Background())
	if errUseCase != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), errUseCase.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, groups)
}
