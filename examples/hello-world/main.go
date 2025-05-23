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

func NewRoot(c *app.Ctx) *app.C {

	return stack.New(c, func(c *app.Ctx) []*app.C {
		return []*app.C{
			text.New(c, "Hello World!", nil),
			divider.New(c),
			text.New(c, "Press [ctrl-c] to quit.", nil),
		}
	})
}

func main() {
	c := app.NewCtx()

	bubbleApp := app.New(c, NewRoot)
	p := tea.NewProgram(bubbleApp, tea.WithAltScreen(), tea.WithMouseAllMotion())
	bubbleApp.SetTeaProgram(p)
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
