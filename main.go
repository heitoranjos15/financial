package main

import (
	"fmt"
	// "os"

	// "github.com/charmbracelet/bubbles/list"
	// "github.com/charmbracelet/bubbles/textarea"
	// "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	// "github.com/charmbracelet/lipgloss"
)


type Model struct {
  Screen string
  Cursor int
  world bool
}

func (m Model) Init() tea.Cmd{
  return nil
}

func word() string {
  return "world"
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd){
  switch msgType := msg.(type) {
  case tea.KeyMsg:
    switch msgType.String() {
    case "q":
      return m, tea.Quit
    case "n": 
      m.world = true
      return m, nil
    }
  }
  m.world = false
  var cmd tea.Cmd
  return m, cmd
}

func (m Model) View() string{
  if(m.world) {
    return "world"
  }
  return "Hellow"
}

func main() {
  model := Model{}
  p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
	}
}

