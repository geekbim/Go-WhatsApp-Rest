package transaction_test

import (
	"context"
	"errors"
	"go-rest-ddd/domain/entity"
	"go-rest-ddd/internal/mocks"
	transaction_usecase "go-rest-ddd/internal/usecase/transaction"
	"go-rest-ddd/testdata"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListTransaction(t *testing.T) {
	ctx := context.TODO()

	transactionRepo := new(mocks.TransactionRepositoryMock)

	transactions := []*entity.Transaction{testdata.NewTransaction()}

	startDate, _ := time.Parse("2006-01-02", "2021-11-01")
	endDate, _ := time.Parse("2006-01-02", "2021-11-30")

	transactionRepo.
		On("GetTransactionByUserIdAndDate", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(transactions, nil)
	transactionRepo.
		On("CountTransactionByUserIdAndDate", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(int32(len(transactions)), nil)

	useCase := transaction_usecase.NewTransactionInteractor(transactionRepo)

	res, count, err := useCase.ListTransaction(ctx, 1, startDate, endDate, nil)
	assert.Nil(t, err)
	assert.Equal(t, transactions[0].Id, res[0].Id)
	assert.Equal(t, int32(len(transactions)), count)
}

func TestListTransactionErrGet(t *testing.T) {
	ctx := context.TODO()

	transactionRepo := new(mocks.TransactionRepositoryMock)

	transactions := []*entity.Transaction{testdata.NewTransaction()}

	startDate, _ := time.Parse("2006-01-02", "2021-11-01")
	endDate, _ := time.Parse("2006-01-02", "2021-11-30")

	err := errors.New("error get list transaction")
	expectedErr := []error{
		err,
	}

	transactionRepo.
		On("GetTransactionByUserIdAndDate", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(transactions, err)
	transactionRepo.
		On("CountTransactionByUserIdAndDate", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(int32(len(transactions)), nil)

	useCase := transaction_usecase.NewTransactionInteractor(transactionRepo)

	res, count, errUseCase := useCase.ListTransaction(ctx, 1, startDate, endDate, nil)
	assert.Nil(t, res)
	assert.Equal(t, int32(0), count)
	assert.Equal(t, errUseCase.Errors.Errors, expectedErr)
}

func TestListTransactionErrCount(t *testing.T) {
	ctx := context.TODO()

	transactionRepo := new(mocks.TransactionRepositoryMock)

	transactions := []*entity.Transaction{testdata.NewTransaction()}

	startDate, _ := time.Parse("2006-01-02", "2021-11-01")
	endDate, _ := time.Parse("2006-01-02", "2021-11-30")

	err := errors.New("error count list transaction")
	expectedErr := []error{
		err,
	}

	transactionRepo.
		On("GetTransactionByUserIdAndDate", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(transactions, nil)
	transactionRepo.
		On("CountTransactionByUserIdAndDate", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(int32(len(transactions)), err)

	useCase := transaction_usecase.NewTransactionInteractor(transactionRepo)

	res, count, errUseCase := useCase.ListTransaction(ctx, 1, startDate, endDate, nil)
	assert.Nil(t, res)
	assert.Equal(t, int32(0), count)
	assert.Equal(t, errUseCase.Errors.Errors, expectedErr)
}
