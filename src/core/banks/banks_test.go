package banks_test

import (
	"financial/src/core/banks"
	"financial/src/core/expenses"
	"financial/src/core/incomes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewBank(t *testing.T) {
  bank := banks.NewBank(200)
  assert.Equal(t, bank.Amount, float64(200))
}

func TestBankAddExpense(t *testing.T) {
  bank := banks.NewBank(100)
  expense := expenses.SaveExpense(2, time.Now(), "market", "no")
  bank.AddExpense(expense)
  assert.Equal(t, bank.Amount, float64(98))
  assert.Equal(t, bank.Expenses, []expenses.Expense{expense})
}

func TestBankAddIncome(t *testing.T) {
  bank := banks.NewBank(100)
  income := incomes.SaveIncome(2, time.Now(), "salary", "no")
  bank.AddIncome(income)
  assert.Equal(t, bank.Amount, float64(102))
  assert.Equal(t, bank.Incomes, []incomes.Income{income})
}
