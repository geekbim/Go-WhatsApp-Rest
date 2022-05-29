package mocks

import (
	"context"
	"gokomodo/domain/entity"
	"gokomodo/internal/delivery/request"
	"gokomodo/pkg/common"

	"github.com/stretchr/testify/mock"
)

type OrderRepositoryMock struct {
	mock.Mock
}

func (repository *OrderRepositoryMock) GetOrderSeller(ctx context.Context, userId common.ID, options *request.Option) ([]*entity.Order, error) {
	args := repository.Called(ctx, userId, options)
	return args.Get(0).([]*entity.Order), args.Error(1)
}

func (repository *OrderRepositoryMock) CountOrderSeller(ctx context.Context, userId common.ID, options *request.Option) (int32, error) {
	args := repository.Called(ctx, userId, options)
	return args.Get(0).(int32), args.Error(1)
}

func (repository *OrderRepositoryMock) FindOrderBySellerIdAndOrderId(ctx context.Context, userId, orderId common.ID) (*entity.Order, error) {
	args := repository.Called(ctx, userId, orderId)
	return args.Get(0).(*entity.Order), args.Error(1)
}

func (repository *OrderRepositoryMock) UpdateOrderStatus(ctx context.Context, order *entity.Order) error {
	args := repository.Called(ctx, order)
	return args.Error(0)
}

func (repository *OrderRepositoryMock) StoreOrder(ctx context.Context, order *entity.Order) error {
	args := repository.Called(ctx, order)
	return args.Error(0)
}

func (repository *OrderRepositoryMock) GetOrderBuyer(ctx context.Context, userId common.ID, options *request.Option) ([]*entity.Order, error) {
	args := repository.Called(ctx, userId, options)
	return args.Get(0).([]*entity.Order), args.Error(1)
}

func (repository *OrderRepositoryMock) CountOrderBuyer(ctx context.Context, userId common.ID, options *request.Option) (int32, error) {
	args := repository.Called(ctx, userId, options)
	return args.Get(0).(int32), args.Error(1)
}
