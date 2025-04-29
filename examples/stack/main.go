package main

import (
	"os"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/box"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/style"

	zone "github.com/alexanderbh/bubblezone/v2"
	tea "github.com/charmbracelet/bubbletea/v2"
)

func NewRoot() model {
	ctx := &app.Context{
		Styles:       style.DefaultStyles(),
		FocusManager: app.NewFocusManager(),
		Zone:         zone.New(),
	}

	stack := stack.New(ctx)
	stack.AddChildren(
		box.New(ctx, box.WithBg(ctx.Styles.Colors.Danger)),
		box.New(ctx, box.WithBg(ctx.Styles.Colors.Warning)),
		box.New(ctx, box.WithBg(ctx.Styles.Colors.Success)),
	)

	base := app.New(ctx, app.AsRoot())
	base.AddChild(stack)

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

func (m model) View() string {
	return m.base.View()
}

func main() {
	p := tea.NewProgram(NewRoot(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
