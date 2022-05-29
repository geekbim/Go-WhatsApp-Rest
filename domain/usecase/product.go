package usecase

import (
	"context"
	"gokomodo/domain/entity"
	"gokomodo/internal/delivery/request"
	"gokomodo/pkg/common"
	"gokomodo/pkg/exceptions"
)

type ProductUseCase interface {
	ListProductSeller(ctx context.Context, userId common.ID, options *request.Option) ([]*entity.Product, int32, *exceptions.CustomerError)
	CreateProduct(ctx context.Context, product *entity.Product) (*entity.Product, *exceptions.CustomerError)
	ListProductBuyer(ctx context.Context, options *request.Option) ([]*entity.Product, int32, *exceptions.CustomerError)
}
