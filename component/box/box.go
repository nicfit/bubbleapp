package box

import (
	"image/color"
	"strings"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/charmbracelet/bubbles/v2/viewport"
	"github.com/charmbracelet/lipgloss/v2"
)

// BoxProps defines the properties for the Box component.
type BoxProps struct {
	Bg            color.Color
	DisableFollow bool
	Children      app.Children
	Size          app.Size
}

// BoxProp is a function type for setting BoxProps.
type BoxProp func(*BoxProps)

// Box is the functional component for rendering a box with a viewport.
func Box(c *app.Ctx, props app.Props) string {
	boxProps, ok := props.(BoxProps)
	if !ok {
		return ""
	}

	id := app.UseID(c)

	vp, _ := app.UseState(c, viewport.New())

	// Get children
	childrenContent := app.UseChildren(c, boxProps.Children)
	renderedChildren := strings.Join(childrenContent, "\n")

	// For prevContent, T is string, so prevContent is string, setPrevContent is func(string).
	prevContent, setPrevContent := app.UseState(c, "")

	// Get dimensions from the UI state (populated by the layout system)
	width := c.UIState.GetWidth(id)
	height := c.UIState.GetHeight(id)

	vp.SetWidth(width)
	vp.SetHeight(height)

	// Update viewport content if it has changed
	if prevContent != renderedChildren {
		vp.SetContent(renderedChildren)
		setPrevContent(renderedChildren) // Update string state for comparison
		if !boxProps.DisableFollow {
			vp.GotoBottom() // or vp.GotoTop() depending on desired behavior
		}
	}

	style := lipgloss.NewStyle()
	if boxProps.Bg != nil {
		style = style.Background(boxProps.Bg)
	}

	finalRender := style.Width(width).Height(height).Render(vp.View())

	return c.Zone.Mark(id, finalRender)
}

// New creates a new Box component.
func New(c *app.Ctx, children app.Children, opts ...BoxProp) string {
	appliedProps := BoxProps{
		Children: children,
		// Default values
		DisableFollow: false,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(&appliedProps)
		}
	}
	return c.Render(Box, appliedProps)
}

// NewEmpty creates a new Box component with no children.
func NewEmpty(c *app.Ctx, opts ...BoxProp) string {
	return New(c, nil, opts...)
}

// --- Prop Option Functions ---

// WithBg sets the background color for the box.
func WithBg(bg color.Color) BoxProp {
	return func(props *BoxProps) {
		props.Bg = bg
	}
}

// WithDisableFollow disables the viewport's auto-scrolling to the bottom on content change.
func WithDisableFollow(disable bool) BoxProp {
	return func(props *BoxProps) {
		props.DisableFollow = disable
	}
}

func WithGrow(grow bool) BoxProp {
	return func(props *BoxProps) {
		props.Size.GrowX = grow
		props.Size.GrowY = grow
	}
}
