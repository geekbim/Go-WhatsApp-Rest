package testdata

import "go_wa_rest/domain/entity"

func NewWhatsAppDTO() *entity.WhatsAppDTO {
	return &entity.WhatsAppDTO{
		Msisdn:  "622150942316",
		Message: "hello",
	}
}

func NewWhatsApp(whatsAppDTO *entity.WhatsAppDTO) *entity.WhatsApp {
	return &entity.WhatsApp{
		Msisdn:  whatsAppDTO.Msisdn,
		Message: whatsAppDTO.Message,
	}
}
