package divider

import (
	"image/color"
	"strings"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/charmbracelet/lipgloss/v2"
)

// TODO: Support GrowY/Vertical divider

type Props struct {
	Char    string
	FGColor color.Color
	app.Layout
}

type prop func(*Props)

func Divider(c *app.Ctx, rawProps app.Props) string {
	props, ok := rawProps.(Props)
	if !ok {
		return ""
	}

	style := lipgloss.NewStyle().Foreground(props.FGColor)

	char := props.Char

	var length int
	width, _ := app.UseSize(c)
	if width > 0 {
		length = width
	} else {
		length = 0
	}

	if length <= 0 {
		return ""
	}

	return style.Render(strings.Repeat(char, length))
}

// New creates a new divider element.
func New(c *app.Ctx, opts ...prop) app.C {
	p := Props{
		Char:    "─",
		FGColor: c.Theme.Colors.Base600,
		Layout: app.Layout{
			GrowX: true,
			GrowY: false,
		},
	}

	for _, opt := range opts {
		if opt != nil {
			opt(&p)
		}
	}
	return c.Render(Divider, p)
}

// --- Prop Option Functions ---

// WithChar sets the character to be repeated for the divider.
func WithChar(char string) prop {
	return func(props *Props) {
		props.Char = char
	}
}

// WithFGColor sets the foreground color of the divider.
func WithFGColor(fg color.Color) prop {
	return func(props *Props) {
		props.FGColor = fg
	}
}

// WithGrowX sets whether the divider should grow horizontally.
func WithGrowX(grow bool) prop {
	return func(props *Props) {
		props.Layout.GrowX = grow
	}
}
