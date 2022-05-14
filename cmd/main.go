package main

import (
	"context"
	"database/sql"
	"fmt"
	"majoo/internal/config"
	transaction_handler "majoo/internal/delivery/http/transaction"
	user_handler "majoo/internal/delivery/http/user"
	transaction_repository "majoo/internal/repository/psql/transaction"
	user_repository "majoo/internal/repository/psql/user"
	"majoo/pkg/logger"
	"majoo/pkg/service/jwt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

var (
	cfg             = config.Server()
	appLogger       = logger.NewApiLogger()
	db              = config.InitDatabase()
	jwtService      = jwt.NewJWTService()
	userRepo        = user_repository.NewUserRepository(db)
	transactionRepo = transaction_repository.NewTransactionRepository(db)
)

func main() {
	psqlConn := config.InitDatabase()
	defer func(db *sql.DB) { _ = db.Close() }(psqlConn)

	router := mux.NewRouter()

	initHandler(router)
	http.Handle("/", router)
	appLogger.Info("Majoo Service Run on " + cfg.Port)

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
		appLogger.Error(fmt.Sprintf("signal.Notify: %v", v))
	case done := <-ctx.Done():
		appLogger.Error(fmt.Sprintf("ctx.Done: %v", done))
	}
}

func initHandler(router *mux.Router) {
	user_handler.UserHandler(router, jwtService, userRepo)
	transaction_handler.TransactionHandler(router, jwtService, transactionRepo)
}
