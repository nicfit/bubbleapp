package main

import (
	"os"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/divider"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"

	tea "github.com/charmbracelet/bubbletea/v2"
)

type CustomData struct{}

func NewRoot(ctx *app.Context[CustomData]) app.Fc[CustomData] {

	stack := stack.New(ctx, func(ctx *app.Context[CustomData]) []app.Fc[CustomData] {
		return []app.Fc[CustomData]{
			text.New(ctx, "Hello World!", nil),
			divider.New(ctx),
			text.New(ctx, "Press [ctrl-c] to quit.", nil),
		}
	}, nil)

	return stack
}

func main() {
	ctx := app.NewContext(&CustomData{})

	p := tea.NewProgram(app.NewApp(ctx, NewRoot), tea.WithAltScreen(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
