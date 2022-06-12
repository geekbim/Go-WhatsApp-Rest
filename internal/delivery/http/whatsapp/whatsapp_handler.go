package whatsapp_handler

import (
	"go-wa-rest/domain/usecase"
	whatsapp_usecase "go-wa-rest/internal/usecase/whatsapp"

	"github.com/gorilla/mux"
	"go.mau.fi/whatsmeow"
)

type whatsAppHandler struct {
	waClient        *whatsmeow.Client
	whatsAppUseCase usecase.WhatsAppUseCase
}

func WhatsAppHandler(
	r *mux.Router,
	waClient *whatsmeow.Client,
) {
	whatsAppUseCase := whatsapp_usecase.NewWhatsAppInteractor(waClient)
	handler := &whatsAppHandler{
		waClient:        waClient,
		whatsAppUseCase: whatsAppUseCase,
	}
	r.HandleFunc("/api/v1/whatsapp/login", handler.Login).Methods("POST")
	r.HandleFunc("/api/v1/whatsapp/send/text", handler.SendText).Methods("POST")
}
