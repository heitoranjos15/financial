package pages

import (
	"financial/src/core/banks"
	"financial/src/repository"
	"financial/tui/tools"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Actions interface {
	View() string
  NormalModeActions(key string)
}

type DefaultPage struct {
	Name    string
	Actions Actions
}

type bankPage struct {
	service banks.BankCore
	Page    DefaultPage
	Tools   *tools.Tools
}

func NewBankPage(db *pgxpool.Pool) *bankPage {
	repo := repository.NewBankRepository(db)
	service := banks.New(repo)

	banks, _ := service.GetAllBanks()
	var fields []tools.Field
	for _, bank := range banks {
		fields = append(fields, tools.Field{Input: strconv.FormatFloat(bank.Amount, 'f', -1, 64), Value: ""})
	}

	bank := &bankPage{
		service: service,
    Tools: &tools.Tools{
      Fields: fields,
    },
	}
	return bank
}

func (b *bankPage) NormalModeActions(key string) {
  switch key {
  case "up", ",":
    if b.Tools.GetCursor() > 0 {
      b.Tools.SetCursor(b.Tools.GetCursor() - 1)
    }
  default:
    b.Tools.CursorActions(key)
  }
}

func (b *bankPage) View() string {
	var s string
	for i, field := range b.Tools.GetFields() {
		cursorStr := " "
		if b.Tools.GetCursor() == i {
			cursorStr = ">"
		}
		s += fmt.Sprintf("%s [bank %d] : %s %d\n", cursorStr, i, field.GetInput(), b.Tools.GetCursor())
	}

	return s
}
