package main

import (
	"os"

	"github.com/alexanderbh/bubbleapp/app"
	tea "github.com/charmbracelet/bubbletea/v2"
)

func main() {
	ctx := app.NewContext(&CustomData{})

	p := tea.NewProgram(app.NewApp(ctx, NewLoginRoot), tea.WithAltScreen(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
