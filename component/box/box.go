package box

import (
	"image/color"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/charmbracelet/bubbles/v2/viewport"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type Options struct {
	Bg            color.Color
	DisableFollow bool
}
type box[T any] struct {
	base         *app.Base
	opts         *Options
	style        lipgloss.Style
	viewport     viewport.Model
	contentCache string
	child        app.Fc[T]
}

func NewEmpty[T any](ctx *app.Context[T], options *Options, baseOptions ...app.BaseOption) *box[T] {
	return New(ctx, func(ctx *app.Context[T]) app.Fc[T] { return nil }, options, baseOptions...)
}

func New[T any](ctx *app.Context[T], child func(ctx *app.Context[T]) app.Fc[T], options *Options, baseOptions ...app.BaseOption) *box[T] {
	if options == nil {
		options = &Options{}
	}
	if baseOptions == nil {
		baseOptions = []app.BaseOption{}
	}

	base, cleanup := app.NewBase(ctx, "box", append([]app.BaseOption{app.WithGrow(true)}, baseOptions...)...)
	defer cleanup()

	viewport := viewport.New()

	style := lipgloss.NewStyle()
	if options.Bg != nil {
		style = style.Background(options.Bg)
	}

	c := child(ctx)

	return &box[T]{
		base:         base,
		opts:         options,
		style:        style,
		viewport:     viewport,
		contentCache: "",
		child:        c,
	}
}

func (m box[T]) Render(ctx *app.Context[T]) string {

	m.viewport.SetWidth(ctx.UIState.GetWidth(m.base.ID))
	m.viewport.SetHeight(ctx.UIState.GetHeight(m.base.ID))

	childFc := m.child
	if childFc != nil {
		child := childFc.Render(ctx)
		if m.contentCache != child {
			m.viewport.SetContent(child)
			if !m.opts.DisableFollow {
				m.viewport.GotoBottom()
			}
		}
		m.contentCache = child
	}

	// This does not work when not Grow. Needs to not limit on LayoutPhase. Fix.
	return m.style.Height(ctx.UIState.GetHeight(m.base.ID)).Width(ctx.UIState.GetWidth(m.base.ID)).Render(m.viewport.View())
}

func (m *box[T]) Children(ctx *app.Context[T]) []app.Fc[T] {
	child := m.child
	if child != nil {
		return []app.Fc[T]{child}
	}
	return []app.Fc[T]{}

}

func (m *box[T]) Update(ctx *app.Context[T], msg tea.Msg) bool {
	return false
}

func (m *box[T]) Base() *app.Base {
	return m.base
}
