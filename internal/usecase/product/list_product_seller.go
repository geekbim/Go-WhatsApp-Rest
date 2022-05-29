package product

import (
	"context"
	"gokomodo/domain/entity"
	"gokomodo/internal/delivery/request"
	"gokomodo/pkg/common"
	"gokomodo/pkg/exceptions"

	"github.com/hashicorp/go-multierror"
)

func (interactor *productInteractor) ListProductSeller(ctx context.Context, userId common.ID, options *request.Option) ([]*entity.Product, int32, *exceptions.CustomerError) {
	var multierr *multierror.Error

	products, errRepo := interactor.productRepository.GetProductSeller(ctx, userId, options)
	if errRepo != nil {
		multierr = multierror.Append(multierr, errRepo)
		return nil, int32(0), &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	total, errRepo := interactor.productRepository.CountProductSeller(ctx, userId, options)
	if errRepo != nil {
		multierr = multierror.Append(multierr, errRepo)
		return nil, int32(0), &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	return products, total, nil
}
