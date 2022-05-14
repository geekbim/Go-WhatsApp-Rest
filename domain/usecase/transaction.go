package usecase

import (
	"context"
	"majoo/domain/entity"
	"majoo/internal/delivery/request"
	"majoo/pkg/exceptions"
	"time"
)

type TransactionUseCase interface {
	ListTransaction(ctx context.Context, userId int, startDate, endDate time.Time, options *request.Option) ([]*entity.Transaction, int32, *exceptions.CustomerError)
}
