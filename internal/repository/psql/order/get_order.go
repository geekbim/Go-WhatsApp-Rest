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

func (repository *orderRepository) getSellerListOrderQuery() string {
	var condition string

	stmt := fmt.Sprintf(`SELECT o.id, o.buyer_id as user_id, u.name as user_name, u.email as user_email, u.address as user_address, o.product_id, p.name as product_name, p.description as product_description, p.price as product_price, p.created_at as product_created_at, p.updated_at as product_updated_at, o.qty, o.total_price, o.status, o.created_at, o.updated_at 
		FROM %s o 
		LEFT JOIN %s u ON o.buyer_id = u.id 
		LEFT JOIN %s p ON o.product_id = p.id %s`, models.Order{}.TableName(), models.User{}.TableName(), models.Product{}.TableName(), condition)

	return stmt
}

func (repository *orderRepository) getBuyerListOrderQuery() string {
	var condition string

	stmt := fmt.Sprintf(`SELECT o.id, o.seller_id as user_id, u.name as user_name, u.email as user_email, u.address as user_address, o.product_id, p.name as product_name, p.description as product_description, p.price as product_price, p.created_at as product_created_at, p.updated_at as product_updated_at, o.qty, o.total_price, o.status, o.created_at, o.updated_at 
		FROM %s o 
		LEFT JOIN %s u ON o.seller_id = u.id 
		LEFT JOIN %s p ON o.product_id = p.id %s`, models.Order{}.TableName(), models.User{}.TableName(), models.Product{}.TableName(), condition)

	return stmt
}

func (repository *orderRepository) GetOrderSeller(ctx context.Context, userId common.ID, options *request.Option) ([]*entity.Order, error) {
	var args []interface{}

	opts := &dbq.Options{SingleResult: false, ConcreteStruct: models.OrderList{}, DecoderConfig: dbq.StdTimeConversionConfig()}

	stmt := repository.getSellerListOrderQuery()

	stmt += fmt.Sprintf(` WHERE o.seller_id = $1 LIMIT $2 OFFSET $3`)
	args = append(args, userId.String())
	args = append(args, options.Pagination.Limit)
	args = append(args, options.Pagination.Page)

	results := dbq.MustQ(ctx, repository.db, stmt, opts, args...)
	if results != nil {
		orders := mapper.ToDomainSellerListOrder(results.([]*models.OrderList))
		return orders, nil
	} else {
		return nil, nil
	}
}

func (repository *orderRepository) CountOrderSeller(ctx context.Context, userId common.ID, options *request.Option) (int32, error) {
	var (
		args  []interface{}
		count sql.NullInt32
		err   error
	)

	stmt := repository.getSellerListOrderQuery()

	stmt = fmt.Sprintf(`SELECT COUNT(id) from (%s WHERE o.seller_id = $1) as q`, stmt)
	args = append(args, userId.String())

	err = repository.db.QueryRowContext(ctx, stmt, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count.Int32, nil
}

func (repository *orderRepository) FindOrderBySellerIdAndOrderId(ctx context.Context, userId, orderId common.ID) (*entity.Order, error) {
	var args []interface{}

	opts := &dbq.Options{SingleResult: true, ConcreteStruct: models.OrderList{}, DecoderConfig: dbq.StdTimeConversionConfig()}

	stmt := repository.getSellerListOrderQuery()

	stmt += fmt.Sprintf(` WHERE o.seller_id = $1 AND o.id = $2 LIMIT 1`)
	args = append(args, userId.String())
	args = append(args, orderId.String())

	results := dbq.MustQ(ctx, repository.db, stmt, opts, args...)
	if results != nil {
		order := mapper.ToDomainSellerOrder(results.(*models.OrderList))
		return order, nil
	} else {
		return nil, nil
	}
}

func (repository *orderRepository) GetOrderBuyer(ctx context.Context, userId common.ID, options *request.Option) ([]*entity.Order, error) {
	var args []interface{}

	opts := &dbq.Options{SingleResult: false, ConcreteStruct: models.OrderList{}, DecoderConfig: dbq.StdTimeConversionConfig()}

	stmt := repository.getBuyerListOrderQuery()

	stmt += fmt.Sprintf(` WHERE o.buyer_id = $1 LIMIT $2 OFFSET $3`)
	args = append(args, userId.String())
	args = append(args, options.Pagination.Limit)
	args = append(args, options.Pagination.Page)

	results := dbq.MustQ(ctx, repository.db, stmt, opts, args...)
	if results != nil {
		orders := mapper.ToDomainBuyerListOrder(results.([]*models.OrderList))
		return orders, nil
	} else {
		return nil, nil
	}
}

func (repository *orderRepository) CountOrderBuyer(ctx context.Context, userId common.ID, options *request.Option) (int32, error) {
	var (
		args  []interface{}
		count sql.NullInt32
		err   error
	)

	stmt := repository.getBuyerListOrderQuery()

	stmt = fmt.Sprintf(`SELECT COUNT(id) from (%s WHERE o.buyer_id = $1) as q`, stmt)
	args = append(args, userId.String())

	err = repository.db.QueryRowContext(ctx, stmt, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count.Int32, nil
}
