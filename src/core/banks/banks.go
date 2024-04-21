package banks

import (
	"financial/src/core/expenses"
	"financial/src/core/incomes"
	"financial/src/dto"
	// "time"
)

type Bank struct {
	Id       int
	Incomes  []incomes.Income
	Expenses []expenses.Expense
	Amount   float64
}

type BankCore struct {
	repository dto.BankRepository
}

func (b BankCore) NewBank(repository dto.BankRepository) BankCore {
	return BankCore{
		repository: repository,
	}
}

func (b BankCore) AddIncome(bank *Bank, income incomes.Income) {
	bank.Incomes = append(bank.Incomes, income)
	bank.Amount += income.Amount
  b.repository.AddIncome(bank.Id, income)
}

func (b BankCore) AddExpense(bank *Bank, expense expenses.Expense) {
	bank.Expenses = append(bank.Expenses, expense)
	bank.Amount -= expense.Amount
  b.repository.AddExpense(bank.Id, expense)
}

// func (b Bank) GetCurrentMontBalance() {
//   month := time.Now().Month();
//
// }
