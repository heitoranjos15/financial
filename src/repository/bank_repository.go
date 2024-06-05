package repository

import (
	"context"
	"financial/src/model"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BankRepository interface {
	Create(amount float64) (model.Bank, error)
	Get(bankID int) (model.Bank, error)
	GetAll() ([]model.Bank, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewBankRepository(db *pgxpool.Pool) BankRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(amount float64) (model.Bank, error) {
	ctx := context.Background()
	bank := model.Bank{}

	err := r.db.QueryRow(
		ctx,
		"INSERT INTO bank (amount, status) VALUES ($1, 1) returning *",
		amount,
	).Scan(
		&bank.ID,
		&bank.Amount,
		&bank.CreatedAt,
		&bank.Status,
	)

	if err != nil {
		return model.Bank{}, err
	}

	return bank, nil
}

func (r *repository) Get(bankID int) (model.Bank, error) {
	ctx := context.Background()
	bank := model.Bank{}

	err := r.db.QueryRow(
		ctx,
		"SELECT * from bank where bank_id = $1",
		bankID,
	).Scan(
		&bank.ID,
		&bank.Amount,
		&bank.CreatedAt,
		&bank.Status,
	)

	if err != nil {
		return model.Bank{}, err
	}

	return bank, nil
}

func (r *repository) GetAll() ([]model.Bank, error) {
	ctx := context.Background()

	rows, err := r.db.Query(ctx, "SELECT * from bank;")

	if err != nil {
		fmt.Println(err)
		return []model.Bank{}, err
	}

  banks, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Bank])

	return banks, err
}
