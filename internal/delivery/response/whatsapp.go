package response

import "go-wa-rest/domain/entity"

type WhatsApp struct {
	Msisdn  string `json:"msisdn"`
	Message string `json:"message"`
}

func MapWhatsAppDomainToResponse(whatsApp *entity.WhatsApp) *WhatsApp {
	return &WhatsApp{
		Msisdn:  whatsApp.Msisdn,
		Message: whatsApp.Message,
	}
}
