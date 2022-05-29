package product

import (
	"context"
	"gokomodo/domain/entity"
	"gokomodo/internal/delivery/request"
	"gokomodo/pkg/exceptions"

	"github.com/hashicorp/go-multierror"
)

func (interactor *productInteractor) ListProductBuyer(ctx context.Context, options *request.Option) ([]*entity.Product, int32, *exceptions.CustomerError) {
	var multierr *multierror.Error

	products, errRepo := interactor.productRepository.GetProduct(ctx, options)
	if errRepo != nil {
		multierr = multierror.Append(multierr, errRepo)
		return nil, int32(0), &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	total, errRepo := interactor.productRepository.CountProduct(ctx, options)
	if errRepo != nil {
		multierr = multierror.Append(multierr, errRepo)
		return nil, int32(0), &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	return products, total, nil
}
