package postgres_repository

import (
	"database/sql"
	"go-rest-ddd/domain/repository"
)

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) repository.TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}
