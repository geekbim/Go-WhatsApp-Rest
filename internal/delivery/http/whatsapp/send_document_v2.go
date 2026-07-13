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
	"log"
	"net/http"
)

func (handler *whatsAppHandler) SendDocumentV2(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("id")

	chatType := r.FormValue("chatType")
	msisdn := r.FormValue("msisdn")
	message := r.FormValue("message")

	file, fileHeader, err := r.FormFile("document")
	if err != nil {
		log.Printf("[whatsapp.send_document_v2] read document error id=%s chatType=%s msisdn=%s err=%v", id, chatType, msisdn, err)
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), []error{errors.New("failed to read document file")})
		return
	}
	defer file.Close()

	fileName := fileHeader.Filename
	fileType := fileHeader.Header.Get("Content-Type")

	documentBytes, err := io.ReadAll(file)
	if err != nil {
		log.Printf("[whatsapp.send_document_v2] read bytes error id=%s chatType=%s msisdn=%s file=%s err=%v", id, chatType, msisdn, fileName, err)
		utils.RespondWithError(w, http.StatusInternalServerError, []error{err})
		return
	}

	newChatType, err := valueobject.NewChatTypeFromString(chatType)
	if err != nil {
		log.Printf("[whatsapp.send_document_v2] chat type error id=%s chatType=%s msisdn=%s err=%v", id, chatType, msisdn, err)
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), []error{err})
		return
	}

	whatsAppDocument, errValidate := entity.NewWhatsAppDocument(&entity.WhatsAppDocumentDTO{
		ChatType: newChatType.GetValue(),
		Msisdn:   msisdn,
		Message:  message,
		Mentions: formMentions(r),
		Document: documentBytes,
		FileName: fileName,
		FileType: fileType,
	})
	if errValidate != nil {
		log.Printf("[whatsapp.send_document_v2] validation error id=%s chatType=%s msisdn=%s file=%s err=%v", id, chatType, msisdn, fileName, errValidate.Error())
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), errValidate.Errors)
		return
	}

	whatsAppDocument, errUseCase := handler.whatsAppUseCase.SendDocumentV2(context.Background(), whatsAppDocument, id)
	if errUseCase != nil {
		log.Printf("[whatsapp.send_document_v2] usecase error id=%s chatType=%s msisdn=%s file=%s err=%v", id, chatType, msisdn, fileName, errUseCase.Errors.Error())
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), errUseCase.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response.MapWhatsAppDocumentDomainToResponse(whatsAppDocument))
}
