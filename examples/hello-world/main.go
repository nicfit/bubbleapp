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

func NewRoot(c *app.Ctx, _ app.Props) string {

	stack := stack.New(c, func(c *app.Ctx) {
		text.New(c, "Hello World!", nil)
		divider.New(c)
		text.New(c, "Press [ctrl-c] to quit.", nil)
	})

	return stack
}

func main() {
	ctx := app.NewCtx()

	bubbleApp := app.New(ctx, NewRoot)
	p := tea.NewProgram(bubbleApp, tea.WithAltScreen(), tea.WithMouseAllMotion())
	bubbleApp.SetTeaProgram(p)
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
