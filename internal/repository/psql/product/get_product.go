package postgres_repository

import (
	"context"
	"database/sql"
	"fmt"
	"gokomodo/domain/entity"
	"gokomodo/internal/delivery/request"
	"gokomodo/internal/repository/psql/mapper"
	"gokomodo/internal/repository/psql/models"
	"gokomodo/pkg/common"

	"github.com/rocketlaunchr/dbq/v2"
)

func (repository *productRepository) getListProductQuery() string {
	var condition string

	stmt := fmt.Sprintf(`SELECT p.id, p.name, p.description, p.price, p.created_at, p.updated_at 
		FROM %s p %s`, models.Product{}.TableName(), condition)

	return stmt
}

func (repository *productRepository) GetProductSeller(ctx context.Context, userId common.ID, options *request.Option) ([]*entity.Product, error) {
	var args []interface{}

	opts := &dbq.Options{SingleResult: false, ConcreteStruct: models.Product{}, DecoderConfig: dbq.StdTimeConversionConfig()}

	stmt := repository.getListProductQuery()

	stmt += fmt.Sprintf(` WHERE p.seller_id = $1 LIMIT $2 OFFSET $3`)
	args = append(args, userId.String())
	args = append(args, options.Pagination.Limit)
	args = append(args, options.Pagination.Page)

	results := dbq.MustQ(ctx, repository.db, stmt, opts, args...)
	if results != nil {
		products := mapper.ToDomainListProduct(results.([]*models.Product))
		return products, nil
	} else {
		return nil, nil
	}
}

func (repository *productRepository) CountProductSeller(ctx context.Context, userId common.ID, options *request.Option) (int32, error) {
	var (
		args  []interface{}
		count sql.NullInt32
		err   error
	)

	stmt := repository.getListProductQuery()

	stmt = fmt.Sprintf(`SELECT count(id) from (%s WHERE p.seller_id = $1) as q`, stmt)
	args = append(args, userId.String())

	err = repository.db.QueryRowContext(ctx, stmt, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count.Int32, nil
}

func (repository *productRepository) GetProduct(ctx context.Context, options *request.Option) ([]*entity.Product, error) {
	var args []interface{}

	opts := &dbq.Options{SingleResult: false, ConcreteStruct: models.Product{}, DecoderConfig: dbq.StdTimeConversionConfig()}

	stmt := repository.getListProductQuery()

	stmt += fmt.Sprintf(` LIMIT $1 OFFSET $2`)
	args = append(args, options.Pagination.Limit)
	args = append(args, options.Pagination.Page)

	results := dbq.MustQ(ctx, repository.db, stmt, opts, args...)
	if results != nil {
		products := mapper.ToDomainListProduct(results.([]*models.Product))
		return products, nil
	} else {
		return nil, nil
	}
}

func (repository *productRepository) CountProduct(ctx context.Context, options *request.Option) (int32, error) {
	var (
		args  []interface{}
		count sql.NullInt32
		err   error
	)

	stmt := repository.getListProductQuery()

	stmt = fmt.Sprintf(`SELECT count(id) from (%s) as q`, stmt)

	err = repository.db.QueryRowContext(ctx, stmt, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count.Int32, nil
}
