package main

import (
	"os"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/box"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/grid"
	"github.com/alexanderbh/bubbleapp/component/loader"
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

	gridView := grid.New(ctx)
	gridView.AddItems(
		grid.NewItem(box.New(ctx, box.WithBg(ctx.Styles.Colors.PrimaryDark), box.WithChild(
			text.New(ctx, "I wish I could center text! Some day...")),
		), grid.WithXs(12)),
		grid.NewItem(box.New(ctx, box.WithBg(ctx.Styles.Colors.Warning)), grid.WithXs(6)),
		grid.NewItem(button.New(ctx, "BUTTON 1", button.WithVariant(button.Success)), grid.WithXs(6)),
		grid.NewItem(button.New(ctx, "BUTTON 2"), grid.WithXs(3)),
		grid.NewItem(box.New(ctx, box.WithBg(ctx.Styles.Colors.InfoDark), box.WithChild(
			stack.New(ctx, stack.WithChildren(
				text.New(ctx, "I am in a stack!"),
				loader.New(ctx, loader.Meter, loader.WithText("Text style messes up bg. Fix!"), loader.WithColor(ctx.Styles.Colors.Black))),
			),
		)), grid.WithXs(6)),
		grid.NewItem(box.New(ctx, box.WithBg(ctx.Styles.Colors.Success)), grid.WithXs(3)),
	)

	base := app.New(ctx, app.AsRoot())
	base.AddChild(gridView)

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
