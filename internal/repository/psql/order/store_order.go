package postgres_repository

import (
	"context"
	"gokomodo/domain/entity"
	"gokomodo/internal/repository/psql/helper"

	"github.com/rocketlaunchr/dbq/v2"
)

func (repository *orderRepository) StoreOrder(ctx context.Context, order *entity.Order) error {
	var err error

	err = dbq.Tx(ctx, repository.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		err = helper.StoreOrder(ctx, E, order)
		if err != nil {
			return
		}

		txCommit()
	})

	return err
}
