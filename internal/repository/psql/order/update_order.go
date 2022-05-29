package postgres_repository

import (
	"context"
	"fmt"
	"gokomodo/domain/entity"
	"gokomodo/internal/repository/psql/models"

	"github.com/rocketlaunchr/dbq/v2"
)

func (repository *orderRepository) UpdateOrderStatus(ctx context.Context, order *entity.Order) error {
	var err error

	stmt := fmt.Sprintf(`UPDATE %s SET status = $2, updated_at = now() where id = $1`, models.Order{}.TableName())

	_, err = dbq.E(ctx, repository.db, stmt, nil, order.Id.String(), order.Status.GetValue())

	return err
}
