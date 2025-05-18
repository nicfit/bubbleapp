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
	app.Layout
}

type Prop func(*Props)

type Type int

const (
	Normal Type = iota
	Bordered
	Flat
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
		// Can this be part of theming somehow?
		buttonProps.Text = "⟦ " + buttonProps.Text + " ⟧"
	}

	return c.MouseZone(style.Render(buttonProps.Text))
}

func New(c *app.Ctx, text string, onAction func(), props ...Prop) app.C {
	p := Props{
		Text:     text,
		OnAction: onAction,
	}
	for _, prop := range props {
		prop(&p)
	}
	return c.Render(Button, p)
}

func WithVariant(variant Variant) Prop {
	return func(props *Props) {
		props.Variant = variant
	}
}
func WithType(btnType Type) Prop {
	return func(props *Props) {
		props.Type = btnType
	}
}
func WithWidth(width int) Prop {
	return func(props *Props) {
		props.Width = width
	}
}
func WithHeight(height int) Prop {
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
				return c.Styles.ButtonCompact.PrimaryHovered
			case Secondary:
				return c.Styles.ButtonCompact.SecondaryHovered
			case Tertiary:
				return c.Styles.ButtonCompact.TertiaryHovered
			case Success:
				return c.Styles.ButtonCompact.SuccessHovered
			case Danger:
				return c.Styles.ButtonCompact.DangerHovered
			case Info:
				return c.Styles.ButtonCompact.InfoHovered
			case Warning:
				return c.Styles.ButtonCompact.WarningHovered
			}
		} else if focused {
			switch variant {
			case Primary:
				return c.Styles.ButtonCompact.PrimaryFocused
			case Secondary:
				return c.Styles.ButtonCompact.SecondaryFocused
			case Tertiary:
				return c.Styles.ButtonCompact.TertiaryFocused
			case Success:
				return c.Styles.ButtonCompact.SuccessFocused
			case Danger:
				return c.Styles.ButtonCompact.DangerFocused
			case Info:
				return c.Styles.ButtonCompact.InfoFocused
			case Warning:
				return c.Styles.ButtonCompact.WarningFocused
			}
		} else {
			switch variant {
			case Primary:
				return c.Styles.ButtonCompact.Primary
			case Secondary:
				return c.Styles.ButtonCompact.Secondary
			case Tertiary:
				return c.Styles.ButtonCompact.Tertiary
			case Success:
				return c.Styles.ButtonCompact.Success
			case Danger:
				return c.Styles.ButtonCompact.Danger
			case Info:
				return c.Styles.ButtonCompact.Info
			case Warning:
				return c.Styles.ButtonCompact.Warning
			}
		}
	} else if btnType == Flat {
		if hovered {
			switch variant {
			case Primary:
				return c.Styles.ButtonFlat.PrimaryHovered
			case Secondary:
				return c.Styles.ButtonFlat.SecondaryHovered
			case Tertiary:
				return c.Styles.ButtonFlat.TertiaryHovered
			case Success:
				return c.Styles.ButtonFlat.SuccessHovered
			case Danger:
				return c.Styles.ButtonFlat.DangerHovered
			case Info:
				return c.Styles.ButtonFlat.InfoHovered
			case Warning:
				return c.Styles.ButtonFlat.WarningHovered
			}
		} else if focused {
			switch variant {
			case Primary:
				return c.Styles.ButtonFlat.PrimaryFocused
			case Secondary:
				return c.Styles.ButtonFlat.SecondaryFocused
			case Tertiary:
				return c.Styles.ButtonFlat.TertiaryFocused
			case Success:
				return c.Styles.ButtonFlat.SuccessFocused
			case Danger:
				return c.Styles.ButtonFlat.DangerFocused
			case Info:
				return c.Styles.ButtonFlat.InfoFocused
			case Warning:
				return c.Styles.ButtonFlat.WarningFocused
			}
		} else {
			switch variant {
			case Primary:
				return c.Styles.ButtonFlat.Primary
			case Secondary:
				return c.Styles.ButtonFlat.Secondary
			case Tertiary:
				return c.Styles.ButtonFlat.Tertiary
			case Success:
				return c.Styles.ButtonFlat.Success
			case Danger:
				return c.Styles.ButtonFlat.Danger
			case Info:
				return c.Styles.ButtonFlat.Info
			case Warning:
				return c.Styles.ButtonFlat.Warning
			}
		}
	} else if btnType == Bordered {
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
	}
	return c.Styles.Button.Primary
}
