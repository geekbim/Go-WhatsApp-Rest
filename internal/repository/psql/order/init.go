package postgres_repository

import (
	"database/sql"
	"gokomodo/domain/repository"
)

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) repository.OrderRepository {
	return &orderRepository{
		db: db,
	}
}
