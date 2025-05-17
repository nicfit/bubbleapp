package text

import (
	"image/color"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/charmbracelet/lipgloss/v2"
)

type Props struct {
	Text       string
	Foreground color.Color
	Background color.Color
	Bold       bool
	app.Margin
	app.Padding
}

type prop func(*Props)

// Text is the core functional component for rendering text.
func Text(c *app.Ctx, props app.Props) string {
	textProps, ok := props.(Props)
	if !ok {
		return ""
	}

	s := lipgloss.NewStyle()

	if textProps.Foreground != nil {
		s = s.Foreground(textProps.Foreground)
	} else {
		s = s.Foreground(lipgloss.NoColor{}) // Default as in original
	}

	if textProps.Background != nil {
		s = s.Background(textProps.Background)
	} else {
		s = s.Background(lipgloss.NoColor{}) // Default as in original
	}

	if textProps.Bold {
		s = s.Bold(true)
	}

	s = app.ApplyMargin(app.ApplyPadding(s, textProps.Padding), textProps.Margin)

	return s.Render(textProps.Text)
}

// New creates a new text element.
func New(c *app.Ctx, text string, opts ...prop) app.C {
	p := Props{
		Text:       text,
		Foreground: lipgloss.NoColor{},
		Background: lipgloss.NoColor{},
	}

	for _, opt := range opts {
		if opt != nil {
			opt(&p)
		}
	}
	return c.Render(Text, p)
}
