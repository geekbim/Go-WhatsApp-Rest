package usecase

import (
	"context"
	"go-rest-ddd/domain/entity"
	"go-rest-ddd/internal/delivery/request"
	"go-rest-ddd/pkg/exceptions"
	"time"
)

type TransactionUseCase interface {
	ListTransaction(ctx context.Context, userId int, startDate, endDate time.Time, options *request.Option) ([]*entity.Transaction, int32, *exceptions.CustomerError)
}
