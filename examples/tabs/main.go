package main

import (
	"os"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/tabs"

	tea "github.com/charmbracelet/bubbletea/v2"
)

var tabsData = []tabs.TabElement[CustomData]{
	{Title: "Overview", Content: NewOverview},
	{Title: "Loaders", Content: NewLoaders},
	{Title: "Scolling", Content: NewScrolling},
}

type CustomData struct {
	HowCoolIsThis string
}

func NewRoot(ctx *app.Context[CustomData]) app.Fc[CustomData] {
	return tabs.New(ctx, tabsData)
}

func main() {
	ctx := app.NewContext(&CustomData{
		HowCoolIsThis: "Very cool!",
	})

	app := app.NewApp(ctx, NewRoot)
	p := tea.NewProgram(app, tea.WithAltScreen(), tea.WithMouseAllMotion())
	app.SetTeaProgram(p)
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
