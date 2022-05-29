package order

import (
	"context"
	"gokomodo/domain/entity"
	"gokomodo/pkg/common"
	"gokomodo/pkg/exceptions"

	"github.com/hashicorp/go-multierror"
)

func (interactor *orderInteractor) AcceptOrder(ctx context.Context, userId, orderId common.ID) (*entity.Order, *exceptions.CustomerError) {
	var multierr *multierror.Error

	order, errRepo := interactor.orderRepository.FindOrderBySellerIdAndOrderId(ctx, userId, orderId)
	if errRepo != nil {
		multierr = multierror.Append(multierr, errRepo)
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	order.SetStatusAccepted()

	errRepo = interactor.orderRepository.UpdateOrderStatus(ctx, order)
	if errRepo != nil {
		multierr = multierror.Append(multierr, errRepo)
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multierr,
		}
	}

	return order, nil
}
