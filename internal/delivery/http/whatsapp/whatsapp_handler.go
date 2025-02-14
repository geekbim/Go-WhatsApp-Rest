package whatsapp_handler

import (
	"go_wa_rest/domain/usecase"
	whatsapp_service "go_wa_rest/internal/service/whatsapp"
	whatsapp_usecase "go_wa_rest/internal/usecase/whatsapp"
	jwt_service "go_wa_rest/pkg/service/jwt"

	"github.com/gorilla/mux"
	"go.mau.fi/whatsmeow"
)

type whatsAppHandler struct {
	waClient        *whatsmeow.Client
	jwtService      jwt_service.JWTService
	whatsAppUseCase usecase.WhatsAppUseCase
}

func NewWhatsAppHandler(
	r *mux.Router,
	waClient *whatsmeow.Client,
) {
	whatsAppService := whatsapp_service.NewWhatsAppService()
	whatsAppUseCase := whatsapp_usecase.NewWhatsAppInteractor(waClient, whatsAppService)
	jwtService := jwt_service.NewJWTService()
	handler := &whatsAppHandler{
		waClient:        waClient,
		jwtService:      jwtService,
		whatsAppUseCase: whatsAppUseCase,
	}
	r.HandleFunc("/api/v1/whatsapp/login", handler.Login).Methods("POST")
	r.HandleFunc("/api/v1/whatsapp/send/text", handler.SendText).Methods("POST")
	r.HandleFunc("/api/v1/whatsapp/send/document", handler.SendDocument).Methods("POST")
	r.HandleFunc("/api/v1/whatsapp/group", handler.GetGroup).Methods("GET")
	r.HandleFunc("/api/v1/whatsapp/logout", handler.Logout).Methods("POST")

	r.HandleFunc("/api/v2/whatsapp/login", handler.LoginV2).Methods("POST")
	r.HandleFunc("/api/v2/whatsapp/send/text", handler.SendTextV2).Methods("POST")
	r.HandleFunc("/api/v2/whatsapp/logout", handler.LogoutV2).Methods("POST")
}
