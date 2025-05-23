package box

import (
	"image/color"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/box/viewport"
	"github.com/alexanderbh/bubbleapp/component/router"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

// BoxProps defines the properties for the Box component.
type BoxProps struct {
	Key           string
	Bg            color.Color
	DisableFollow bool
	Child         app.FC
	app.Layout
	app.Border
}

// BoxProp is a function type for setting BoxProps.
type BoxProp func(*BoxProps)

// Box is the functional component for rendering a box with a viewport.
func Box(c *app.Ctx, props app.Props) string {
	boxProps, ok := props.(BoxProps)
	if !ok {
		panic("Box component requires BoxProps")
	}

	initialViewport := viewport.New()
	initialViewport.MouseWheelDelta = 1
	vp, _ := app.UseState(c, &initialViewport)

	// Scroll to top on route change
	// Not sure if this should be default behavior. It makes sense for now.
	router := router.UseRouterController(c)
	currentRoute := router.Current()
	app.UseEffect(c, func() {
		vp.GotoTop()
	}, []any{currentRoute})
	// End scroll to top on route change

	width, height := app.UseSize(c)
	if boxProps.Layout.Width > 0 {
		width = boxProps.Layout.Width
	}
	if boxProps.Layout.Height > 0 {
		height = boxProps.Layout.Height
	}
	if width <= 0 || height <= 0 {
		return ""
	}

	app.UseMouseHandler(c, func(msg tea.MouseMsg, childID string) bool {
		switch msg := msg.(type) {
		case tea.MouseWheelMsg:
			switch msg.Button {
			case tea.MouseWheelDown:
				if vp.AtBottom() {
					return false
				}
			case tea.MouseWheelUp:
				if vp.AtTop() {
					return false
				}
			}
			newVp, _ := vp.Update(msg)
			vp.SetYOffset(newVp.YOffset())
			vp.SetXOffset(newVp.XOffset())

			return true
		}
		return false
	})

	// Is this right? When trying to get intrinsic size it feels like this should not be set
	vp.SetWidth(width)
	vp.SetHeight(height)

	style := app.ApplyBorder(lipgloss.NewStyle(), boxProps.Border)
	if boxProps.Bg != nil {
		style = style.Background(boxProps.Bg)
		beforeCurrentBg := c.CurrentBg
		c.CurrentBg = boxProps.Bg
		vp.Style = vp.Style.Background(c.CurrentBg)
		defer func() {
			c.CurrentBg = beforeCurrentBg
		}()
	} else if c.CurrentBg != nil {
		style = style.Background(c.CurrentBg)
		vp.Style = vp.Style.Background(c.CurrentBg)
	}

	if boxProps.Child != nil {
		fc := boxProps.Child(c).String()
		vp.SetContent(fc)
		if !boxProps.DisableFollow {
			vp.GotoBottom()
		}
	}

	finalRender := style.Width(width).Height(height).Render(vp.View())

	return c.MouseZone(finalRender)
}

// New creates a new Box component.
func New(c *app.Ctx, child app.FC, opts ...BoxProp) *app.C {
	appliedProps := BoxProps{
		Child:         child,
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
func NewEmpty(c *app.Ctx, opts ...BoxProp) *app.C {
	return New(c, nil, append([]BoxProp{WithDisableFollow(true)}, opts...)...)
}
