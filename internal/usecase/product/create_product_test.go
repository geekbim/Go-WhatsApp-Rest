package product_test

import (
	"context"
	"errors"
	"gokomodo/internal/mocks"
	product_usecase "gokomodo/internal/usecase/product"
	"gokomodo/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateProduct(t *testing.T) {
	ctx := context.TODO()

	productRepo := new(mocks.ProductRepositoryMock)

	productDTO := testdata.NewProductDTO()
	product := testdata.NewProduct(productDTO)

	productRepo.
		On("StoreProduct", mock.Anything, product).
		Return(nil)

	useCase := product_usecase.NewProductInteractor(productRepo)

	res, err := useCase.CreateProduct(ctx, product)

	assert.Nil(t, err)
	assert.Equal(t, product, res)
}

func TestCreateProductErr(t *testing.T) {
	ctx := context.TODO()

	productRepo := new(mocks.ProductRepositoryMock)

	productDTO := testdata.NewProductDTO()
	product := testdata.NewProduct(productDTO)

	err := errors.New("error create product")
	expectedErr := []error{
		err,
	}

	productRepo.
		On("StoreProduct", mock.Anything, product).
		Return(err)

	useCase := product_usecase.NewProductInteractor(productRepo)

	res, errUseCase := useCase.CreateProduct(ctx, product)

	assert.Nil(t, res)
	assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
}
