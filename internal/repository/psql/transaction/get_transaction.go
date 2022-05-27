package postgres_repository

import (
	"context"
	"database/sql"
	"fmt"
	"go-rest-ddd/domain/entity"
	"go-rest-ddd/internal/delivery/request"
	"go-rest-ddd/internal/repository/psql/mapper"
	"go-rest-ddd/internal/repository/psql/models"
	"time"

	"github.com/rocketlaunchr/dbq/v2"
)

func (repository *transactionRepository) getListTransactionQuery(userId int, startDate, endDate time.Time, option *request.Option) (string, []interface{}) {
	var args []interface{}
	args = append(args, userId)
	args = append(args, startDate)
	args = append(args, endDate)
	condition := ""

	stmt := fmt.Sprintf(`SELECT t.id, m.merchant_name, o.outlet_name, t.bill_total as omzet, t.created_at, t.updated_at 
		FROM %s t 
		JOIN merchants m ON m.id = t.merchant_id 
		JOIN outlets o ON o.id = t.outlet_id 
		WHERE m.user_id = $1 
		AND t.created_at BETWEEN DATE($2) AND DATE($3) %s`, models.Transaction{}.TableName(), condition)

	return stmt, args
}

func (repository *transactionRepository) GetTransactionByUserIdAndDate(ctx context.Context, userId int, startDate, endDate time.Time, options *request.Option) ([]*entity.Transaction, error) {
	opts := &dbq.Options{SingleResult: false, ConcreteStruct: models.Transaction{}, DecoderConfig: dbq.StdTimeConversionConfig()}

	stmt, args := repository.getListTransactionQuery(userId, startDate, endDate, options)

	stmt += fmt.Sprintf(` LIMIT $%d OFFSET $%d`, len(args)+1, len(args)+2)
	args = append(args, options.Pagination.Limit)
	args = append(args, options.Pagination.Page)

	results := dbq.MustQ(ctx, repository.db, stmt, opts, args...)
	if results != nil {
		transactions := mapper.ToDomainListTransaction(results.([]*models.Transaction))
		return transactions, nil
	} else {
		return nil, nil
	}
}

func (repository *transactionRepository) CountTransactionByUserIdAndDate(ctx context.Context, userId int, startDate, endDate time.Time, options *request.Option) (int32, error) {
	var (
		count sql.NullInt32
		err   error
	)

	stmt, args := repository.getListTransactionQuery(userId, startDate, endDate, options)

	stmt = fmt.Sprintf(`SELECT count(id) from (%s) as q`, stmt)

	err = repository.db.QueryRowContext(ctx, stmt, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count.Int32, nil
}
