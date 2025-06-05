package main

import (
	"os"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/divider"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
	"github.com/alexanderbh/bubbleapp/component/textfield"
	"github.com/alexanderbh/bubbleapp/style"

	tea "github.com/charmbracelet/bubbletea/v2"
)

type CustomData struct{}

func NewRoot(c *app.Ctx) *app.C {

	textValue, setTextValue := app.UseState(c, "")

	return stack.New(c, func(c *app.Ctx) []*app.C {
		return []*app.C{
			textfield.New(c, func(text string) {
				setTextValue(text)
			}, textValue, textfield.WithTitle("Type something here:")),
			divider.New(c),
			text.New(c, "You typed: "+textValue),
			text.New(c, "Press [ctrl-c] to quit."),
			button.New(c, "Quit", c.Quit, button.WithVariant(style.Danger)),
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
