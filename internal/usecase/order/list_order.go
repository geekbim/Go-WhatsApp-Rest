package order

import (
	"context"
	"gokomodo/domain/entity"
	"gokomodo/domain/valueobject"
	"gokomodo/internal/delivery/request"
	"gokomodo/pkg/common"
	"gokomodo/pkg/exceptions"

	"github.com/hashicorp/go-multierror"
)

func (interactor *orderInteractor) ListOrder(ctx context.Context, userId common.ID, role string, options *request.Option) ([]*entity.Order, int32, *exceptions.CustomerError) {
	var (
		orders   []*entity.Order
		total    int32
		errRepo  error
		multierr *multierror.Error
	)

	roleEnum, errRole := valueobject.NewRoleFromString(role)
	if errRole != nil {
		multierr = multierror.Append(multierr, errRole)
		return nil, int32(0), &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	if roleEnum == valueobject.USER_ROLE_SELLER {
		orders, errRepo = interactor.orderRepository.GetOrderSeller(ctx, userId, options)
		if errRepo != nil {
			multierr = multierror.Append(multierr, errRepo)
			return nil, int32(0), &exceptions.CustomerError{
				Status: exceptions.ERRBUSSINESS,
				Errors: multierr,
			}
		}

		total, errRepo = interactor.orderRepository.CountOrderSeller(ctx, userId, options)
		if errRepo != nil {
			multierr = multierror.Append(multierr, errRepo)
			return nil, int32(0), &exceptions.CustomerError{
				Status: exceptions.ERRBUSSINESS,
				Errors: multierr,
			}
		}
	}

	if roleEnum == valueobject.USER_ROLE_BUYER {
		orders, errRepo = interactor.orderRepository.GetOrderBuyer(ctx, userId, options)
		if errRepo != nil {
			multierr = multierror.Append(multierr, errRepo)
			return nil, int32(0), &exceptions.CustomerError{
				Status: exceptions.ERRBUSSINESS,
				Errors: multierr,
			}
		}

		total, errRepo = interactor.orderRepository.CountOrderBuyer(ctx, userId, options)
		if errRepo != nil {
			multierr = multierror.Append(multierr, errRepo)
			return nil, int32(0), &exceptions.CustomerError{
				Status: exceptions.ERRBUSSINESS,
				Errors: multierr,
			}
		}
	}

	return orders, total, nil
}
