package text

import (
	"image/color"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/style"
	"github.com/charmbracelet/lipgloss/v2"
)

type Props struct {
	Variant    style.Variant
	Text       string
	Foreground color.Color
	Background color.Color
	Bold       bool
	app.Margin
	app.Padding
}

type prop func(*Props)

// Text is the core functional component for rendering text.
func Text(c *app.Ctx, rawProps app.Props) string {
	props, ok := rawProps.(Props)
	if !ok {
		return ""
	}

	s := c.Theme.Text[props.Variant][style.Normal]

	if props.Foreground != nil {
		s = s.Foreground(props.Foreground)
	}
	if props.Background != nil {
		s = s.Background(props.Background)
	} else {
		s = s.Background(lipgloss.NoColor{})
	}
	if props.Bold {
		s = s.Bold(true)
	}

	s = app.ApplyMargin(app.ApplyPadding(s, props.Padding), props.Margin)

	return s.Render(props.Text)
}

// New creates a new text element.
func New(c *app.Ctx, text string, opts ...prop) app.C {
	p := Props{
		Text:    text,
		Variant: style.Base,
	}

	for _, opt := range opts {
		if opt != nil {
			opt(&p)
		}
	}
	return c.Render(Text, p)
}
