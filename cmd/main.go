package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "go-rest-ddd/docs"
	"go-rest-ddd/internal/config"
	transaction_handler "go-rest-ddd/internal/delivery/http/transaction"
	user_handler "go-rest-ddd/internal/delivery/http/user"
	transaction_repository "go-rest-ddd/internal/repository/psql/transaction"
	user_repository "go-rest-ddd/internal/repository/psql/user"
	"go-rest-ddd/pkg/logger"
	"go-rest-ddd/pkg/service/jwt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
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

	router.PathPrefix("/documentation/").Handler(httpSwagger.WrapHandler)

	appLogger.Info("go-rest-ddd Service Run on " + cfg.Port)

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
