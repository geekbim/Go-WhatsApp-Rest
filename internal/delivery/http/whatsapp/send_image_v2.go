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

func (handler *whatsAppHandler) SendImageV2(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("id")

	chatType := r.FormValue("chatType")
	msisdn := r.FormValue("msisdn")
	message := r.FormValue("message")

	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), []error{errors.New("failed to read image file")})
		return
	}
	defer file.Close()

	fileName := fileHeader.Filename
	fileType := fileHeader.Header.Get("Content-Type")

	imageBytes, err := io.ReadAll(file)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, []error{err})
		return
	}

	newChatType, err := valueobject.NewChatTypeFromString(chatType)
	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), []error{err})
		return
	}

	whatsAppImage, errValidate := entity.NewWhatsAppImage(&entity.WhatsAppImageDTO{
		ChatType: newChatType.GetValue(),
		Msisdn:   msisdn,
		Message:  message,
		Image:    imageBytes,
		FileName: fileName,
		FileType: fileType,
	})
	if errValidate != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), errValidate.Errors)
		return
	}

	whatsAppImage, errUseCase := handler.whatsAppUseCase.SendImageV2(context.Background(), whatsAppImage, id)
	if errUseCase != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), errUseCase.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response.MapWhatsAppImageDomainToResponse(whatsAppImage))
}
