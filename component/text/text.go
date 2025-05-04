package text

import (
	"image/color"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/style"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type Options struct {
	Foreground color.Color
	Background color.Color
	Bold       bool
	style.Margin
}

type text[T any] struct {
	base   *app.Base[T]
	render func(ctx *app.Context[T]) string
	opts   *Options
	style  lipgloss.Style
}

func New[T any](ctx *app.Context[T], text string, options *Options, baseOptions ...app.BaseOption) *text[T] {
	return NewDynamic(ctx, func(ctx *app.Context[T]) string {
		return text
	}, options, baseOptions...)
}

func NewDynamic[T any](ctx *app.Context[T], render func(ctx *app.Context[T]) string, options *Options, baseOptions ...app.BaseOption) *text[T] {
	if options == nil {
		options = &Options{}
	}
	if baseOptions == nil {
		baseOptions = []app.BaseOption{}
	}
	base := app.NewBase[T](append([]app.BaseOption{}, baseOptions...)...)

	if options.Foreground == nil {
		options.Foreground = lipgloss.NoColor{}
	}
	if options.Background == nil {
		options.Background = lipgloss.NoColor{}
	}

	s := lipgloss.NewStyle().Foreground(options.Foreground).Background(options.Background)

	s = style.ApplyMargin(s, options.Margin)

	if options.Bold {
		s = s.Bold(true)
	}
	return &text[T]{
		base:   base,
		render: render,
		style:  s,
		opts:   options,
	}
}

func (m *text[T]) Render(ctx *app.Context[T]) string {
	return m.base.ApplyShaderWithStyle(m.render(ctx), m.style)
}

func (m *text[T]) Update(ctx *app.Context[T], msg tea.Msg) {

}

func (m *text[T]) Children(ctx *app.Context[T]) []app.Fc[T] {
	return nil
}

func (m *text[T]) Base() *app.Base[T] {
	return m.base
}
