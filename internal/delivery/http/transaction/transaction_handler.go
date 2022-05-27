package transaction_handler

import (
	"go-rest-ddd/domain/repository"
	"go-rest-ddd/domain/usecase"
	"go-rest-ddd/internal/usecase/transaction"
	"go-rest-ddd/pkg/service/jwt"

	"github.com/gorilla/mux"
)

type transactionHandler struct {
	jwtService         jwt.JWTService
	transactionUseCase usecase.TransactionUseCase
}

func TransactionHandler(
	r *mux.Router,
	jwtService jwt.JWTService,
	transactionRepository repository.TransactionRepository,
) {
	transactionUseCase := transaction.NewTransactionInteractor(transactionRepository)
	handler := &transactionHandler{
		jwtService:         jwtService,
		transactionUseCase: transactionUseCase,
	}
	r.HandleFunc("/apis/v1/transaction/list", handler.GetList).Methods("GET")
}
