package text

import (
	"image/color"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/shader"
	"github.com/alexanderbh/bubbleapp/style"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type Options struct {
	Foreground color.Color
	Background color.Color
	Bold       bool
	Shader     shader.Shader
	style.Margin
}

type model[T any] struct {
	base  *app.Base[T]
	text  string
	opts  *Options
	style lipgloss.Style
}

func New[T any](ctx *app.Context[T], text string, options *Options) *app.Base[T] {
	if options == nil {
		options = &Options{}
	}

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
	return model[T]{
		base:  app.New(ctx, app.WithShader(options.Shader)),
		text:  text,
		style: s,
		opts:  options,
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
	return m.base.ApplyShaderWithStyle(m.text, m.style)
}

func (m model[T]) Base() *app.Base[T] {
	m.base.Model = m
	return m.base
}
