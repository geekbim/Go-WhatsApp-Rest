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

func TestCreateOrder(t *testing.T) {
	ctx := context.TODO()

	productRepo := new(mocks.ProductRepositoryMock)
	orderRepo := new(mocks.OrderRepositoryMock)

	orderDTO := testdata.NewOrderDTO()
	order := testdata.NewOrder(orderDTO)

	productDTO := testdata.NewProductDTO()
	product := testdata.NewProduct(productDTO)

	productRepo.
		On("FindProductById", mock.Anything, order.Product.Id).
		Return(product, nil)
	orderRepo.
		On("StoreOrder", mock.Anything, order).
		Return(nil)

	useCase := order_usecase.NewOrderInteractor(orderRepo, productRepo)

	res, err := useCase.CreateOrder(ctx, order)

	assert.Nil(t, err)
	assert.Equal(t, order, res)
}

func TestCreateOrderErrStore(t *testing.T) {
	ctx := context.TODO()

	productRepo := new(mocks.ProductRepositoryMock)
	orderRepo := new(mocks.OrderRepositoryMock)

	orderDTO := testdata.NewOrderDTO()
	order := testdata.NewOrder(orderDTO)

	productDTO := testdata.NewProductDTO()
	product := testdata.NewProduct(productDTO)

	err := errors.New("store order failed")
	expectedErr := []error{
		err,
	}

	productRepo.
		On("FindProductById", mock.Anything, order.Product.Id).
		Return(product, nil)
	orderRepo.
		On("StoreOrder", mock.Anything, order).
		Return(err)

	useCase := order_usecase.NewOrderInteractor(orderRepo, productRepo)

	res, errUseCase := useCase.CreateOrder(ctx, order)

	assert.Nil(t, res)
	assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
}

func TestCreateOrderErrFind(t *testing.T) {
	ctx := context.TODO()

	productRepo := new(mocks.ProductRepositoryMock)
	orderRepo := new(mocks.OrderRepositoryMock)

	orderDTO := testdata.NewOrderDTO()
	order := testdata.NewOrder(orderDTO)

	productDTO := testdata.NewProductDTO()
	product := testdata.NewProduct(productDTO)

	err := errors.New("product not found")
	expectedErr := []error{
		err,
	}

	productRepo.
		On("FindProductById", mock.Anything, order.Product.Id).
		Return(product, err)
	orderRepo.
		On("StoreOrder", mock.Anything, order).
		Return(nil)

	useCase := order_usecase.NewOrderInteractor(orderRepo, productRepo)

	res, errUseCase := useCase.CreateOrder(ctx, order)

	assert.Nil(t, res)
	assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
}
