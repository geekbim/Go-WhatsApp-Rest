package main

import (
	"context"
	"fmt"
	"go_wa_rest/internal/config"
	docs_handler "go_wa_rest/internal/delivery/http/docs"
	whatsapp_handler "go_wa_rest/internal/delivery/http/whatsapp"
	whatsapp_service "go_wa_rest/internal/service/whatsapp"
	"go_wa_rest/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"go.mau.fi/whatsmeow"
)

var (
	cfg       = config.Server()
	appLogger = logger.NewApiLogger()
)

func main() {
	waClient := whatsapp_service.InitWhatsApp()

	router := mux.NewRouter()

	initHandler(router, waClient, cfg)
	http.Handle("/", router)
	router.HandleFunc("/api", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "ok")
	})

	appLogger.Info("go_wa_rest Service Run on " + cfg.Port)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		err := http.ListenAndServe(":"+cfg.Port, router)
		if err != nil {
			appLogger.Error(err)
			cancel()
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	select {
	case v := <-quit:
		waClient.Disconnect()
		appLogger.Error(fmt.Sprintf("signal.Notify: %v", v))
	case done := <-ctx.Done():
		appLogger.Error(fmt.Sprintf("ctx.Done: %v", done))
	}
}

func initHandler(router *mux.Router, waClient *whatsmeow.Client, cfg config.ServerConfig) {
	docs_handler.DocsHandler(router, cfg)
	whatsapp_handler.NewWhatsAppHandler(router, waClient)
}
