package usecase

import (
	"context"
	"gokomodo/domain/entity"
	"gokomodo/internal/delivery/request"
	"gokomodo/pkg/common"
	"gokomodo/pkg/exceptions"
)

type OrderUseCase interface {
	ListOrder(ctx context.Context, userId common.ID, role string, options *request.Option) ([]*entity.Order, int32, *exceptions.CustomerError)
	AcceptOrder(ctx context.Context, userId, orderId common.ID) (*entity.Order, *exceptions.CustomerError)
	CreateOrder(ctx context.Context, order *entity.Order) (*entity.Order, *exceptions.CustomerError)
}
