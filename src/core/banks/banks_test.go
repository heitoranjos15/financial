package banks_test

import (
	"errors"
	"testing"

	"financial/mocks"
	"financial/src/core/banks"
	"financial/src/model"

	"github.com/stretchr/testify/assert"
)

func TestCreateBank(t *testing.T) {
	repo := mocks.NewBankRepository(t)
  repo.On("Create", 0.0).Return(model.Bank{ID: 1}, nil)

	bank, _ := banks.New(repo).CreateBank(0)
	assert.Equal(t, bank.ID, 1)
}

func TestCreateBank_fails(t *testing.T) {
	repo := mocks.NewBankRepository(t)
  repo.On("Create", 0.0).Return(model.Bank{}, errors.New("herr"))

	_, error := banks.New(repo).CreateBank(0)
  assert.Error(t, error)
}

func TestGetAllBanks(t *testing.T) {
	repo := mocks.NewBankRepository(t)
  repo.On("GetAll").Return([]model.Bank{{ID: 0}}, nil)

	banks, error := banks.New(repo).GetAllBanks()
  assert.Nil(t, error)
  assert.Equal(t, len(banks), 1)
}
