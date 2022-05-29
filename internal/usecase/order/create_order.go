package order

import (
	"context"
	"gokomodo/domain/entity"
	"gokomodo/pkg/exceptions"

	"github.com/hashicorp/go-multierror"
)

func (interactor *orderInteractor) CreateOrder(ctx context.Context, order *entity.Order) (*entity.Order, *exceptions.CustomerError) {
	var multierr *multierror.Error

	product, errFind := interactor.productRepository.FindProductById(ctx, order.Product.Id)
	if errFind != nil {
		multierr = multierror.Append(multierr, errFind)
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	order.SetSeller(product.Seller.Id)
	order.SetTotalPrice(order.Qty, product.Price)

	errStore := interactor.orderRepository.StoreOrder(ctx, order)
	if errStore != nil {
		multierr = multierror.Append(multierr, errStore)
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multierr,
		}
	}

	return order, nil
}
