package whatsapp_handler

import (
	"context"
	"errors"
	"go_wa_rest/domain/entity"
	"go_wa_rest/internal/delivery/response"
	"go_wa_rest/pkg/exceptions"
	"go_wa_rest/pkg/utils"
	"io"
	"net/http"
)

func (handler *whatsAppHandler) SendDocument(w http.ResponseWriter, r *http.Request) {
	msisdn := r.FormValue("msisdn")
	message := r.FormValue("message")

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

	whatsApp, errValidate := entity.NewWhatsAppDocument(&entity.WhatsAppDocumentDTO{
		Msisdn:   msisdn,
		Message:  message,
		Document: documentBytes,
		FileName: fileName,
		FileType: fileType,
	})
	if errValidate != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), errValidate.Errors)
		return
	}

	whatsAppDocument, errUseCase := handler.whatsAppUseCase.SendDocument(context.Background(), whatsApp)
	if errUseCase != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), errUseCase.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response.MapWhatsAppDocumentDomainToResponse(whatsAppDocument))
}
