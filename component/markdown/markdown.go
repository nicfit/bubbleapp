package markdown

import (
	"github.com/alexanderbh/bubbleapp/app"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/glamour"
)

type model[T any] struct {
	base            *app.Base[T]
	text            string
	glamourRenderer *glamour.TermRenderer
}

func New[T any](ctx *app.Context[T], text string, baseOptions ...app.BaseOption) *app.Base[T] {
	if baseOptions == nil {
		baseOptions = []app.BaseOption{}
	}
	r, _ := glamour.NewTermRenderer(
		glamour.WithWordWrap(ctx.Width - 1),
	)

	return model[T]{
		base:            app.NewBase(ctx, baseOptions...),
		text:            text,
		glamourRenderer: r,
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

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.glamourRenderer, _ = glamour.NewTermRenderer(
			glamour.WithWordWrap(msg.Width - 1),
		)
	}

	cmd = m.base.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model[T]) View() string {
	out, _ := m.glamourRenderer.Render(m.text)
	return out
}

func (m model[T]) Base() *app.Base[T] {
	m.base.Model = m
	return m.base
}
