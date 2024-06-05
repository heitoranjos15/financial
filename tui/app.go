package tui

import (
	"context"
	"financial/adapters/postgres"
	"financial/tui/pages"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	TEXTMODE   = "text"
	NORMALMODE = "normal"
)

type TeaModel struct {
	pageIndex int
	mode      string
	pages     []*pages.DefaultPage
}

func newBankPage(db *pgxpool.Pool) *pages.DefaultPage {
	homepage := pages.NewBankPage(db)
	page := pages.DefaultPage{
		Name:    "home",
		Actions: homepage,
	}
	return &page
}

func InitialModel() TeaModel {
	ctx := context.Background()
	conn := postgres.GetConnection(ctx)

	postgres.RunMigrations()

	bankPage := newBankPage(conn)

	return TeaModel{
		pages:     []*pages.DefaultPage{bankPage},
		mode:      NORMALMODE,
		pageIndex: 0,
	}
}

func (m TeaModel) Init() tea.Cmd {
	return nil
}

func (m TeaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.mode {
		case NORMALMODE:
			return m.normalModeActions(msg)
		}
	}

	return m, nil
}

func (m TeaModel) normalModeActions(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	currentPage := m.pages[m.pageIndex]
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
  case "enter":
		m.Next()
  default:
    currentPage.Actions.NormalModeActions(msg.String())
	}
	return m, nil
}

func (m *TeaModel) Next() {
	if m.pageIndex < len(m.pages)-1 {
		m.pageIndex++
	} else {
		m.pageIndex = 0
	}
}

func (m TeaModel) View() string {
	s := "Expenses\n\n"
	currentPage := m.pages[m.pageIndex]
	s += currentPage.Actions.View()

	return s
}
