package transaction

import (
	"go-rest-ddd/domain/repository"
	"go-rest-ddd/domain/usecase"
)

type transactionInteractor struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionInteractor(
	transactionRepository repository.TransactionRepository,
) usecase.TransactionUseCase {
	return &transactionInteractor{
		transactionRepository: transactionRepository,
	}
}
