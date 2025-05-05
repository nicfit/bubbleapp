package main

import (
	"os"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/box"
	"github.com/alexanderbh/bubbleapp/component/stack"

	tea "github.com/charmbracelet/bubbletea/v2"
)

type CustomData struct{}

func NewRoot(ctx *app.Context[CustomData]) app.Fc[CustomData] {

	stack := stack.New(ctx, []app.Fc[CustomData]{
		box.NewEmpty(ctx, &box.Options{Bg: ctx.Styles.Colors.Danger}),
		box.New(ctx, stack.New(ctx, []app.Fc[CustomData]{
			box.NewEmpty(ctx, &box.Options{Bg: ctx.Styles.Colors.Primary}),
			box.NewEmpty(ctx, &box.Options{Bg: ctx.Styles.Colors.Secondary}),
			box.NewEmpty(ctx, &box.Options{Bg: ctx.Styles.Colors.Tertiary}),
		}, &stack.Options{Horizontal: true}), nil),
		box.NewEmpty(ctx, &box.Options{Bg: ctx.Styles.Colors.Warning}),
	}, nil)

	return stack
}

func main() {
	ctx := app.NewContext(&CustomData{})

	p := tea.NewProgram(app.NewApp(ctx, NewRoot), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
