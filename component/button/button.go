package button

import (
	"strings"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/charmbracelet/lipgloss/v2"
)

type Props struct {
	Variant  Variant
	Type     Type
	Text     string
	Disabled bool
	OnAction func()
	Bold     bool
	app.Layout
	app.Margin
	app.Padding
}

type prop func(*Props)

type Type int

const (
	Normal Type = iota
	Bordered
)

type Variant int

const (
	Primary Variant = iota
	Secondary
	Tertiary
	Success
	Danger
	Info
	Warning
)

func Button(c *app.Ctx, props app.Props) string {
	buttonProps, _ := props.(Props)

	focused := app.UseIsFocused(c)
	hovered, _ := app.UseIsHovered(c)

	app.UseAction(c, func(_ string) {
		if buttonProps.OnAction != nil && !buttonProps.Disabled {
			buttonProps.OnAction()
		}
	})

	style := styleResolver(c, buttonProps.Variant, buttonProps.Type, focused, hovered, buttonProps.Disabled)

	if buttonProps.Layout.Height > 0 {
		style = style.Height(buttonProps.Layout.Height)
	}
	if buttonProps.Layout.Width > 0 {
		style = style.Width(buttonProps.Layout.Width)
		if buttonProps.Type == Normal {
			repeasts := buttonProps.Layout.Width - lipgloss.Width(buttonProps.Text) - 4
			if repeasts > 0 {
				rs := strings.Repeat(" ", repeasts)
				buttonProps.Text = "⟦ " + buttonProps.Text + rs + " ⟧"
			}
		}
	} else if buttonProps.Type == Normal {
		if focused {
			buttonProps.Text = "[>" + buttonProps.Text + "<]"
		} else {
			buttonProps.Text = "[ " + buttonProps.Text + " ]"
		}
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

func WithVariant(variant Variant) prop {
	return func(props *Props) {
		props.Variant = variant
	}
}
func WithType(btnType Type) prop {
	return func(props *Props) {
		props.Type = btnType
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

// OMG I hate this. How to make a nice design system that maybe can also work for custom components?
// This is a mess. I need to think about this more.
func styleResolver(c *app.Ctx, variant Variant, btnType Type, focused bool, hovered bool, disabled bool) lipgloss.Style {
	if btnType == Normal {
		if hovered {
			switch variant {
			case Primary:
				return c.Styles.Button.PrimaryHovered
			case Secondary:
				return c.Styles.Button.SecondaryHovered
			case Tertiary:
				return c.Styles.Button.TertiaryHovered
			case Success:
				return c.Styles.Button.SuccessHovered
			case Danger:
				return c.Styles.Button.DangerHovered
			case Info:
				return c.Styles.Button.InfoHovered
			case Warning:
				return c.Styles.Button.WarningHovered
			}
		} else if focused {
			switch variant {
			case Primary:
				return c.Styles.Button.PrimaryFocused
			case Secondary:
				return c.Styles.Button.SecondaryFocused
			case Tertiary:
				return c.Styles.Button.TertiaryFocused
			case Success:
				return c.Styles.Button.SuccessFocused
			case Danger:
				return c.Styles.Button.DangerFocused
			case Info:
				return c.Styles.Button.InfoFocused
			case Warning:
				return c.Styles.Button.WarningFocused
			}
		} else {
			switch variant {
			case Primary:
				return c.Styles.Button.Primary
			case Secondary:
				return c.Styles.Button.Secondary
			case Tertiary:
				return c.Styles.Button.Tertiary
			case Success:
				return c.Styles.Button.Success
			case Danger:
				return c.Styles.Button.Danger
			case Info:
				return c.Styles.Button.Info
			case Warning:
				return c.Styles.Button.Warning
			}
		}
	} else if btnType == Bordered {
		if hovered {
			switch variant {
			case Primary:
				return c.Styles.ButtonBordered.PrimaryHovered
			case Secondary:
				return c.Styles.ButtonBordered.SecondaryHovered
			case Tertiary:
				return c.Styles.ButtonBordered.TertiaryHovered
			case Success:
				return c.Styles.ButtonBordered.SuccessHovered
			case Danger:
				return c.Styles.ButtonBordered.DangerHovered
			case Info:
				return c.Styles.ButtonBordered.InfoHovered
			case Warning:
				return c.Styles.ButtonBordered.WarningHovered
			}
		} else if focused {
			switch variant {
			case Primary:
				return c.Styles.ButtonBordered.PrimaryFocused
			case Secondary:
				return c.Styles.ButtonBordered.SecondaryFocused
			case Tertiary:
				return c.Styles.ButtonBordered.TertiaryFocused
			case Success:
				return c.Styles.ButtonBordered.SuccessFocused
			case Danger:
				return c.Styles.ButtonBordered.DangerFocused
			case Info:
				return c.Styles.ButtonBordered.InfoFocused
			case Warning:
				return c.Styles.ButtonBordered.WarningFocused
			}
		} else {
			switch variant {
			case Primary:
				return c.Styles.ButtonBordered.Primary
			case Secondary:
				return c.Styles.ButtonBordered.Secondary
			case Tertiary:
				return c.Styles.ButtonBordered.Tertiary
			case Success:
				return c.Styles.ButtonBordered.Success
			case Danger:
				return c.Styles.ButtonBordered.Danger
			case Info:
				return c.Styles.ButtonBordered.Info
			case Warning:
				return c.Styles.ButtonBordered.Warning
			}
		}
	}
	return c.Styles.ButtonBordered.Primary
}
