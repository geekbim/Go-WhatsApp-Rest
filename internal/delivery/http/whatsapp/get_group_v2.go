package whatsapp_handler

import (
	"context"
	"go_wa_rest/pkg/exceptions"
	"go_wa_rest/pkg/utils"
	"net/http"
)

func (handler *whatsAppHandler) GetGroupV2(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("id")

	groups, errUseCase := handler.whatsAppUseCase.GetGroupV2(context.Background(), id)
	if errUseCase != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), errUseCase.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, groups)
}
