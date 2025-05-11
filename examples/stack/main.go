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

	stack := stack.New(c, func(c *app.Ctx) {
		box.NewEmpty(c, box.WithBg(c.Styles.Colors.Danger))
		box.New(c, func(c *app.Ctx) {
			stack.New(c, func(c *app.Ctx) {
				box.NewEmpty(c, box.WithBg(c.Styles.Colors.Primary))
				box.NewEmpty(c, box.WithBg(c.Styles.Colors.Secondary))
				box.NewEmpty(c, box.WithBg(c.Styles.Colors.Tertiary))

			}, stack.WithDirection(app.Horizontal))
		})
		box.NewEmpty(c, box.WithBg(c.Styles.Colors.Warning))
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
