package main

import (
	"os"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/box"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/divider"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
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

	boxFill := box.New(ctx)

	addButton := button.New(ctx, "Button 1",
		button.WithVariant(button.Primary),
	)

	quitButton := button.New(ctx, "Quit App",
		button.WithVariant(button.Danger),
	)

	stack := stack.New(ctx)
	stack.AddChildren(
		text.New(ctx, "Tab through the buttons to see focus state!"),
		addButton,
		boxFill,
		divider.New(ctx),
		quitButton,
	)

	base := app.New(ctx, app.AsRoot())
	base.AddChild(stack)

	return model{
		base:         base,
		containerID:  boxFill.Base().ID,
		addButtonID:  addButton.Base().ID,
		quitButtonID: quitButton.Base().ID,
	}
}

type model struct {
	base *app.Base

	containerID  string
	addButtonID  string
	quitButtonID string
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.base.Init(), m.base.Ctx.FocusManager.FocusFirstCmd(m.base.GetChildren()[0]))
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case button.ButtonPressMsg:
		switch msg.ID {
		case m.quitButtonID:
			return m, tea.Quit
		case m.addButtonID:
			m.base.GetChild(m.containerID).Base().AddChild(
				text.New(m.base.Ctx, "Button pressed"),
			)
			return m, nil
		}
	}

	cmd = m.base.Update(msg)

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
