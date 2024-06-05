package banks

import (
	"financial/src/model"
	"financial/src/repository"
)

type BankCore struct {
	repository repository.BankRepository
}

func New(repository repository.BankRepository) BankCore {
	return BankCore{
		repository: repository,
	}
}

func (b BankCore) CreateBank(initialAmount float64) (model.Bank, error) {
	bank, err := b.repository.Create(initialAmount)
	return bank, err
}

func (b BankCore) GetAllBanks() ([]model.Bank, error) {
  banks, err := b.repository.GetAll()
  return banks, err
}
