package repository

import (
	"context"
	"gokomodo/domain/entity"
	"gokomodo/internal/delivery/request"
	"gokomodo/pkg/common"
)

type OrderRepository interface {
	GetOrderSeller(ctx context.Context, userId common.ID, options *request.Option) ([]*entity.Order, error)
	CountOrderSeller(ctx context.Context, userId common.ID, options *request.Option) (int32, error)
	FindOrderBySellerIdAndOrderId(ctx context.Context, userId, orderId common.ID) (*entity.Order, error)
	UpdateOrderStatus(ctx context.Context, order *entity.Order) error
	StoreOrder(ctx context.Context, order *entity.Order) error
	GetOrderBuyer(ctx context.Context, userId common.ID, options *request.Option) ([]*entity.Order, error)
	CountOrderBuyer(ctx context.Context, userId common.ID, options *request.Option) (int32, error)
}
