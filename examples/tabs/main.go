package main

import (
	"os"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/tabs"

	tea "github.com/charmbracelet/bubbletea/v2"
)

var tabsData = []tabs.Tab{
	{Title: "Overview", Content: overview},
	{Title: "Loaders", Content: loaders},
	{Title: "Boxes 🟨", Content: boxes},
}

func NewRoot(c *app.Ctx, _ app.Props) app.C {
	return tabs.New(c, tabsData)
}

func main() {
	c := app.NewCtx()

	bubbleApp := app.New(c, NewRoot)
	p := tea.NewProgram(bubbleApp, tea.WithAltScreen(), tea.WithMouseAllMotion())
	bubbleApp.SetTeaProgram(p)
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
