package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/box"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/divider"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"

	tea "github.com/charmbracelet/bubbletea/v2"
)

func NewRoot(c *app.Ctx) app.C {
	presses, setPresses := app.UseState(c, 0)
	log, setLog := app.UseState(c, []string{})

	return stack.New(c, func(c *app.Ctx) []app.C {
		return []app.C{
			text.New(c, "Tab through the buttons to see focus state!"),

			button.New(c, "Button 1", func() {
				currentLog := log
				currentPresses := presses
				newLog := append(currentLog, "["+strconv.Itoa(currentPresses)+"] "+"Button 1 pressed")
				setLog(newLog)
				setPresses(currentPresses + 1)
			}, button.WithVariant(button.Primary)),

			divider.New(c),

			box.New(c, func(c *app.Ctx) app.C {
				return text.New(c, strings.Join(log, "\n"))
			}),

			divider.New(c),

			button.New(c, "Quit App", func() {
				c.Quit()
			}, button.WithVariant(button.Danger)),
		}

	}, stack.WithGrow(true))
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
