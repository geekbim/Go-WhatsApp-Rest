package postgres_repository

import (
	"context"
	"fmt"
	"gokomodo/domain/entity"
	"gokomodo/internal/repository/psql/mapper"
	"gokomodo/internal/repository/psql/models"
	"gokomodo/pkg/common"

	"github.com/rocketlaunchr/dbq/v2"
)

func (repository *productRepository) FindProductById(ctx context.Context, productId common.ID) (*entity.Product, error) {
	opts := &dbq.Options{SingleResult: true, ConcreteStruct: models.Product{}, DecoderConfig: dbq.StdTimeConversionConfig()}

	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, models.Product{}.TableName())
	results := dbq.MustQ(ctx, repository.db, stmt, opts, productId.String())
	if results != nil {
		product := mapper.ToDomainProduct(results.(*models.Product))
		return product, nil
	} else {
		return nil, nil
	}
}
