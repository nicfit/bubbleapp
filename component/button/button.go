package button

import (
	"strings"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/style"
	"github.com/charmbracelet/lipgloss/v2"
)

type Props struct {
	Variant  style.Variant
	Text     string
	Disabled bool
	OnAction func()
	Bold     bool
	app.Layout
	app.Margin
	app.Padding
}

type prop func(*Props)

func Button(c *app.Ctx, props app.Props) string {
	buttonProps, _ := props.(Props)

	id := app.UseID(c)
	focused := app.UseIsFocused(c)
	hovered, _ := app.UseIsHovered(c)

	app.UseAction(c, func(_ string) {
		if buttonProps.OnAction != nil && !buttonProps.Disabled {
			c.FocusThis(id)
			buttonProps.OnAction()
		}
	})

	state := style.Normal
	if buttonProps.Disabled {
		state = style.Disabled
	} else if hovered {
		state = style.Hover
	} else if focused {
		state = style.Focus
	}

	style := c.Theme.Button[buttonProps.Variant][state]

	if buttonProps.Layout.Height > 0 {
		style = style.Height(buttonProps.Layout.Height)
	}

	var internalLeftPadding string = ""
	var internalRightPadding string = ""
	if buttonProps.Layout.Width > 0 {
		style = style.Width(buttonProps.Layout.Width)

		repeasts := buttonProps.Layout.Width - lipgloss.Width(buttonProps.Text) - 2
		if repeasts > 0 {
			if repeasts%2 == 0 {
				internalLeftPadding = strings.Repeat(" ", repeasts/2)
				internalRightPadding = strings.Repeat(" ", repeasts/2)
			} else {
				internalLeftPadding = strings.Repeat(" ", repeasts/2+1)
				internalRightPadding = strings.Repeat(" ", repeasts/2)
			}
		}
	}

	if focused {
		buttonProps.Text = "⟨" + internalLeftPadding + buttonProps.Text + internalRightPadding + "⟩"
	} else {
		buttonProps.Text = "[" + internalLeftPadding + buttonProps.Text + internalRightPadding + "]"
	}

	if buttonProps.Bold {
		style = style.Bold(true)
	}
	style = app.ApplyMargin(app.ApplyPadding(style, buttonProps.Padding), buttonProps.Margin)

	return c.MouseZone(style.Render(buttonProps.Text))
}

func New(c *app.Ctx, text string, onAction func(), props ...prop) app.C {
	p := Props{
		Text:     text,
		OnAction: onAction,
	}
	for _, prop := range props {
		prop(&p)
	}
	return c.Render(Button, p)
}

func WithVariant(variant style.Variant) prop {
	return func(props *Props) {
		props.Variant = variant
	}
}
func WithWidth(width int) prop {
	return func(props *Props) {
		props.Width = width
	}
}
func WithHeight(height int) prop {
	return func(props *Props) {
		props.Height = height
	}
}
