package product

import (
	"context"
	"gokomodo/domain/entity"
	"gokomodo/pkg/exceptions"

	"github.com/hashicorp/go-multierror"
)

func (interactor *productInteractor) CreateProduct(ctx context.Context, product *entity.Product) (*entity.Product, *exceptions.CustomerError) {
	var multierr *multierror.Error

	errRepo := interactor.productRepository.StoreProduct(ctx, product)
	if errRepo != nil {
		multierr = multierror.Append(multierr, errRepo)
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multierr,
		}
	}

	return product, nil
}
