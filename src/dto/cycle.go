package dto

import "time"

type CreateCycleRequest struct {
	BankID    int
	StartDate time.Time
}
