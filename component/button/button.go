package button

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/shader"
	"github.com/alexanderbh/bubbleapp/style"
	"github.com/charmbracelet/bubbles/v2/key"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type Options struct {
	Variant ButtonVariant
	Type    ButtonType
	Shader  shader.Shader
	style.Margin
}

type ButtonType int

const (
	Normal ButtonType = iota
	Compact
)

type ButtonVariant int

const (
	Primary ButtonVariant = iota // This is the default it seems
	Secondary
	Tertiary
	Success
	Danger
	Info
	Warning
)

type model[T any] struct {
	base         *app.Base[T]
	Text         string
	opts         *Options
	style        lipgloss.Style
	styleFocused lipgloss.Style
	styleHovered lipgloss.Style
	KeyMap       KeyMap
}

type KeyMap struct {
	Submit key.Binding
}

type ButtonPressMsg struct {
	ID string
}

func New[T any](ctx *app.Context[T], text string, options *Options) *app.Base[T] {

	if options == nil {
		options = &Options{}
	}

	s := lipgloss.NewStyle()
	s = style.ApplyMargin(s, options.Margin)

	if options.Type == Normal {
		s = s.Border(lipgloss.RoundedBorder())
	} else if options.Type == Compact {
		text = "[" + text + "]"
	}

	// TODO: This seems too much code. Is there a better way?
	switch options.Variant {
	case Primary:
		s = s.Foreground(ctx.Styles.Colors.Primary)
	case Secondary:
		s = s.Foreground(ctx.Styles.Colors.Secondary)
	case Tertiary:
		s = s.Foreground(ctx.Styles.Colors.Tertiary)
	case Success:
		s = s.Foreground(ctx.Styles.Colors.Success)
	case Danger:
		s = s.Foreground(ctx.Styles.Colors.Danger)
	case Warning:
		s = s.Foreground(ctx.Styles.Colors.Warning)
	case Info:
		s = s.Foreground(ctx.Styles.Colors.Info)
	}

	styleFocused := s

	if options.Type == Normal {
		switch options.Variant {
		case Primary:
			styleFocused = styleFocused.BorderForeground(ctx.Styles.Colors.Primary).Foreground(ctx.Styles.Colors.Primary)
		case Secondary:
			styleFocused = styleFocused.BorderForeground(ctx.Styles.Colors.Secondary).Foreground(ctx.Styles.Colors.Secondary)
		case Tertiary:
			styleFocused = styleFocused.BorderForeground(ctx.Styles.Colors.Tertiary).Foreground(ctx.Styles.Colors.Tertiary)
		case Success:
			styleFocused = styleFocused.BorderForeground(ctx.Styles.Colors.Success).Foreground(ctx.Styles.Colors.Success)
		case Danger:
			styleFocused = styleFocused.BorderForeground(ctx.Styles.Colors.Danger).Foreground(ctx.Styles.Colors.Danger)
		case Warning:
			styleFocused = styleFocused.BorderForeground(ctx.Styles.Colors.Warning).Foreground(ctx.Styles.Colors.Warning)
		case Info:
			styleFocused = styleFocused.BorderForeground(ctx.Styles.Colors.Info).Foreground(ctx.Styles.Colors.Info)
		}
	} else if options.Type == Compact {
		switch options.Variant {
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

	styleHovered := s

	if options.Type == Normal {
		switch options.Variant {
		case Primary:
			styleHovered = styleFocused.BorderForeground(ctx.Styles.Colors.PrimaryLight)
		case Secondary:
			styleHovered = styleFocused.BorderForeground(ctx.Styles.Colors.SecondaryLight)
		case Tertiary:
			styleHovered = styleFocused.BorderForeground(ctx.Styles.Colors.TertiaryLight)
		case Success:
			styleHovered = styleFocused.BorderForeground(ctx.Styles.Colors.SuccessLight)
		case Danger:
			styleHovered = styleFocused.BorderForeground(ctx.Styles.Colors.DangerLight)
		case Warning:
			styleHovered = styleFocused.BorderForeground(ctx.Styles.Colors.WarningLight)
		case Info:
			styleHovered = styleFocused.BorderForeground(ctx.Styles.Colors.InfoLight)
		}
	} else if options.Type == Compact {
		switch options.Variant {
		case Primary:
			styleHovered = styleFocused.Background(ctx.Styles.Colors.PrimaryLight)
		case Secondary:
			styleHovered = styleFocused.Background(ctx.Styles.Colors.SecondaryLight)
		case Tertiary:
			styleHovered = styleFocused.Background(ctx.Styles.Colors.TertiaryLight)
		case Success:
			styleHovered = styleFocused.Background(ctx.Styles.Colors.SuccessLight)
		case Danger:
			styleHovered = styleFocused.Background(ctx.Styles.Colors.DangerLight)
		case Warning:
			styleHovered = styleFocused.Background(ctx.Styles.Colors.WarningLight)
		case Info:
			styleHovered = styleFocused.Background(ctx.Styles.Colors.InfoLight)
		}
	}

	return model[T]{
		base:         app.New(ctx, app.WithFocusable(true), app.WithShader(options.Shader)),
		Text:         text,
		style:        s,
		styleFocused: styleFocused,
		styleHovered: styleHovered,
		opts:         options,
		KeyMap: KeyMap{
			Submit: key.NewBinding(
				key.WithKeys("enter"),
				key.WithHelp("enter", "submit"),
			),
		},
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

	switch msg := msg.(type) {
	case tea.MouseClickMsg:
		if msg.Button == tea.MouseLeft {
			if m.base.Ctx.Zone.Get(m.Base().ID).InBounds(msg) {
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

func (m model[T]) View() string {
	style := m.style
	if m.base.Focused {
		style = m.styleFocused
	}
	if m.base.Hovered {
		style = m.styleHovered
	}

	return m.base.Ctx.Zone.Mark(m.base.ID, m.base.ApplyShaderWithStyle(m.Text, style))
}

func (m model[T]) Base() *app.Base[T] {
	m.base.Model = m
	return m.base
}
