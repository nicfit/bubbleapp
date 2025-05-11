package main

import (
	"os"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/box"
	"github.com/alexanderbh/bubbleapp/component/stack"

	tea "github.com/charmbracelet/bubbletea/v2"
)

type CustomData struct{}

func NewRoot(c *app.Ctx, _ app.Props) string {

	stack := stack.New(c, func(ctx *app.Ctx) {

		box.NewEmpty(ctx, box.WithBg(ctx.Styles.Colors.Danger))
		box.New(ctx, func(ctx *app.Ctx) {
			stack.New(ctx, func(ctx *app.Ctx) {

				box.NewEmpty(ctx, box.WithBg(ctx.Styles.Colors.Primary))
				box.NewEmpty(ctx, box.WithBg(ctx.Styles.Colors.Secondary))
				box.NewEmpty(ctx, box.WithBg(ctx.Styles.Colors.Tertiary))

			}, stack.WithDirection(app.Horizontal))
		})
		box.NewEmpty(ctx, box.WithBg(ctx.Styles.Colors.Warning))

	})

	return stack
}

func main() {
	ctx := app.NewCtx()

	app := app.New(ctx, NewRoot)
	p := tea.NewProgram(app, tea.WithAltScreen())
	app.SetTeaProgram(p)

	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
