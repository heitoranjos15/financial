package view

import (
	"context"
	"fmt"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"

	"financial/adapters/postgres"
	"financial/src/core/banks"
	"financial/src/repository"
)

const (
	TEXTMODE   = "text"
	NORMALMODE = "normal"
)

type TeaModel struct {
	fields     []field
	cursor     int
	selected   map[int]struct{}
	mode       string
	services   core
	errorMsg   string
	actualPage string
}

type field struct {
	name  string
	value string
}

type core struct {
	bankCore banks.BankCore
}

func InitialModel() TeaModel {
	ctx := context.Background()
	conn := postgres.GetConnection(ctx)

	postgres.RunMigrations()

	bankRepo := repository.NewBankRepository(conn)
	bankCore := banks.New(bankRepo)

	services := core{
		bankCore: bankCore,
	}

	return TeaModel{
		fields: []field{{ name: "home", value: "home"}},

		selected:   make(map[int]struct{}),
		mode:       NORMALMODE,
		services:   services,
		errorMsg:   "",
		actualPage: "bank",
	}
}

func (m TeaModel) Init() tea.Cmd {
	return nil
}

func (m TeaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	selectedField := &m.fields[m.cursor]

	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch m.mode {

		case TEXTMODE:
			return m.textModeActions(msg, selectedField)
		case NORMALMODE:
			return m.normalModeActions(msg, selectedField)
		}
	}

	return m, nil
}

func (m TeaModel) textModeActions(msg tea.KeyMsg, selectedField *field) (tea.Model, tea.Cmd) {

	switch msg.String() {
	case "esc":
		m.mode = NORMALMODE
		return m, nil
	case "backspace":

		selectedField.value = selectedField.value[:len(selectedField.value)-1]

	// return m, tea.Quit
	default:
		selectedField.value = selectedField.value + msg.String()
	}

	return m, nil
}

func (m TeaModel) normalModeActions(msg tea.KeyMsg, selectedField *field) (tea.Model, tea.Cmd) {
	switch msg.String() {

	case "ctrl+c", "q":
		return m, tea.Quit

	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}

	case "down", "j":
		if m.cursor < len(m.fields)-1 {
			m.cursor++
		}

	case "a":
		m.mode = TEXTMODE

	case "w":
		amount, _ := strconv.ParseFloat(m.fields[0].value, 64)
		_, err := m.services.bankCore.CreateBank(amount)

		if err != nil {
			m.errorMsg = err.Error()
		}

	case "enter", " ":
		m.actualPage = "new_transaction"
	}

	return m, nil

}

func (m TeaModel) View() string {
	s := "New expense\n\n"

	switch m.actualPage {
	case "bank":
		s += BankView(&m)
  case "new_transaction":
    s += "hehe"
	}

	s += fmt.Sprintf("\nPress q to quit. --- %s ----- %s\n", m.mode, m.errorMsg)

	return s
}
