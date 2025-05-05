package markdown

import (
	"github.com/alexanderbh/bubbleapp/app"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/glamour"
)

// TODO:
//   - Add support for styles and custom styles

type markdown[T any] struct {
	base            *app.Base
	render          func(ctx *app.Context[T]) string
	glamourRenderer *glamour.TermRenderer
}

func New[T any](ctx *app.Context[T], text string, baseOptions ...app.BaseOption) *markdown[T] {
	return NewDynamic(ctx, func(ctx *app.Context[T]) string {
		return text
	}, baseOptions...)
}

func NewDynamic[T any](ctx *app.Context[T], render func(ctx *app.Context[T]) string, baseOptions ...app.BaseOption) *markdown[T] {
	if baseOptions == nil {
		baseOptions = []app.BaseOption{}
	}
	r, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(ctx.Width-1),
	)

	return &markdown[T]{
		base:            app.NewBase[T]("markdown", append([]app.BaseOption{app.WithGrow(true)}, baseOptions...)...),
		render:          render,
		glamourRenderer: r,
	}
}

func (m *markdown[T]) Render(ctx *app.Context[T]) string {
	out, _ := m.glamourRenderer.Render(m.render(ctx))
	return out
}

func (m *markdown[T]) Update(ctx *app.Context[T], msg tea.Msg) {

}

func (m *markdown[T]) Children(ctx *app.Context[T]) []app.Fc[T] {
	return nil
}

func (m *markdown[T]) Base() *app.Base {
	return m.base
}
