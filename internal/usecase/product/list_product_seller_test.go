package product_test

import (
	"context"
	"errors"
	"gokomodo/domain/entity"
	"gokomodo/internal/mocks"
	product_usecase "gokomodo/internal/usecase/product"
	"gokomodo/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListProductSeller(t *testing.T) {
	ctx := context.TODO()

	productRepo := new(mocks.ProductRepositoryMock)

	productDTO := testdata.NewProductDTO()
	product := testdata.NewProduct(productDTO)
	products := []*entity.Product{product}

	productRepo.
		On("GetProductSeller", mock.Anything, product.Seller.Id, mock.Anything).
		Return(products, nil)
	productRepo.
		On("CountProductSeller", mock.Anything, product.Seller.Id, mock.Anything).
		Return(int32(len(products)), nil)

	useCase := product_usecase.NewProductInteractor(productRepo)

	res, count, err := useCase.ListProductSeller(ctx, product.Seller.Id, nil)

	assert.Nil(t, err)
	assert.Equal(t, products, res)
	assert.Equal(t, int32(len(products)), count)
}

func TestListProductSellerErrCount(t *testing.T) {
	ctx := context.TODO()

	productRepo := new(mocks.ProductRepositoryMock)

	productDTO := testdata.NewProductDTO()
	product := testdata.NewProduct(productDTO)
	products := []*entity.Product{product}

	err := errors.New("error count product")
	expectedErr := []error{
		err,
	}

	productRepo.
		On("GetProductSeller", mock.Anything, product.Seller.Id, mock.Anything).
		Return(products, nil)
	productRepo.
		On("CountProductSeller", mock.Anything, product.Seller.Id, mock.Anything).
		Return(int32(0), err)

	productRepo.
		On("StoreProduct", mock.Anything, product).
		Return(err)

	useCase := product_usecase.NewProductInteractor(productRepo)

	res, count, errUseCase := useCase.ListProductSeller(ctx, product.Seller.Id, nil)

	assert.Nil(t, res)
	assert.Equal(t, int32(0), count)
	assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
}

func TestListProductSellerErrGet(t *testing.T) {
	ctx := context.TODO()

	productRepo := new(mocks.ProductRepositoryMock)

	productDTO := testdata.NewProductDTO()
	product := testdata.NewProduct(productDTO)
	products := []*entity.Product{product}

	err := errors.New("error get product")
	expectedErr := []error{
		err,
	}

	productRepo.
		On("GetProductSeller", mock.Anything, product.Seller.Id, mock.Anything).
		Return(products, err)
	productRepo.
		On("CountProductSeller", mock.Anything, product.Seller.Id, mock.Anything).
		Return(int32(0), nil)

	productRepo.
		On("StoreProduct", mock.Anything, product).
		Return(err)

	useCase := product_usecase.NewProductInteractor(productRepo)

	res, count, errUseCase := useCase.ListProductSeller(ctx, product.Seller.Id, nil)

	assert.Nil(t, res)
	assert.Equal(t, int32(0), count)
	assert.Equal(t, expectedErr, errUseCase.Errors.Errors)
}
