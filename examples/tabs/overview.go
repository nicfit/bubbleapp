package main

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"

	tea "github.com/charmbracelet/bubbletea/v2"
)

func NewOverview(ctx *app.Context) app.UIModel {
	quitButton := button.New(ctx, "Quit", button.WithVariant(button.Danger))

	stack := stack.New(ctx)
	stack.AddChildren(
		text.New(ctx, "\nFor now you navigate tabs with arrow keys.\nThey should have shortcuts probably. And perhaps navigate with tab? Or vim keys?\n\n"),
		quitButton,
	)

	base := app.New(ctx)
	base.AddChild(stack)

	return overviewModel{
		base: base,

		quitButtonID: quitButton.Base().ID,
	}
}

type overviewModel struct {
	base         *app.Base
	quitButtonID string
}

func (m overviewModel) Init() tea.Cmd {
	return m.base.Init()
}

func (m overviewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case button.ButtonPressMsg:
		if msg.ID == m.quitButtonID {
			return m, tea.Quit
		}
	}
	cmd := m.base.Update(msg)

	return m, cmd

}

func (m overviewModel) View() string {
	return m.base.View()
}

func (m overviewModel) Base() *app.Base {
	return m.base
}
