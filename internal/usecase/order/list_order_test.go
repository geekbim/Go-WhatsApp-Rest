package order_test

import (
	"context"
	"errors"
	"gokomodo/domain/entity"
	"gokomodo/internal/mocks"
	order_usecase "gokomodo/internal/usecase/order"
	"gokomodo/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListOrder(t *testing.T) {
	ctx := context.TODO()

	orderRepo := new(mocks.OrderRepositoryMock)

	orderDTO := testdata.NewOrderDTO()
	order := testdata.NewOrder(orderDTO)
	orders := []*entity.Order{order}

	t.Run("ListOrderSeller", func(t *testing.T) {
		orderRepo.
			On("GetOrderSeller", mock.Anything, order.Seller.Id, mock.Anything).
			Return(orders, nil)
		orderRepo.
			On("CountOrderSeller", mock.Anything, order.Seller.Id, mock.Anything).
			Return(int32(len(orders)), nil)

		useCase := order_usecase.NewOrderInteractor(orderRepo, nil)

		res, count, err := useCase.ListOrder(ctx, order.Seller.Id, "SELLER", nil)

		assert.Nil(t, err)
		assert.Equal(t, orders, res)
		assert.Equal(t, int32(len(orders)), count)
	})

	t.Run("ListOrderBuyer", func(t *testing.T) {
		orderRepo.
			On("GetOrderBuyer", mock.Anything, order.Buyer.Id, mock.Anything).
			Return(orders, nil)
		orderRepo.
			On("CountOrderBuyer", mock.Anything, order.Buyer.Id, mock.Anything).
			Return(int32(len(orders)), nil)

		useCase := order_usecase.NewOrderInteractor(orderRepo, nil)

		res, count, err := useCase.ListOrder(ctx, order.Buyer.Id, "BUYER", nil)

		assert.Nil(t, err)
		assert.Equal(t, orders, res)
		assert.Equal(t, int32(len(orders)), count)
	})
}

func TestListOrderErrCount(t *testing.T) {
	ctx := context.TODO()

	orderRepo := new(mocks.OrderRepositoryMock)

	orderDTO := testdata.NewOrderDTO()
	order := testdata.NewOrder(orderDTO)
	orders := []*entity.Order{order}

	err := errors.New("error count order")
	expectedErr := []error{
		err,
	}

	t.Run("ListOrderSellerErrCount", func(t *testing.T) {
		orderRepo.
			On("GetOrderSeller", mock.Anything, order.Seller.Id, mock.Anything).
			Return(orders, nil)
		orderRepo.
			On("CountOrderSeller", mock.Anything, order.Seller.Id, mock.Anything).
			Return(int32(len(orders)), err)

		useCase := order_usecase.NewOrderInteractor(orderRepo, nil)

		res, count, errUseCase := useCase.ListOrder(ctx, order.Seller.Id, "SELLER", nil)

		assert.Nil(t, res)
		assert.Equal(t, int32(0), count)
		assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
	})

	t.Run("ListOrderBuyerErrCount", func(t *testing.T) {
		orderRepo.
			On("GetOrderBuyer", mock.Anything, order.Buyer.Id, mock.Anything).
			Return(orders, nil)
		orderRepo.
			On("CountOrderBuyer", mock.Anything, order.Buyer.Id, mock.Anything).
			Return(int32(len(orders)), err)

		useCase := order_usecase.NewOrderInteractor(orderRepo, nil)

		res, count, errUseCase := useCase.ListOrder(ctx, order.Buyer.Id, "BUYER", nil)

		assert.Nil(t, res)
		assert.Equal(t, int32(0), count)
		assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
	})
}

func TestListOrderErrGet(t *testing.T) {
	ctx := context.TODO()

	orderRepo := new(mocks.OrderRepositoryMock)

	orderDTO := testdata.NewOrderDTO()
	order := testdata.NewOrder(orderDTO)
	orders := []*entity.Order{order}

	err := errors.New("error get order")
	expectedErr := []error{
		err,
	}

	t.Run("ListOrderSellerErrCount", func(t *testing.T) {
		orderRepo.
			On("GetOrderSeller", mock.Anything, order.Seller.Id, mock.Anything).
			Return(orders, err)
		orderRepo.
			On("CountOrderSeller", mock.Anything, order.Seller.Id, mock.Anything).
			Return(int32(len(orders)), nil)

		useCase := order_usecase.NewOrderInteractor(orderRepo, nil)

		res, count, errUseCase := useCase.ListOrder(ctx, order.Seller.Id, "SELLER", nil)

		assert.Nil(t, res)
		assert.Equal(t, int32(0), count)
		assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
	})

	t.Run("ListOrderBuyerErrCount", func(t *testing.T) {
		orderRepo.
			On("GetOrderBuyer", mock.Anything, order.Buyer.Id, mock.Anything).
			Return(orders, err)
		orderRepo.
			On("CountOrderBuyer", mock.Anything, order.Buyer.Id, mock.Anything).
			Return(int32(len(orders)), nil)

		useCase := order_usecase.NewOrderInteractor(orderRepo, nil)

		res, count, errUseCase := useCase.ListOrder(ctx, order.Buyer.Id, "BUYER", nil)

		assert.Nil(t, res)
		assert.Equal(t, int32(0), count)
		assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
	})
}

func TestListOrderErrRole(t *testing.T) {
	ctx := context.TODO()

	orderDTO := testdata.NewOrderDTO()
	order := testdata.NewOrder(orderDTO)

	err := errors.New("role not found")
	expectedErr := []error{
		err,
	}

	useCase := order_usecase.NewOrderInteractor(nil, nil)

	res, count, errUseCase := useCase.ListOrder(ctx, order.Seller.Id, "GUEST", nil)

	assert.Nil(t, res)
	assert.Equal(t, int32(0), count)
	assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
}
