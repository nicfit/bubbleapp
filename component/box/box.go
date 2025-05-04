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
	base         *app.Base[T]
	opts         *Options
	style        lipgloss.Style
	viewport     viewport.Model
	contentCache string
	child        func(ctx *app.Context[T]) app.Fc[T]
}

func New[T any](ctx *app.Context[T], child app.Fc[T], options *Options, baseOptions ...app.BaseOption) *box[T] {
	return NewDynamic(ctx, func(ctx *app.Context[T]) app.Fc[T] { return child }, options, baseOptions...)
}

func NewDynamic[T any](ctx *app.Context[T], child func(ctx *app.Context[T]) app.Fc[T], options *Options, baseOptions ...app.BaseOption) *box[T] {
	if options == nil {
		options = &Options{}
	}
	if baseOptions == nil {
		baseOptions = []app.BaseOption{}
	}
	base := app.NewBase[T](append([]app.BaseOption{app.WithGrow(true)}, baseOptions...)...)

	viewport := viewport.New()

	style := lipgloss.NewStyle()
	if options.Bg != nil {
		style = style.Background(options.Bg)
	}

	return &box[T]{
		base:         base,
		opts:         options,
		style:        style,
		viewport:     viewport,
		contentCache: "",
		child:        child,
	}
}

func (m box[T]) Render(ctx *app.Context[T]) string {
	child := m.child(ctx).Render(ctx)

	m.viewport.SetWidth(m.base.Width)
	m.viewport.SetHeight(m.base.Height)

	if m.contentCache != child {
		m.viewport.SetContent(child)
		if !m.opts.DisableFollow {
			m.viewport.GotoBottom()
		}
	}
	m.contentCache = child

	return m.style.Height(m.base.Height).Width(m.base.Width).Render(m.viewport.View())
}

func (m *box[T]) Children(ctx *app.Context[T]) []app.Fc[T] {
	return []app.Fc[T]{m.child(ctx)}

}

func (m *box[T]) Update(ctx *app.Context[T], msg tea.Msg) {

}

func (m *box[T]) Base() *app.Base[T] {
	return m.base
}
