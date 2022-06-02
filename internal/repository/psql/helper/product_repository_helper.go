package helper

import (
	"context"
	"database/sql"
	"fmt"
	"gokomodo/domain/entity"
	"gokomodo/internal/repository/psql/mapper"
	"gokomodo/internal/repository/psql/models"
	"gokomodo/pkg/common"

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

func FindProductById(ctx context.Context, db *sql.DB, productId common.ID) (*entity.Product, error) {
	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, models.Product{}.TableName())

	opts := &dbq.Options{DecoderConfig: dbq.StdTimeConversionConfig()}

	results, err := dbq.Qs(ctx, db, stmt, models.Product{}, opts, productId.String())
	if err != nil {
		return nil, err
	} else {
		return mapper.ToDomainProduct(results.(*models.Product)), err
	}
}
