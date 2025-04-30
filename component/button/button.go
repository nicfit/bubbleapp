package button

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/charmbracelet/bubbles/v2/key"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type Options struct {
	Variant ButtonVariant
}
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

func New[T any](ctx *app.Context[T], text string, options *Options) model[T] {

	if options == nil {
		options = &Options{}
	}

	style := lipgloss.NewStyle()

	// TODO: This seems too much code. Is there a better way?
	switch options.Variant {
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

	styleFocused := style
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

	styleHovered := style
	switch options.Variant {
	case Primary:
		styleHovered = styleFocused.Background(ctx.Styles.Colors.PrimaryDark)
	case Secondary:
		styleHovered = styleFocused.Background(ctx.Styles.Colors.SecondaryDark)
	case Tertiary:
		styleHovered = styleFocused.Background(ctx.Styles.Colors.TertiaryDark)
	case Success:
		styleHovered = styleFocused.Background(ctx.Styles.Colors.SuccessDark)
	case Danger:
		styleHovered = styleFocused.Background(ctx.Styles.Colors.DangerDark)
	case Warning:
		styleHovered = styleFocused.Background(ctx.Styles.Colors.WarningDark)
	case Info:
		styleHovered = styleFocused.Background(ctx.Styles.Colors.InfoDark)
	}

	return model[T]{
		base:         app.New(ctx, app.WithFocusable(true)),
		Text:         text,
		style:        style,
		styleFocused: styleFocused,
		styleHovered: styleHovered,
		opts:         options,
		KeyMap: KeyMap{
			Submit: key.NewBinding(
				key.WithKeys("enter"),
				key.WithHelp("enter", "submit"),
			),
		}}
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
	return m.base.Ctx.Zone.Mark(m.base.ID, style.Render("[ "+m.Text+" ]"))
}

func (m model[T]) Base() *app.Base[T] {
	m.base.Model = m
	return m.base
}
