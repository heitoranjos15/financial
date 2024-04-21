package dto

import (
	"financial/src/core/expenses"
	"financial/src/core/incomes"
)

type Bank struct {
	Incomes  []incomes.Income
	Expenses []expenses.Expense
	Amount   float64
}

type BankRepository interface {
  Create(bank Bank) Bank
  AddIncome(bankId int, income incomes.Income) *Bank
  AddExpense(bankId int, expenses expenses.Expense) *Bank
  Get(bankId string) Bank
}
