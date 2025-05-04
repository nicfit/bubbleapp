package main

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
	"github.com/alexanderbh/bubbleapp/component/tickfps"

	tea "github.com/charmbracelet/bubbletea/v2"
)

func NewAuthModel(ctx *app.Context[CustomData]) authModel {
	stack := stack.New(ctx, &stack.Options[CustomData]{
		Children: []*app.Base[CustomData]{
			text.New(ctx, "You are logged in as: "+ctx.Data.UserID, nil), // Find a way to generically have custom data in app.Context to save userID and more
			text.New(ctx, "Press [q] to quit.\n", nil),
			tickfps.New(ctx),
		}}, app.AsRoot(),
	)

	return authModel{
		base: stack,
	}
}

type authModel struct {
	base *app.Base[CustomData]
}

func (m authModel) Init() tea.Cmd {
	return m.base.Init()
}

func (m authModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	cmd := m.base.Update(msg)

	return m, cmd

}

func (m authModel) View() string {
	return m.base.Render()
}
