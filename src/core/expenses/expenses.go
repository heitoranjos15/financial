package expenses

import "time"

type Expense struct {
	Amount      float64
	Date        time.Time
	Category    string
	Description string
}

func SaveExpense(amount float64, date time.Time, category, description string) Expense {
	return Expense{
		Amount:      amount,
		Date:        date,
		Category:    category,
		Description: description,
	}
}
