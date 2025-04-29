package main

import (
	"os"

	tea "github.com/charmbracelet/bubbletea/v2"
)

func main() {
	p := tea.NewProgram(NewLogin(), tea.WithAltScreen(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
