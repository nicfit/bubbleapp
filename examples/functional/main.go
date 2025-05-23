package main

import (
	"os"
	"strconv"
	"time"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/box"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
	"github.com/alexanderbh/bubbleapp/style"

	tea "github.com/charmbracelet/bubbletea/v2"
)

func NewRoot(c *app.Ctx) *app.C {
	clicks, setClicks := app.UseState(c, 0)
	greeting, setGreeting := app.UseState(c, "Knock knock!")

	app.UseEffect(c, func() {
		go func() {
			time.Sleep(2 * time.Second)
			setGreeting("Who's there?")
		}()
	}, []any{})

	return stack.New(c, func(c *app.Ctx) []*app.C {
		return []*app.C{
			button.New(c, "Count clicks here!", func() {
				setClicks(clicks + 1)
			}),

			text.New(c, "Clicks: "+strconv.Itoa(clicks), text.WithFg(c.Theme.Colors.Warning)),
			text.New(c, "Greeting: "+greeting, text.WithFg(c.Theme.Colors.Warning)),

			box.NewEmpty(c),

			button.New(c, "Quit", c.Quit, button.WithVariant(style.Danger)),
		}
	}, stack.WithGap(1), stack.WithGrow(true))
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
