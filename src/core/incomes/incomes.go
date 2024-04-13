package incomes

import "time"


type Income struct {
  Amount float64
  Date time.Time
  Category string
  Description string
}

func SaveIncome(amount float64, date time.Time, category, description string) Income {
  return Income{
    Amount: amount,
    Date: date,
    Category: category,
    Description: description,
  }
}
