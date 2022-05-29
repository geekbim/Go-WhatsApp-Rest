package order_test

import (
	"context"
	"errors"
	"gokomodo/internal/mocks"
	order_usecase "gokomodo/internal/usecase/order"
	"gokomodo/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAcceptOrder(t *testing.T) {
	ctx := context.TODO()

	orderRepo := new(mocks.OrderRepositoryMock)

	orderDTO := testdata.NewOrderDTO()
	order := testdata.NewOrder(orderDTO)

	orderRepo.
		On("FindOrderBySellerIdAndOrderId", mock.Anything, order.Seller.Id, order.Id).
		Return(order, nil)
	orderRepo.
		On("UpdateOrderStatus", mock.Anything, order).
		Return(nil)

	useCase := order_usecase.NewOrderInteractor(orderRepo, nil)

	res, err := useCase.AcceptOrder(ctx, order.Seller.Id, order.Id)

	assert.Nil(t, err)
	assert.Equal(t, order, res)
}

func TestAcceptOrderErrUpdate(t *testing.T) {
	ctx := context.TODO()

	orderRepo := new(mocks.OrderRepositoryMock)

	orderDTO := testdata.NewOrderDTO()
	order := testdata.NewOrder(orderDTO)

	err := errors.New("update order status failed")
	expectedErr := []error{
		err,
	}

	orderRepo.
		On("FindOrderBySellerIdAndOrderId", mock.Anything, order.Seller.Id, order.Id).
		Return(order, nil)
	orderRepo.
		On("UpdateOrderStatus", mock.Anything, order).
		Return(err)

	useCase := order_usecase.NewOrderInteractor(orderRepo, nil)

	res, errUseCase := useCase.AcceptOrder(ctx, order.Seller.Id, order.Id)

	assert.Nil(t, res)
	assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
}

func TestAcceptOrderErrFind(t *testing.T) {
	ctx := context.TODO()

	orderRepo := new(mocks.OrderRepositoryMock)

	orderDTO := testdata.NewOrderDTO()
	order := testdata.NewOrder(orderDTO)

	err := errors.New("order not found")
	expectedErr := []error{
		err,
	}

	orderRepo.
		On("FindOrderBySellerIdAndOrderId", mock.Anything, order.Seller.Id, order.Id).
		Return(order, err)
	orderRepo.
		On("UpdateOrderStatus", mock.Anything, order).
		Return(nil)

	useCase := order_usecase.NewOrderInteractor(orderRepo, nil)

	res, errUseCase := useCase.AcceptOrder(ctx, order.Seller.Id, order.Id)

	assert.Nil(t, res)
	assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
}
