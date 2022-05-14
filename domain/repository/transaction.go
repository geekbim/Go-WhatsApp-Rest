package repository

import (
	"context"
	"majoo/domain/entity"
	"majoo/internal/delivery/request"
	"time"
)

type TransactionRepository interface {
	GetTransactionByUserIdAndDate(ctx context.Context, userId int, startDate, endDate time.Time, options *request.Option) ([]*entity.Transaction, error)
}
