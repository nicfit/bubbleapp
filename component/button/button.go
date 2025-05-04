package button

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/style"
	"github.com/charmbracelet/bubbles/v2/key"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type Options struct {
	Variant ButtonVariant
	Type    ButtonType
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

type button[T any] struct {
	base         *app.Base[T]
	render       func(ctx *app.Context[T]) string
	onClick      func(ctx *app.Context[T])
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

func New[T any](ctx *app.Context[T], render func(ctx *app.Context[T]) string, onClick func(ctx *app.Context[T]), options *Options, baseOptions ...app.BaseOption) *button[T] {
	if options == nil {
		options = &Options{}
	}
	if baseOptions == nil {
		baseOptions = []app.BaseOption{}
	}

	s := lipgloss.NewStyle()
	s = style.ApplyMargin(s, options.Margin)

	if options.Type == Normal {
		s = s.Border(lipgloss.RoundedBorder())
	} else if options.Type == Compact {
		render = func(ctx *app.Context[T]) string {
			return "⟦" + render(ctx) + "⟧"
		}
	}

	switch options.Variant {
	case Primary:
		s = s.BorderForeground(ctx.Styles.Colors.Primary).Foreground(ctx.Styles.Colors.Primary)
	case Secondary:
		s = s.BorderForeground(ctx.Styles.Colors.Secondary).Foreground(ctx.Styles.Colors.Secondary)
	case Tertiary:
		s = s.BorderForeground(ctx.Styles.Colors.Tertiary).Foreground(ctx.Styles.Colors.Tertiary)
	case Success:
		s = s.BorderForeground(ctx.Styles.Colors.Success).Foreground(ctx.Styles.Colors.Success)
	case Danger:
		s = s.BorderForeground(ctx.Styles.Colors.Danger).Foreground(ctx.Styles.Colors.Danger)
	case Warning:
		s = s.BorderForeground(ctx.Styles.Colors.Warning).Foreground(ctx.Styles.Colors.Warning)
	case Info:
		s = s.BorderForeground(ctx.Styles.Colors.Info).Foreground(ctx.Styles.Colors.Info)
	}

	styleFocused := s

	if options.Type == Normal {
		switch options.Variant {
		case Primary:
			styleFocused = styleFocused.BorderForeground(ctx.Styles.Colors.PrimaryLight).Foreground(ctx.Styles.Colors.White).Background(ctx.Styles.Colors.PrimaryDark)
		case Secondary:
			styleFocused = styleFocused.BorderForeground(ctx.Styles.Colors.SecondaryLight).Foreground(ctx.Styles.Colors.White).Background(ctx.Styles.Colors.SecondaryDark)
		case Tertiary:
			styleFocused = styleFocused.BorderForeground(ctx.Styles.Colors.TertiaryLight).Foreground(ctx.Styles.Colors.Black).Background(ctx.Styles.Colors.TertiaryDark)
		case Success:
			styleFocused = styleFocused.BorderForeground(ctx.Styles.Colors.SuccessLight).Foreground(ctx.Styles.Colors.Black).Background(ctx.Styles.Colors.SuccessDark)
		case Danger:
			styleFocused = styleFocused.BorderForeground(ctx.Styles.Colors.DangerLight).Foreground(ctx.Styles.Colors.White).Background(ctx.Styles.Colors.DangerDark)
		case Warning:
			styleFocused = styleFocused.BorderForeground(ctx.Styles.Colors.WarningLight).Foreground(ctx.Styles.Colors.Black).Background(ctx.Styles.Colors.WarningDark)
		case Info:
			styleFocused = styleFocused.BorderForeground(ctx.Styles.Colors.InfoLight).Foreground(ctx.Styles.Colors.Black).Background(ctx.Styles.Colors.InfoDark)
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

	styleHovered := styleFocused

	if options.Type == Normal {
		switch options.Variant {
		case Primary:
			styleHovered = styleHovered.BorderForeground(ctx.Styles.Colors.PrimaryDark).Foreground(ctx.Styles.Colors.Black).Background(ctx.Styles.Colors.PrimaryLight)
		case Secondary:
			styleHovered = styleHovered.BorderForeground(ctx.Styles.Colors.SecondaryDark).Foreground(ctx.Styles.Colors.Black).Background(ctx.Styles.Colors.SecondaryLight)
		case Tertiary:
			styleHovered = styleHovered.BorderForeground(ctx.Styles.Colors.TertiaryDark).Foreground(ctx.Styles.Colors.Black).Background(ctx.Styles.Colors.TertiaryLight)
		case Success:
			styleHovered = styleHovered.BorderForeground(ctx.Styles.Colors.SuccessDark).Foreground(ctx.Styles.Colors.Black).Background(ctx.Styles.Colors.SuccessLight)
		case Danger:
			styleHovered = styleHovered.BorderForeground(ctx.Styles.Colors.DangerDark).Foreground(ctx.Styles.Colors.Black).Background(ctx.Styles.Colors.DangerLight)
		case Warning:
			styleHovered = styleHovered.BorderForeground(ctx.Styles.Colors.WarningDark).Foreground(ctx.Styles.Colors.Black).Background(ctx.Styles.Colors.WarningLight)
		case Info:
			styleHovered = styleHovered.BorderForeground(ctx.Styles.Colors.InfoDark).Foreground(ctx.Styles.Colors.Black).Background(ctx.Styles.Colors.InfoLight)
		}
	} else if options.Type == Compact {
		switch options.Variant {
		case Primary:
			styleHovered = styleHovered.Background(ctx.Styles.Colors.PrimaryLight)
		case Secondary:
			styleHovered = styleHovered.Background(ctx.Styles.Colors.SecondaryLight)
		case Tertiary:
			styleHovered = styleHovered.Background(ctx.Styles.Colors.TertiaryLight)
		case Success:
			styleHovered = styleHovered.Background(ctx.Styles.Colors.SuccessLight)
		case Danger:
			styleHovered = styleHovered.Background(ctx.Styles.Colors.DangerLight)
		case Warning:
			styleHovered = styleHovered.Background(ctx.Styles.Colors.WarningLight)
		case Info:
			styleHovered = styleHovered.Background(ctx.Styles.Colors.InfoLight)
		}
	}

	return &button[T]{
		base:         app.NewBase[T](append([]app.BaseOption{app.WithFocusable(true)}, baseOptions...)...),
		render:       render,
		onClick:      onClick,
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
	}
}

func (m *button[T]) Render(ctx *app.Context[T]) string {
	style := m.style
	if ctx.Focused == m {
		style = m.styleFocused
	}
	if m.base.Hovered {
		style = m.styleHovered
	}

	rendered := m.base.ApplyShaderWithStyle(m.render(ctx), style)

	return app.RegisterMouse(ctx, m.base.ID, m, rendered)
}

func (m *button[T]) Update(ctx *app.Context[T], msg tea.Msg) {
	switch msg := msg.(type) {
	case tea.MouseClickMsg:
		if msg.Button == tea.MouseLeft {
			m.onClick(ctx)
		}
	}
}

func (m *button[T]) Children(ctx *app.Context[T]) []app.Fc[T] {
	return nil
}
func (m *button[T]) Base() *app.Base[T] {
	return m.base
}
