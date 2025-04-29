package main

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
	"github.com/alexanderbh/bubbleapp/component/tickfps"
	"github.com/alexanderbh/bubbleapp/style"

	zone "github.com/alexanderbh/bubblezone/v2"
	tea "github.com/charmbracelet/bubbletea/v2"
)

func NewAuthModel(userID string) authModel {
	ctx := &app.Context{
		Styles:       style.DefaultStyles(),
		FocusManager: app.NewFocusManager(),
		Zone:         zone.New(),
	}

	stack := stack.New(ctx)
	stack.AddChildren(
		text.New(ctx, "You are logged in as: "+userID), // Find a way to generically have custom data in app.Context to save userID and more
		text.New(ctx, "Press [q] to quit.\n"),
		tickfps.New(ctx),
	)

	base := app.New(ctx, app.AsRoot())
	base.AddChild(stack)

	return authModel{
		base: base,
	}
}

type authModel struct {
	base *app.Base
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
	return m.base.View()
}
