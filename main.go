package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	TEXTMODE   = "text"
	NORMALMODE = "normal"
)

type model struct {
	fields   []field
	cursor   int
	selected map[int]struct{}
	mode     string
}

type field struct {
	name  string
	value string
}

func initialModel() model {
	return model{
    fields: []field{{name: "Amount"}, {name: "Date"}, {name: "Category"}, {name: "Description"}},

		selected: make(map[int]struct{}),
		mode:     NORMALMODE,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m model) textModeActions(msg tea.KeyMsg, selectedField *field) (tea.Model, tea.Cmd) {

	switch msg.String() {
  case "esc" :
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

func (m model) normalModeActions(msg tea.KeyMsg, selectedField *field) (tea.Model, tea.Cmd) {
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

	case "enter", " ":
		_, ok := m.selected[m.cursor]
		if ok {
			delete(m.selected, m.cursor)
		} else {
			m.selected[m.cursor] = struct{}{}
		}
	}

	return m, nil

}

func (m model) View() string {
	s := "New expense\n\n"

	for i, field := range m.fields {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s [%s]: %s\n", cursor, field.name, field.value)
	}

	s += fmt.Sprintf("\nPress q to quit. --- %s\n", m.mode)

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
