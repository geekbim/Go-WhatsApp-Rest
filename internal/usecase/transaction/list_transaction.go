package transaction

import (
	"context"
	"majoo/domain/entity"
	"majoo/internal/delivery/request"
	"majoo/pkg/exceptions"
	"time"

	"github.com/hashicorp/go-multierror"
)

func (interactor *transactionInteractor) ListTransaction(ctx context.Context, userId int, startDate, endDate time.Time, options *request.Option) ([]*entity.Transaction, int32, *exceptions.CustomerError) {
	var multierr *multierror.Error

	transactions, err := interactor.transactionRepository.GetTransactionByUserIdAndDate(ctx, userId, startDate, endDate, options)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		return nil, 0, &exceptions.CustomerError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multierr,
		}
	}

	total, err := interactor.transactionRepository.CountTransactionByUserIdAndDate(ctx, userId, startDate, endDate, options)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		return nil, 0, &exceptions.CustomerError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multierr,
		}
	}

	return transactions, total, nil
}
