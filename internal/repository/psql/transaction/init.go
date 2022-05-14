package postgres_repository

import (
	"database/sql"
	"majoo/domain/repository"
)

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) repository.TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}
