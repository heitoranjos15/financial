package dto

import (
	"time"
)

type CreateTransactionRequest struct {
	Amount      float64
	Date        time.Time
	Category    string
	Description string
	Type        string
	CycleID     int
}
