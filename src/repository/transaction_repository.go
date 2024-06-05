package repository

import (
	"context"
	"financial/src/dto"
	"financial/src/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionRepository interface {
	CreateTransaction(data dto.CreateTransactionRequest) (model.Transaction, error)
}

func NewTransactionRepository(db *pgxpool.Pool) TransactionRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateTransaction(data dto.CreateTransactionRequest) (model.Transaction, error) {
	ctx := context.Background()

	transaction := model.Transaction{}

	err := r.db.QueryRow(
		ctx,
		"INSERT INTO transactions (amount, cycle_id, transaction_type, category, description) ($1, $1, $1, $1, $1) returning *",
		data.Amount,
		data.CycleID,
		data.Type,
		data.Category,
		data.Description,
	).Scan(
		&transaction.ID,
		&transaction.Amount,
		&transaction.Type,
		&transaction.Category,
		&transaction.Description,
	)

	if err != nil {
		return model.Transaction{}, err
	}

	return transaction, nil
}
