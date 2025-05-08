package main

import (
	"os"
	"strconv"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"

	tea "github.com/charmbracelet/bubbletea/v2"
)

func NewRoot(c *app.FCContext, _ app.Props) string {

	clicks, setClicks := app.UseState(c, 0)
	greeting, setGreeting := app.UseState(c, "Knock knock!")

	return stack.New(c, func(c *app.FCContext) {
		button.NewButton(c, "Count clicks here!", func() {
			setClicks(clicks + 1)
		}, button.WithType(button.Compact))

		text.NewText(c, "Clicks: "+strconv.Itoa(clicks), text.WithForeground(c.Styles.Colors.Warning))
		text.NewText(c, "Greeting: "+greeting, text.WithForeground(c.Styles.Colors.Warning))

		button.NewButton(c, "Quit", func() {
			c.Quit()
		}, button.WithVariant(button.Danger), button.WithType(button.Compact))
	}, stack.WithGap(1))

}

func main() {
	ctx := app.NewFCContext()

	app := app.New(ctx, NewRoot)
	p := tea.NewProgram(app, tea.WithAltScreen(), tea.WithMouseAllMotion())
	app.SetTeaProgram(p)
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
