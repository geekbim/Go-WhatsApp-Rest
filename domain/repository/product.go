package repository

import (
	"context"
	"gokomodo/domain/entity"
	"gokomodo/internal/delivery/request"
	"gokomodo/pkg/common"
)

type ProductRepository interface {
	GetProductSeller(ctx context.Context, userId common.ID, options *request.Option) ([]*entity.Product, error)
	CountProductSeller(ctx context.Context, userId common.ID, options *request.Option) (int32, error)
	StoreProduct(ctx context.Context, product *entity.Product) error
	GetProduct(ctx context.Context, options *request.Option) ([]*entity.Product, error)
	CountProduct(ctx context.Context, options *request.Option) (int32, error)
	FindProductById(ctx context.Context, productId common.ID) (*entity.Product, error)
}
