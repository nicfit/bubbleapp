package divider

import (
	"strings"

	"github.com/alexanderbh/bubbleapp/app"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type model struct {
	base  *app.Base
	style lipgloss.Style
}

func New(ctx *app.Context) model {
	style := lipgloss.NewStyle().Foreground(ctx.Styles.Colors.Ghost)

	return model{
		base:  app.New(ctx, app.WithFocusable(false)),
		style: style,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	cmd = m.base.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if m.base.Width == 0 {
		return ""
	}
	return m.style.Render(strings.Repeat("─", m.base.Width-1))
}

func (m model) Base() *app.Base {
	return m.base
}
