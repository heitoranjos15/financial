package cycles

import (
	"financial/src/dto"
	"financial/src/model"
	"financial/src/repository"
)

type CycleCore struct {
  repo repository.CycleRepository
}

func New(repo repository.CycleRepository) CycleCore {
  return CycleCore{
    repo: repo,
  }
}

func (c CycleCore) CreateCycle(data dto.CreateCycleRequest) (model.Cycle, error) {
  cycle, err := c.repo.CreateCycle(data)
  return cycle, err
}
