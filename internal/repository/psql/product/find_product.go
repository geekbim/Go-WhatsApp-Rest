package postgres_repository

import (
	"context"
	"gokomodo/domain/entity"
	"gokomodo/internal/repository/psql/helper"
	"gokomodo/pkg/common"
)

func (repository *productRepository) FindProductById(ctx context.Context, productId common.ID) (*entity.Product, error) {
	product, err := helper.FindProductById(ctx, repository.db, productId)
	if err != nil {
		return nil, err
	}

	return product, nil
}
