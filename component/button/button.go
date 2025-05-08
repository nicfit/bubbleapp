package button

import (
	"github.com/alexanderbh/bubbleapp/app"
)

type ButtonProps struct {
	Text     string
	OnAction func()
}

func Button(c *app.FCContext, props app.Props) string {
	buttonProps, _ := props.(ButtonProps)
	id := c.UseID()

	focused := c.UseFocus()

	if !focused {
		return c.Zone.Mark(id, "button {"+buttonProps.Text+"}")
	}
	return c.Zone.Mark(id, "BUTTON {"+buttonProps.Text+"}")
}

func NewButton(c *app.FCContext, text string, onAction func()) string {
	return c.Render(Button, ButtonProps{
		Text:     text,
		OnAction: onAction,
	})
}
