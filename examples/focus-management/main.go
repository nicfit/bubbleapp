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

type CustomData struct{}

func NewRoot() model {
	ctx := &app.Context[CustomData]{
		Styles: style.DefaultStyles(),
		Zone:   zone.New(),
	}

	boxFill := box.New(ctx, &box.Options[CustomData]{})
	addButton := button.New(ctx, "Button 1", &button.Options{Variant: button.Primary})
	quitButton := button.New(ctx, "Quit App", &button.Options{Variant: button.Danger})

	stack := stack.New(ctx, &stack.Options[CustomData]{
		Children: []*app.Base[CustomData]{
			text.New(ctx, "Tab through the buttons to see focus state!", nil),
			addButton,
			boxFill,
			divider.New(ctx),
			quitButton,
		}}, app.AsRoot(),
	)

	return model{
		base:         stack,
		containerID:  boxFill.ID,
		addButtonID:  addButton.ID,
		quitButtonID: quitButton.ID,
	}
}

type model struct {
	base *app.Base[CustomData]

	containerID  string
	addButtonID  string
	quitButtonID string
}

func (m model) Init() tea.Cmd {
	return m.base.Init()
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
			m.base.GetChild(m.containerID).AddChild(
				text.New(m.base.Ctx, "Button pressed", nil),
			)
			return m, nil
		}
	}

	cmd = m.base.Update(msg)

	return m, cmd

}

func (m model) View() string {
	return m.base.Render()
}

func main() {
	p := tea.NewProgram(NewRoot(), tea.WithAltScreen(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
