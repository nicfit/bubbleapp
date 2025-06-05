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
	app.Layout
}

type prop func(*Props)

// Text is the core functional component for rendering text.
func Text(c *app.Ctx, rawProps app.Props) string {
	props, ok := rawProps.(Props)
	if !ok {
		return ""
	}

	width, height := app.UseSize(c)

	s := c.Theme.Text[props.Variant][style.Normal]

	if props.Foreground != nil {
		s = s.Foreground(props.Foreground)
	} else if s.GetForeground() == nil {
		s = s.Foreground(lipgloss.NoColor{})
	}
	if props.Background != nil {
		s = s.Background(props.Background)
	} else if c.CurrentBg != nil {
		s = s.Background(c.CurrentBg)
	}
	if props.Bold {
		s = s.Bold(true)
	}

	if props.Layout.Height > 0 {
		s = s.Height(props.Layout.Height)
	}
	if props.Layout.Width > 0 {
		s = s.Width(props.Layout.Width)
	}

	s = app.ApplyMargin(app.ApplyPadding(s, props.Padding), props.Margin)

	return s.MaxWidth(width).MaxHeight(height).Render(props.Text)
}

// New creates a new text element.
func New(c *app.Ctx, text string, opts ...prop) *app.C {
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
