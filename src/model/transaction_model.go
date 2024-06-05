package model

import (
	"time"
)

type Transaction struct {
	ID          int
	Amount      float64
	Date        time.Time
	Category    string
	Description string
	Type        string
	CycleID     int
	CreatedAt   time.Time
}
