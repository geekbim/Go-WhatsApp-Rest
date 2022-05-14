package transaction

import (
	"majoo/domain/repository"
	"majoo/domain/usecase"
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
