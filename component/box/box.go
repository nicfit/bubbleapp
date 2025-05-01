package box

import (
	"image/color"
	"strings"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/charmbracelet/bubbles/v2/viewport"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type Options[T any] struct {
	Bg    color.Color
	Child *app.Base[T]
}
type model[T any] struct {
	base         *app.Base[T]
	opts         *Options[T]
	style        lipgloss.Style
	viewport     viewport.Model
	contentCache string
}

func New[T any](ctx *app.Context[T], options *Options[T]) *app.Base[T] {
	if options == nil {
		options = &Options[T]{}
	}
	base := app.New(ctx, app.WithGrow(true))

	if options.Child != nil {
		base.AddChild(options.Child)
	}

	viewport := viewport.New()

	style := lipgloss.NewStyle()
	if options.Bg != nil {
		style = style.Background(options.Bg)
	}

	return model[T]{
		base:         base,
		opts:         options,
		style:        style,
		viewport:     viewport,
		contentCache: "",
	}.Base()
}

func (m model[T]) Init() tea.Cmd {
	return m.base.Init()
}

func (m model[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.viewport.SetWidth(msg.Width)
		m.viewport.SetHeight(msg.Height)
		// Does not support multiple children with fill

	case tea.KeyMsg:
		if m.base.Focused {
			m.viewport, cmd = m.viewport.Update(msg)
			cmds = append(cmds, cmd)
		}
	case tea.MouseMsg:
		if m.base.Ctx.Zone.Get(m.base.ID).InBounds(msg) {
			m.viewport, cmd = m.viewport.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	cmd = m.base.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model[T]) View() string {
	if len(m.base.Children) > 0 {
		childContent := make([]string, len(m.base.Children))
		for i, child := range m.base.Children {
			childContent[i] = child.Model.View()
		}

		cont := strings.Join(childContent, "\n")
		if m.contentCache != cont {
			m.viewport.SetContent(cont)
		}
		m.contentCache = cont
	}
	return m.base.Ctx.Zone.Mark(m.base.ID, m.style.Height(m.base.Height).Width(m.base.Width).Render(m.viewport.View()))
}

func (m model[T]) Base() *app.Base[T] {
	m.base.Model = m
	return m.base
}
