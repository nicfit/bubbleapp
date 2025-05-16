package main

import (
	"os"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/box"
	"github.com/alexanderbh/bubbleapp/component/stack"

	tea "github.com/charmbracelet/bubbletea/v2"
)

type CustomData struct{}

func NewRoot(c *app.Ctx) app.C {
	return stack.New(c, func(c *app.Ctx) []app.C {
		return []app.C{
			box.NewEmpty(c, box.WithBg(c.Styles.Colors.Danger)),
			box.New(c, func(c *app.Ctx) app.C {
				return stack.New(c, func(c *app.Ctx) []app.C {
					return []app.C{
						box.NewEmpty(c, box.WithBg(c.Styles.Colors.Primary)),
						box.NewEmpty(c, box.WithBg(c.Styles.Colors.Secondary)),
						box.NewEmpty(c, box.WithBg(c.Styles.Colors.Tertiary)),
					}
				}, stack.WithDirection(app.Horizontal))
			}),
			box.NewEmpty(c, box.WithBg(c.Styles.Colors.Warning)),
		}
	})
}

func main() {
	c := app.NewCtx()

	bubbleApp := app.New(c, NewRoot)
	p := tea.NewProgram(bubbleApp, tea.WithAltScreen())
	bubbleApp.SetTeaProgram(p)

	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
