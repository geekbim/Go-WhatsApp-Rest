package whatsapp

import (
	"go_wa_rest/domain/service"
	"go_wa_rest/domain/usecase"

	"go.mau.fi/whatsmeow"
)

type whatsAppInteractor struct {
	waClient        *whatsmeow.Client
	whatsAppService service.WhatsAppService
}

func NewWhatsAppInteractor(
	waClient *whatsmeow.Client,
	whatsAppService service.WhatsAppService,
) usecase.WhatsAppUseCase {
	return &whatsAppInteractor{
		waClient:        waClient,
		whatsAppService: whatsAppService,
	}
}
