package main

import (
	"os"

	"github.com/alexanderbh/bubbleapp/app"
	tea "github.com/charmbracelet/bubbletea/v2"
)

func main() {
	ctx := app.NewContext(&CustomData{})

	app := app.NewApp(ctx, NewLoginRoot)
	p := tea.NewProgram(app, tea.WithAltScreen(), tea.WithMouseAllMotion())
	app.SetTeaProgram(p)
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
