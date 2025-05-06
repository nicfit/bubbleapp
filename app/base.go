package app

import (
	"github.com/charmbracelet/lipgloss/v2"
)

type Base struct {
	ID              string
	LayoutDirection LayoutDirection
	Shader          Shader
	Opts            BaseOptions
}

type LayoutDirection int

const (
	Vertical LayoutDirection = iota
	Horizontal
)

type BaseOptions struct {
	GrowX           bool
	GrowY           bool
	Focusable       bool
	LayoutDirection LayoutDirection
	Shader          Shader
}

type BaseOption func(*BaseOptions)

func WithGrowX(grow bool) BaseOption {
	return func(o *BaseOptions) {
		o.GrowX = grow
	}
}
func WithGrowY(grow bool) BaseOption {
	return func(o *BaseOptions) {
		o.GrowY = grow
	}
}

func WithGrow(grow bool) BaseOption {
	return func(o *BaseOptions) {
		o.GrowX = grow
		o.GrowY = grow
	}
}
func WithShader(shader Shader) BaseOption {
	return func(o *BaseOptions) {
		o.Shader = shader
	}
}

func WithFocusable(focusable bool) BaseOption {
	return func(o *BaseOptions) {
		o.Focusable = focusable
	}
}
func WithLayoutDirection(direction LayoutDirection) BaseOption {
	return func(o *BaseOptions) {
		o.LayoutDirection = direction
	}
}

// Creates a new base which includes the ID. Returns a cleanup function that must be deferred.
func NewBase[T any](ctx *Context[T], name string, opts ...BaseOption) (*Base, func()) {
	if opts == nil {
		opts = []BaseOption{}
	}
	options := BaseOptions{
		GrowX:           false,
		GrowY:           false,
		Focusable:       false,
		Shader:          nil,
		LayoutDirection: Vertical,
	}

	for _, opt := range opts {
		if opt == nil {
			continue
		}
		opt(&options)
	}

	b := &Base{
		ID:              ctx.id.push(name),
		Opts:            options,
		Shader:          options.Shader,
		LayoutDirection: options.LayoutDirection,
	}

	return b, func() {
		ctx.id.pop()
	}
}

func (base *Base) ApplyShader(input string) string {
	if base.Shader != nil {
		return base.Shader.Render(input, nil)
	}
	return input
}
func (base *Base) ApplyShaderWithStyle(input string, style lipgloss.Style) string {
	if base.Shader != nil {
		return base.Shader.Render(input, &style)
	}
	return style.Render(input)
}
func (base *Base) Tick() {
	if ds, ok := base.Shader.(DynamicShader); ok && ds != nil {
		ds.Tick()
	}
}
