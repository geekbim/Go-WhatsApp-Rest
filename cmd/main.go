package main

import (
	"context"
	"fmt"
	_ "go-wa-rest/docs"
	"go-wa-rest/internal/config"
	docs_handler "go-wa-rest/internal/delivery/http/docs"
	whatsapp_handler "go-wa-rest/internal/delivery/http/whatsapp"
	"go-wa-rest/pkg/logger"
	"go-wa-rest/pkg/service/whatsapp"
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
	waClient := whatsapp.InitWhatsApp()

	router := mux.NewRouter()

	initHandler(router, waClient)
	http.Handle("/", router)

	appLogger.Info("go-wa-rest Service Run on " + cfg.Port)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		err := http.ListenAndServe(cfg.Port, router)
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

func initHandler(router *mux.Router, waClient *whatsmeow.Client) {
	docs_handler.DocsHandler(router)
	whatsapp_handler.WhatsAppHandler(router, waClient)
}
