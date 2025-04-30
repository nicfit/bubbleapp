package main

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"

	tea "github.com/charmbracelet/bubbletea/v2"
)

func NewOverview(ctx *app.Context[CustomData]) app.UIModel[CustomData] {
	quitButton := button.New(ctx, "Quit", button.WithVariant(button.Danger))

	stack := stack.New(ctx, stack.Options[CustomData]{
		Children: []*app.Base[CustomData]{
			text.New(ctx, "\nFor now you navigate tabs with arrow keys.\nThey should have shortcuts probably. And perhaps navigate with tab? Or vim keys?\n\n").Base(),
			text.New(ctx, "From global data: "+ctx.Data.HowCoolIsThis).Base(),
			quitButton.Base(),
		}},
	)

	base := app.New(ctx)
	base.AddChild(stack.Base())

	return overviewModel[CustomData]{
		base:         base,
		quitButtonID: quitButton.Base().ID,
	}
}

type overviewModel[T CustomData] struct {
	base         *app.Base[T]
	quitButtonID string
}

func (m overviewModel[T]) Init() tea.Cmd {
	return m.base.Init()
}

func (m overviewModel[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m overviewModel[T]) View() string {
	return m.base.Render()
}

func (m overviewModel[T]) Base() *app.Base[T] {
	m.base.Model = m
	return m.base
}
