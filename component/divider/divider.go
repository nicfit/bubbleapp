package divider

import (
	"image/color"
	"strings"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/charmbracelet/lipgloss/v2"
)

// TODO: Support GrowY/Vertical divider

type Props struct {
	Char    string      // If not set defaults to "─"
	FGColor color.Color // If nil, defaults to app.Ctx.Styles.Colors.Ghost.
	app.Layout
}

type prop func(*Props)

func Divider(c *app.Ctx, props app.Props) string {
	divProps, ok := props.(Props)
	if !ok {
		return ""
	}

	finalFGColor := divProps.FGColor
	if finalFGColor == nil {
		finalFGColor = c.Styles.Colors.Ghost
	}
	style := lipgloss.NewStyle().Foreground(finalFGColor)

	char := divProps.Char

	var length int
	width, _ := app.UseSize(c)
	if width > 0 {
		length = width - 1
	} else {
		length = 0
	}

	if length <= 0 {
		return ""
	}

	return style.Render(strings.Repeat(char, length))
}

// New creates a new divider element.
func New(c *app.Ctx, opts ...prop) string {
	p := Props{
		Char:    "─",
		FGColor: c.Styles.Colors.Ghost,
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

// Note: WithGrowY might not be typical for a horizontal divider,
// but could be added if vertical dividers or other use cases emerge.
// func WithGrowY(grow bool) prop {
// 	return func(props *Props) {
// 		props.Layout.GrowY = grow
// 	}
// }
