package model

import "time"

type Cycle struct {
	ID          int
	TotalExpend float64
	StartDate   time.Time
	CreatedAt   time.Time
	Bank        Bank
	BankID      int
}
