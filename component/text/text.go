package text

import (
	"image/color"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/style"
	"github.com/charmbracelet/lipgloss/v2"
)

// Props defines the properties for the Text component.
type Props struct {
	Content    func(*app.Ctx) string
	Foreground color.Color
	Background color.Color
	Bold       bool
	style.Margin
	// TODO: Consider adding Width, Height, MaxWidth, MaxHeight if explicit control is needed directly in props
}

// prop is a function type for setting Props.
type prop func(*Props)

// TextFC is the core functional component for rendering text.
func Text(c *app.Ctx, props app.Props) string {
	textProps, ok := props.(Props)
	if !ok {
		// In a real scenario, you might log an error or return a specific error string.
		return ""
	}

	var renderedContent string

	if textProps.Content != nil {
		renderedContent = textProps.Content(c)
	} else {
		renderedContent = "" // Handle nil function case
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

	s = style.ApplyMargin(s, textProps.Margin)

	return s.Render(renderedContent)
}

// NewText creates a new text element.
// Content can be a static string or a dynamic function: func(c *app.FCContext) string.
func NewText(c *app.Ctx, content string, opts ...prop) string {
	p := Props{
		Content: func(c *app.Ctx) string {
			return content
		},
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

// --- Prop Option Functions ---

// WithFg sets the foreground color.
func WithFg(fg color.Color) prop {
	return func(props *Props) {
		props.Foreground = fg
	}
}

// WithBg sets the background color.
func WithBg(bg color.Color) prop {
	return func(props *Props) {
		props.Background = bg
	}
}

// WithBold enables or disables bold text.
func WithBold(bold bool) prop {
	return func(props *Props) {
		props.Bold = bold
	}
}

// WithMarginAll sets uniform margin for all sides.
func WithMarginAll(m int) prop {
	return func(props *Props) {
		props.Margin.M = m
	}
}

// WithMargin sets individual margins.
func WithMargin(top, right, bottom, left int) prop {
	return func(props *Props) {
		props.Margin.MT = top
		props.Margin.MR = right
		props.Margin.MB = bottom
		props.Margin.ML = left
	}
}

// WithMarginTop sets the top margin.
func WithMarginTop(m int) prop {
	return func(props *Props) {
		props.Margin.MT = m
	}
}

// WithMarginRight sets the right margin.
func WithMarginRight(m int) prop {
	return func(props *Props) {
		props.Margin.MR = m
	}
}

// WithMarginBottom sets the bottom margin.
func WithMarginBottom(m int) prop {
	return func(props *Props) {
		props.Margin.MB = m
	}
}

// WithMarginLeft sets the left margin.
func WithMarginLeft(m int) prop {
	return func(props *Props) {
		props.Margin.ML = m
	}
}

// WithMarginX sets horizontal (left and right) margins.
func WithMarginX(m int) prop {
	return func(props *Props) {
		props.Margin.MX = m
	}
}

// WithMarginY sets vertical (top and bottom) margins.
func WithMarginY(m int) prop {
	return func(props *Props) {
		props.Margin.MY = m
	}
}
