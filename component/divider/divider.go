package divider

import (
	"strings"

	"github.com/alexanderbh/bubbleapp/app"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type divider[T any] struct {
	base  *app.Base
	style lipgloss.Style
}

func New[T any](ctx *app.Context[T], baseOptions ...app.BaseOption) *divider[T] {
	if baseOptions == nil {
		baseOptions = []app.BaseOption{}
	}
	style := lipgloss.NewStyle().Foreground(ctx.Styles.Colors.Ghost)

	return &divider[T]{
		// TODO: Support GrowY/Vertical divider
		base:  app.NewBase[T](append([]app.BaseOption{app.WithGrowX(true)}, baseOptions...)...),
		style: style,
	}
}

func (m *divider[T]) Render(ctx *app.Context[T]) string {
	if m.base.Width == 0 {
		return ""
	}
	return m.style.Render(strings.Repeat("─", m.base.Width-1))
}

func (m *divider[T]) Update(ctx *app.Context[T], msg tea.Msg) {

}

func (m *divider[T]) Children(ctx *app.Context[T]) []app.Fc[T] {
	return nil
}

func (m *divider[T]) Base() *app.Base {
	return m.base
}
