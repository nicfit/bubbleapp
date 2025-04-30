package divider

import (
	"strings"

	"github.com/alexanderbh/bubbleapp/app"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type model[T any] struct {
	base  *app.Base[T]
	style lipgloss.Style
}

func New[T any](ctx *app.Context[T]) *app.Base[T] {
	style := lipgloss.NewStyle().Foreground(ctx.Styles.Colors.Ghost)

	return model[T]{
		base:  app.New(ctx, app.WithFocusable(false)),
		style: style,
	}.Base()
}

func (m model[T]) Init() tea.Cmd {
	return nil
}

func (m model[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	cmd = m.base.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model[T]) View() string {
	if m.base.Width == 0 {
		return ""
	}
	return m.style.Render(strings.Repeat("─", m.base.Width-1))
}

func (m model[T]) Base() *app.Base[T] {
	m.base.Model = m
	return m.base
}
