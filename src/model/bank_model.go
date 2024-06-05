package model

import "time"

type Bank struct {
	ID        int
	Amount    float64
	Status    int
	CreatedAt time.Time
}
