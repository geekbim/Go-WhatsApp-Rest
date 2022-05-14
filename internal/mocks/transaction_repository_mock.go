package mocks

import (
	"context"
	"majoo/domain/entity"
	"majoo/internal/delivery/request"
	"time"

	"github.com/stretchr/testify/mock"
)

type TransactionRepositoryMock struct {
	mock.Mock
}

func (repository *TransactionRepositoryMock) GetTransactionByUserIdAndDate(ctx context.Context, userId int, startDate, endDate time.Time, options *request.Option) ([]*entity.Transaction, error) {
	args := repository.Called(ctx, userId, startDate, endDate, options)
	return args.Get(0).([]*entity.Transaction), args.Error(1)
}

func (repository *TransactionRepositoryMock) CountTransactionByUserIdAndDate(ctx context.Context, userId int, startDate, endDate time.Time, options *request.Option) (int32, error) {
	args := repository.Called(ctx, userId, startDate, endDate, options)
	return args.Get(0).(int32), args.Error(1)
}
