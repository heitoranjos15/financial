package transactions

import (
	"financial/src/dto"
	"financial/src/model"
	"financial/src/repository"
)

type TransactionCore struct {
  repo repository.TransactionRepository
}

func New(repo repository.TransactionRepository) TransactionCore {
  return TransactionCore{
    repo: repo,
  }
}

func (t TransactionCore) CreateTransaction(data dto.CreateTransactionRequest) (model.Transaction, error) {
  transaction, err := t.repo.CreateTransaction(data)
  return transaction, err
}
