package main

import (
	"os"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/box"
	"github.com/alexanderbh/bubbleapp/component/stack"

	tea "github.com/charmbracelet/bubbletea/v2"
)

type CustomData struct{}

func NewRoot(c *app.Ctx) *app.C {
	return stack.New(c, func(c *app.Ctx) []*app.C {
		return []*app.C{
			box.NewEmpty(c, box.WithBg(c.Theme.Colors.DangerLight)),
			stack.New(c, func(c *app.Ctx) []*app.C {
				return []*app.C{
					box.NewEmpty(c, box.WithBg(c.Theme.Colors.PrimaryLight)),
					box.NewEmpty(c, box.WithBg(c.Theme.Colors.SecondaryLight)),
					box.NewEmpty(c, box.WithBg(c.Theme.Colors.TertiaryLight)),
				}
			}, stack.WithDirection(app.Horizontal), stack.WithGrow(true)),

			box.NewEmpty(c, box.WithBg(c.Theme.Colors.WarningLight)),
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
