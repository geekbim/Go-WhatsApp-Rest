package whatsapp

import (
	"go_wa_rest/domain/usecase"

	"go.mau.fi/whatsmeow"
)

type whatsAppInteractor struct {
	waClient *whatsmeow.Client
}

func NewWhatsAppInteractor(
	waClient *whatsmeow.Client,
) usecase.WhatsAppUseCase {
	return &whatsAppInteractor{
		waClient: waClient,
	}
}
