package testdata

import (
	"go_wa_rest/domain/entity"
	"go_wa_rest/valueobject"
)

func NewWhatsAppDTO() *entity.WhatsAppDTO {
	return &entity.WhatsAppDTO{
		ChatType: valueobject.Private,
		Msisdn:   "622150942316",
		Message:  "hello",
	}
}

func NewWhatsApp(whatsAppDTO *entity.WhatsAppDTO) *entity.WhatsApp {
	return &entity.WhatsApp{
		ChatType: valueobject.NewChatType(whatsAppDTO.ChatType),
		Msisdn:   whatsAppDTO.Msisdn,
		Message:  whatsAppDTO.Message,
	}
}
