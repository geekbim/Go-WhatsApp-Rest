package response

import "go_wa_rest/domain/entity"

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

func MapWhatsAppDocumentDomainToResponse(whatsAppDocument *entity.WhatsAppDocument) *WhatsApp {
	return &WhatsApp{
		Msisdn:  whatsAppDocument.Msisdn,
		Message: whatsAppDocument.Message,
	}
}
