package response

import "go_wa_rest/domain/entity"

type WhatsApp struct {
	Id      string `json:"id"`
	Msisdn  string `json:"msisdn"`
	Message string `json:"message"`
}

func MapWhatsAppDomainToResponse(whatsApp *entity.WhatsApp) *WhatsApp {
	return &WhatsApp{
		Id:      whatsApp.Id,
		Msisdn:  whatsApp.Msisdn,
		Message: whatsApp.Message,
	}
}

func MapWhatsAppDocumentDomainToResponse(whatsAppDocument *entity.WhatsAppDocument) *WhatsApp {
	return &WhatsApp{
		Id:     whatsAppDocument.Id,
		Msisdn: whatsAppDocument.Msisdn,
	}
}

func MapWhatsAppImageDomainToResponse(whatsAppImage *entity.WhatsAppImage) *WhatsApp {
	return &WhatsApp{
		Id:     whatsAppImage.Id,
		Msisdn: whatsAppImage.Msisdn,
	}
}
