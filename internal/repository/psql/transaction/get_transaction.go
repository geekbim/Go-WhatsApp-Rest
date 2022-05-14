package postgres_repository

import (
	"context"
	"fmt"
	"majoo/domain/entity"
	"majoo/internal/delivery/request"
	"majoo/internal/repository/psql/mapper"
	"majoo/internal/repository/psql/models"
	"time"

	"github.com/rocketlaunchr/dbq/v2"
)

func (repository *transactionRepository) getListTransactionQuery(userId int, startDate, endDate time.Time, option *request.Option) (string, []interface{}) {
	var args []interface{}
	args = append(args, userId)
	args = append(args, startDate)
	args = append(args, endDate)
	condition := ""

	stmt := fmt.Sprintf(`SELECT t.id, m.merchant_name, o.outlet_name, t.bill_total as omzet, t.created_at, t.created_by, t.updated_at, t.updated_by 
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
