package repository

import (
	"context"
	"go-rest-ddd/domain/entity"
	"go-rest-ddd/internal/delivery/request"
	"time"
)

type TransactionRepository interface {
	GetTransactionByUserIdAndDate(ctx context.Context, userId int, startDate, endDate time.Time, options *request.Option) ([]*entity.Transaction, error)
	CountTransactionByUserIdAndDate(ctx context.Context, userId int, startDate, endDate time.Time, options *request.Option) (int32, error)
}
