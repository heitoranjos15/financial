package banks

import (
	"financial/src/core/expenses"
	"financial/src/core/incomes"
	// "time"
)

type Bank struct {
	Incomes  []incomes.Income
	Expenses []expenses.Expense
	Amount   float64
}

func NewBank(initalAmount float64) Bank {
	return Bank{
		Amount: initalAmount,
	}
}

func (b *Bank) AddIncome(income incomes.Income) {
	b.Incomes = append(b.Incomes, income)
	b.Amount += income.Amount
}

func (b *Bank) AddExpense(expense expenses.Expense) {
	b.Expenses = append(b.Expenses, expense)
	b.Amount -= expense.Amount
}

// func (b Bank) GetCurrentMontBalance() {
//   month := time.Now().Month();
//
// }
