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

	tea "github.com/charmbracelet/bubbletea/v2"
)

func NewRoot(c *app.Ctx, _ app.Props) string {
	clicks, setClicks := app.UseState(c, 0)
	greeting, setGreeting := app.UseState(c, "Knock knock!")

	app.UseEffect(c, func() {
		go func() {
			time.Sleep(2 * time.Second)
			setGreeting("Who's there?")
		}()
	}, []any{})

	return stack.New(c, func(c *app.Ctx) {
		button.New(c, "Count clicks here!", func() {
			setClicks(clicks + 1)
		})

		text.New(c, "Clicks: "+strconv.Itoa(clicks), text.WithFg(c.Styles.Colors.Warning))
		text.New(c, "Greeting: "+greeting, text.WithFg(c.Styles.Colors.Warning))

		box.NewEmpty(c)

		button.New(c, "Quit", func() {
			c.Quit()
		}, button.WithVariant(button.Danger))
	}, stack.WithGap(1), stack.WithGrow(true))
}

func main() {
	ctx := app.NewCtx()

	app := app.New(ctx, NewRoot)
	p := tea.NewProgram(app, tea.WithAltScreen(), tea.WithMouseAllMotion())
	app.SetTeaProgram(p)
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
