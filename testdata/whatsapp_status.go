package testdata

import (
	"go_wa_rest/domain/entity"
)

func NewWhatsAppStatusDTO() *entity.WhatsAppStatusDTO {
	return &entity.WhatsAppStatusDTO{
		MessageId: "xxxxxxx",
	}
}

func NewWhatsAppStatus(whatsAppStatusDTO *entity.WhatsAppStatusDTO) *entity.WhatsAppStatus {
	return &entity.WhatsAppStatus{
		MessageId: whatsAppStatusDTO.MessageId,
		Status:    "read",
	}
}
