package postgres_repository

import (
	"context"
	"gokomodo/domain/entity"
	"gokomodo/internal/repository/psql/helper"

	"github.com/rocketlaunchr/dbq/v2"
)

func (repository *productRepository) StoreProduct(ctx context.Context, product *entity.Product) error {
	var err error

	err = dbq.Tx(ctx, repository.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		err = helper.StoreProduct(ctx, E, product)
		if err != nil {
			return
		}

		txCommit()
	})

	return err
}
