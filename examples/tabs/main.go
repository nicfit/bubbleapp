package main

import (
	"os"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/tabs"
	"github.com/alexanderbh/bubbleapp/style"

	zone "github.com/alexanderbh/bubblezone/v2"
	tea "github.com/charmbracelet/bubbletea/v2"
)

var tabsData = []tabs.TabElement{
	{
		Title: "Overview",
		Content: func(ctx *app.Context) app.UIModel {
			return NewOverview(ctx)
		},
	},
	{
		Title: "Loaders",
		Content: func(ctx *app.Context) app.UIModel {
			return NewLoaders(ctx)
		},
	},
	{
		Title: "Scolling",
		Content: func(ctx *app.Context) app.UIModel {
			return NewScrolling(ctx)
		},
	},
}

func NewRoot() model {
	ctx := &app.Context{
		Styles:       style.DefaultStyles(),
		FocusManager: app.NewFocusManager(),
		Zone:         zone.New(),
	}

	tabs := tabs.New(ctx, tabsData)

	base := app.New(ctx, app.AsRoot())
	base.AddChild(tabs)

	return model{
		base: base,
	}
}

type model struct {
	base *app.Base
}

func (m model) Init() tea.Cmd {
	return m.base.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmd := m.base.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.base.View()
}

func main() {
	p := tea.NewProgram(NewRoot(), tea.WithAltScreen(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
