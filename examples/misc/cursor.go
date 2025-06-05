package main

import (
	"os"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type Model struct {
	cursor *tea.Cursor
	text   string
}

func NewModel() Model {
	return Model{
		text: "",
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			if m.cursor == nil {
				m.cursor = tea.NewCursor(2+lipgloss.Width(m.text), 1)
			} else {
				m.cursor = nil
			}
		default:
			m.text += msg.String()
		}
	}
	return m, nil
}

func (m Model) View() (string, *tea.Cursor) {
	desc := "Press enter to toggle cursor. It will not show up. Press a letter (maybe two) to change the output and then the cursor shows up."
	return desc + "\n> " + m.text, m.cursor
}

func main() {

	p := tea.NewProgram(NewModel(), tea.WithAltScreen(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
