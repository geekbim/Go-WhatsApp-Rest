package whatsapp_handler

import (
	"context"
	"errors"
	"go_wa_rest/domain/entity"
	"go_wa_rest/internal/delivery/response"
	"go_wa_rest/pkg/exceptions"
	"go_wa_rest/pkg/utils"
	"go_wa_rest/valueobject"
	"io"
	"net/http"
)

func (handler *whatsAppHandler) SendDocumentV2(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("id")

	chatType := r.FormValue("chatType")
	msisdn := r.FormValue("msisdn")

	file, fileHeader, err := r.FormFile("document")
	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), []error{errors.New("failed to read document file")})
		return
	}
	defer file.Close()

	fileName := fileHeader.Filename
	fileType := fileHeader.Header.Get("Content-Type")

	documentBytes, err := io.ReadAll(file)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, []error{err})
		return
	}

	newChatType, err := valueobject.NewChatTypeFromString(chatType)
	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), []error{err})
		return
	}

	whatsAppDocument, errValidate := entity.NewWhatsAppDocument(&entity.WhatsAppDocumentDTO{
		ChatType: newChatType.GetValue(),
		Msisdn:   msisdn,
		Document: documentBytes,
		FileName: fileName,
		FileType: fileType,
	})
	if errValidate != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), errValidate.Errors)
		return
	}

	whatsAppDocument, errUseCase := handler.whatsAppUseCase.SendDocumentV2(context.Background(), whatsAppDocument, id)
	if errUseCase != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), errUseCase.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response.MapWhatsAppDocumentDomainToResponse(whatsAppDocument))
}
