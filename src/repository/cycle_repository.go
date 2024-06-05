package repository

import (
	"context"

	"financial/src/dto"
	"financial/src/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CycleRepository interface {
	CreateCycle(cycle dto.CreateCycleRequest) (model.Cycle, error)
}

func NewCycleRepository(db *pgxpool.Pool) BankRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateCycle(data dto.CreateCycleRequest) (model.Cycle, error) {
	ctx := context.Background()

	cycle := model.Cycle{}

	err := r.db.QueryRow(
		ctx,
		"INSERT INTO cycles (total_spent, start_date, bank_id) (0, $1, $1) returning *",
		data.StartDate,
		data.BankID,
	).Scan(
		&cycle.ID,
		&cycle.StartDate,
		&cycle.BankID,
	)

	if err != nil {
		return model.Cycle{}, err
	}

	return cycle, nil
}
