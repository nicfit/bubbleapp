package text

import (
	"github.com/alexanderbh/bubbleapp/app"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type model struct {
	base *app.Base
	text string
}

func New(ctx *app.Context, text string) model {
	return model{
		base: app.New(ctx),
		text: text,
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
	return m.text
}

func (m model) Base() *app.Base {
	return m.base
}
