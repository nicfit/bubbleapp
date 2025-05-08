package main

import (
	"os"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/stack"

	tea "github.com/charmbracelet/bubbletea/v2"
)

func NewRoot(c *app.FCContext, _ app.Props) string {

	return stack.New(c, func(c *app.FCContext) {
		button.NewButton(c, "CLICK ME 1", func() {
			c.Quit()
		})
		button.NewButton(c, "CLICK ME 2", nil)
		button.NewButton(c, "CLICK ME 3", nil)
		button.NewButton(c, "CLICK ME 4", nil)
	}, stack.WithDirection(app.Horizontal), stack.WithGap(2))

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
