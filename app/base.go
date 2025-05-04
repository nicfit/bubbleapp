package app

import (
	"github.com/alexanderbh/bubbleapp/shader"
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/google/uuid"
)

type Base[T any] struct {
	ID              string
	Parent          Fc[T]
	LayoutDirection LayoutDirection
	Shader          shader.Shader
	Width           int
	Height          int
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
	Shader          shader.Shader
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
func WithShader(shader shader.Shader) BaseOption {
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

func NewBase[T any](opts ...BaseOption) *Base[T] {
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
		opt(&options)
	}

	b := &Base[T]{
		ID:              uuid.New().String(),
		Opts:            options,
		Shader:          options.Shader,
		LayoutDirection: options.LayoutDirection,
	}

	return b
}

func (base *Base[T]) ApplyShader(input string) string {
	if base.Shader != nil {
		return base.Shader.Render(input, nil)
	}
	return input
}
func (base *Base[T]) ApplyShaderWithStyle(input string, style lipgloss.Style) string {
	if base.Shader != nil {
		return base.Shader.Render(input, &style)
	}
	return style.Render(input)
}
