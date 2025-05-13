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
	Key           string
	Bg            color.Color
	DisableFollow bool
	Children      app.Children
	app.Layout
}

// BoxProp is a function type for setting BoxProps.
type BoxProp func(*BoxProps)

// Box is the functional component for rendering a box with a viewport.
func Box(c *app.Ctx, props app.Props) string {
	boxProps, ok := props.(BoxProps)
	if !ok {
		panic("Box component requires BoxProps")
	}

	// Create a new viewport model instance first
	initialViewport := viewport.New()
	// Then pass its address to UseState to store a pointer
	vp, _ := app.UseState(c, &initialViewport)

	// Get children
	childrenContent := app.UseChildren(c, boxProps.Children)
	renderedChildren := strings.Join(childrenContent, "\n")

	width, height := app.UseSize(c)

	if width <= 0 || height <= 0 {
		return ""
	}

	// Is this right? When trying to get intrinsic size it feels like this should not be set
	vp.SetWidth(width)
	vp.SetHeight(height)

	vp.SetContent(renderedChildren)
	if !boxProps.DisableFollow {
		vp.GotoBottom()
	}

	style := lipgloss.NewStyle()
	if boxProps.Bg != nil {
		style = style.Background(boxProps.Bg)
	}

	finalRender := style.Width(width).Height(height).Render(vp.View())

	return finalRender
}

// New creates a new Box component.
func New(c *app.Ctx, children app.Children, opts ...BoxProp) string {
	appliedProps := BoxProps{
		Children:      children,
		DisableFollow: false,
		Layout: app.Layout{
			GrowX: true,
			GrowY: true,
		},
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

func WithKey(key string) BoxProp {
	return func(props *BoxProps) {
		props.Key = key
	}
}

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
		props.Layout.GrowX = grow
		props.Layout.GrowY = grow
	}
}

func WithGrowX(grow bool) BoxProp {
	return func(props *BoxProps) {
		props.Layout.GrowX = grow
	}
}
func WithGrowY(grow bool) BoxProp {
	return func(props *BoxProps) {
		props.Layout.GrowY = grow
	}
}
