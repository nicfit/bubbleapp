package divider

import (
	"strings"

	"github.com/alexanderbh/bubbleapp/app"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

// TODO: Support GrowY/Vertical divider
type divider[T any] struct {
	base  *app.Base
	style lipgloss.Style
}

func New[T any](ctx *app.Context[T], baseOptions ...app.BaseOption) *divider[T] {
	if baseOptions == nil {
		baseOptions = []app.BaseOption{}
	}
	style := lipgloss.NewStyle().Foreground(ctx.Styles.Colors.Ghost)
	base, cleanup := app.NewBase(ctx, "divider", append([]app.BaseOption{app.WithGrowX(true)}, baseOptions...)...)
	defer cleanup()

	return &divider[T]{
		base:  base,
		style: style,
	}
}

func (m *divider[T]) Render(ctx *app.Context[T]) string {
	if ctx.UIState.GetWidth(m.base.ID) == 0 {
		return ""
	}
	// Why -1 here? Check if this is required and so why
	return m.style.Render(strings.Repeat("─", ctx.UIState.GetWidth(m.base.ID)-1))
}

func (m *divider[T]) Update(ctx *app.Context[T], msg tea.Msg) bool {
	return false
}

func (m *divider[T]) Children(ctx *app.Context[T]) []app.Fc[T] {
	return nil
}

func (m *divider[T]) Base() *app.Base {
	return m.base
}
