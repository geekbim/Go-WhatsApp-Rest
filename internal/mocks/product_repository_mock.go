package mocks

import (
	"context"
	"gokomodo/domain/entity"
	"gokomodo/internal/delivery/request"
	"gokomodo/pkg/common"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (repository *ProductRepositoryMock) GetProductSeller(ctx context.Context, userId common.ID, options *request.Option) ([]*entity.Product, error) {
	args := repository.Called(ctx, userId, options)
	return args.Get(0).([]*entity.Product), args.Error(1)
}

func (repository *ProductRepositoryMock) CountProductSeller(ctx context.Context, userId common.ID, options *request.Option) (int32, error) {
	args := repository.Called(ctx, userId, options)
	return args.Get(0).(int32), args.Error(1)
}

func (repository *ProductRepositoryMock) StoreProduct(ctx context.Context, product *entity.Product) error {
	args := repository.Called(ctx, product)
	return args.Error(0)
}

func (repository *ProductRepositoryMock) GetProduct(ctx context.Context, options *request.Option) ([]*entity.Product, error) {
	args := repository.Called(ctx, options)
	return args.Get(0).([]*entity.Product), args.Error(1)
}

func (repository *ProductRepositoryMock) CountProduct(ctx context.Context, options *request.Option) (int32, error) {
	args := repository.Called(ctx, options)
	return args.Get(0).(int32), args.Error(1)
}

func (repository *ProductRepositoryMock) FindProductById(ctx context.Context, productId common.ID) (*entity.Product, error) {
	args := repository.Called(ctx, productId)
	return args.Get(0).(*entity.Product), args.Error(1)
}
