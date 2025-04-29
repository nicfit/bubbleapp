package text

import (
	"image/color"

	"github.com/alexanderbh/bubbleapp/app"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type TextOptions struct {
	Foreground color.Color
	Background color.Color
}
type TextOption func(o *TextOptions)

func WithColor(color color.Color) TextOption {
	return func(o *TextOptions) {
		o.Foreground = color
	}
}
func WithBackgroundColor(color color.Color) TextOption {
	return func(o *TextOptions) {
		o.Background = color
	}
}

type model struct {
	base  *app.Base
	text  string
	opts  TextOptions
	style lipgloss.Style
}

func New(ctx *app.Context, text string, opts ...TextOption) model {
	options := TextOptions{
		Foreground: lipgloss.NoColor{},
		Background: lipgloss.NoColor{},
	}
	for _, opt := range opts {
		opt(&options)
	}

	style := lipgloss.NewStyle().Foreground(options.Foreground).Background(options.Background)

	return model{
		base:  app.New(ctx),
		text:  text,
		style: style,
		opts:  options,
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
	return m.style.UnsetWidth().Render(m.text)
}

func (m model) Base() *app.Base {
	return m.base
}
