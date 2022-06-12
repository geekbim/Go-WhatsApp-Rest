package entity

type WhatsApp struct {
	Msisdn  string
	Message string
}

type WhatsAppDTO struct {
	Msisdn  string
	Message string
}

func NewWhatsApp(whatsAppDTO *WhatsAppDTO) *WhatsApp {
	return &WhatsApp{
		Msisdn:  whatsAppDTO.Msisdn,
		Message: whatsAppDTO.Message,
	}
}
