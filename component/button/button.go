package button

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/charmbracelet/bubbles/v2/key"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type ButtonOptions struct {
	Color *ButtonColor
}
type ButtonOption func(o *ButtonOptions)

type ButtonColor int

const (
	Primary ButtonColor = iota
	Secondary
	Tertiary
	Success
	Danger
	Info
	Warning
)

type model struct {
	base         *app.Base
	Text         string
	opts         ButtonOptions
	style        lipgloss.Style
	styleFocused lipgloss.Style
	KeyMap       KeyMap
}

type KeyMap struct {
	Submit key.Binding
}

type ButtonPressMsg struct {
	ID string
}

func WithVariant(Color ButtonColor) ButtonOption {
	return func(o *ButtonOptions) {
		o.Color = &Color
	}
}

func New(ctx *app.Context, text string, opts ...ButtonOption) model {
	options := ButtonOptions{
		Color: nil, // TODO theming for default value
	}
	for _, opt := range opts {
		opt(&options)
	}

	style := lipgloss.NewStyle()
	if options.Color != nil {
		switch *options.Color {
		case Primary:
			style = style.Foreground(ctx.Styles.Colors.Primary)
		case Secondary:
			style = style.Foreground(ctx.Styles.Colors.Secondary)
		case Tertiary:
			style = style.Foreground(ctx.Styles.Colors.Tertiary)
		case Success:
			style = style.Foreground(ctx.Styles.Colors.Success)
		case Danger:
			style = style.Foreground(ctx.Styles.Colors.Danger)
		case Warning:
			style = style.Foreground(ctx.Styles.Colors.Warning)
		case Info:
			style = style.Foreground(ctx.Styles.Colors.Info)
		}
	}

	styleFocused := style
	if options.Color != nil {
		switch *options.Color {
		case Primary:
			styleFocused = styleFocused.Background(ctx.Styles.Colors.Primary).Foreground(ctx.Styles.Colors.White)
		case Secondary:
			styleFocused = styleFocused.Background(ctx.Styles.Colors.Secondary).Foreground(ctx.Styles.Colors.White)
		case Tertiary:
			styleFocused = styleFocused.Background(ctx.Styles.Colors.Tertiary).Foreground(ctx.Styles.Colors.Black)
		case Success:
			styleFocused = styleFocused.Background(ctx.Styles.Colors.Success).Foreground(ctx.Styles.Colors.Black)
		case Danger:
			styleFocused = styleFocused.Background(ctx.Styles.Colors.Danger).Foreground(ctx.Styles.Colors.White)
		case Warning:
			styleFocused = styleFocused.Background(ctx.Styles.Colors.Warning).Foreground(ctx.Styles.Colors.Black)
		case Info:
			styleFocused = styleFocused.Background(ctx.Styles.Colors.Info).Foreground(ctx.Styles.Colors.Black)
		}
	}

	return model{
		base:         app.New(ctx, app.WithFocusable(true)),
		Text:         text,
		style:        style,
		styleFocused: styleFocused,
		opts:         options,
		KeyMap: KeyMap{
			Submit: key.NewBinding(
				key.WithKeys("enter"),
				key.WithHelp("enter", "submit"),
			),
		},
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

	if m.base.Focused {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, m.KeyMap.Submit):
				return m, func() tea.Msg {
					return ButtonPressMsg{ID: m.base.ID}
				}
			}
		}
	}

	cmd = m.base.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	style := m.style
	if m.base.Focused {
		style = m.styleFocused
	}
	return m.base.Ctx.Zone.Mark(m.base.ID, style.Render("[ "+m.Text+" ]"))
}

func (m model) Base() *app.Base {
	return m.base
}
