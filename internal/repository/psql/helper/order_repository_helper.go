package helper

import (
	"context"
	"gokomodo/domain/entity"
	"gokomodo/internal/repository/psql/mapper"
	"gokomodo/internal/repository/psql/models"

	"github.com/rocketlaunchr/dbq/v2"
)

func StoreOrder(ctx context.Context, E dbq.EFn, order *entity.Order) error {
	orderDbq := mapper.ToDbqStructOrder(order)

	stmt := dbq.INSERTStmt(models.Order{}.TableName(), models.TableOrder(), len(orderDbq), dbq.PostgreSQL)

	_, err := E(ctx, stmt, nil, orderDbq)
	if err != nil {
		return err
	}

	return nil
}
