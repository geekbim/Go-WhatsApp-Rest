package postgres_repository

import (
	"database/sql"
	"gokomodo/domain/repository"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) repository.ProductRepository {
	return &productRepository{
		db: db,
	}
}
