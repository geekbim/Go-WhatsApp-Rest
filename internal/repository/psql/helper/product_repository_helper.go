package helper

import (
	"context"
	"gokomodo/domain/entity"
	"gokomodo/internal/repository/psql/mapper"
	"gokomodo/internal/repository/psql/models"

	"github.com/rocketlaunchr/dbq/v2"
)

func StoreProduct(ctx context.Context, E dbq.EFn, product *entity.Product) error {
	productDbq := mapper.ToDbqStructProduct(product)

	stmt := dbq.INSERTStmt(models.Product{}.TableName(), models.TableProduct(), len(productDbq), dbq.PostgreSQL)

	_, err := E(ctx, stmt, nil, productDbq)
	if err != nil {
		return err
	}

	return nil
}
