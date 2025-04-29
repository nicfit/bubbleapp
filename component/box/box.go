package box

import (
	"image/color"
	"strings"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/charmbracelet/bubbles/v2/key"
	"github.com/charmbracelet/bubbles/v2/viewport"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type BoxOptions struct {
	Bg           *color.Color
	initialChild app.UIModel
}
type model struct {
	base         *app.Base
	opts         BoxOptions
	style        lipgloss.Style
	viewport     viewport.Model
	contentCache string
}

type KeyMap struct {
	Submit key.Binding
}

type BoxOption func(o *BoxOptions)

func WithBg(color color.Color) BoxOption {
	return func(o *BoxOptions) {
		o.Bg = &color
	}
}

func WithChild(item app.UIModel) BoxOption {
	return func(o *BoxOptions) {
		o.initialChild = item
	}
}

func New(ctx *app.Context, opts ...BoxOption) model {
	options := BoxOptions{
		Bg: nil, // TODO theming for default value
	}
	for _, opt := range opts {
		opt(&options)
	}
	base := app.New(ctx, app.WithGrow(true))

	if options.initialChild != nil {
		base.AddChild(options.initialChild)
	}

	viewport := viewport.New()

	style := lipgloss.NewStyle()
	if options.Bg != nil {
		style = style.Background(*options.Bg)
	}

	return model{
		base:         base,
		opts:         options,
		style:        style,
		viewport:     viewport,
		contentCache: "",
	}
}
func (m *model) AddChild(item app.UIModel) {
	m.base.AddChild(item)
}
func (m *model) AddChildren(items ...app.UIModel) {
	for _, item := range items {
		m.base.AddChild(item)
	}
}

func (m model) Init() tea.Cmd {
	return m.base.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m model) View() string {
	if m.base.GetChildren() != nil && len(m.base.GetChildren()) > 0 {
		childContent := make([]string, len(m.base.GetChildren()))
		for i, child := range m.base.GetChildren() {
			childContent[i] = child.View()
		}

		cont := strings.Join(childContent, "\n")
		if m.contentCache != cont {
			m.viewport.SetContent(cont)
		}
		m.contentCache = cont
	}
	return m.base.Ctx.Zone.Mark(m.base.ID, m.style.Height(m.base.Height).Width(m.base.Width).Render(m.viewport.View()))
}

func (m model) Base() *app.Base {
	return m.base
}
