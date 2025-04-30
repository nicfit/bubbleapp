package loader

import (
	"image/color"
	"time"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/charmbracelet/lipgloss/v2"

	tea "github.com/charmbracelet/bubbletea/v2"
)

type model[T any] struct {
	base         *app.Base[T]
	options      options
	styleSpinner lipgloss.Style
	styleText    lipgloss.Style
	spinner      Spinner
	frame        int
	lastTick     time.Time // Track last frame update
}

type options struct {
	text                string
	textColor           *color.Color
	textBackgroundColor *color.Color
	color               *color.Color
}

type option func(*options)

func WithText(text string) option {
	return func(o *options) {
		o.text = text
	}
}
func WithTextColor(color color.Color) option {
	return func(o *options) {
		o.textColor = &color
	}
}
func WithTextBackgroundColor(color color.Color) option {
	return func(o *options) {
		o.textBackgroundColor = &color
	}
}

func WithColor(color color.Color) option {
	return func(o *options) {
		o.color = &color
	}
}

func New[T any](ctx *app.Context[T], variant Spinner, opts ...option) model[T] {
	color := ctx.Styles.Colors.Info
	options := options{
		text:  "",
		color: &color,
	}
	for _, opt := range opts {
		opt(&options)
	}

	styleText := lipgloss.NewStyle()
	styleSpinner := lipgloss.NewStyle()

	if options.textColor != nil {
		styleText = styleText.Foreground(*options.textColor)
	}
	if options.textBackgroundColor != nil {
		styleText = styleText.Background(*options.textBackgroundColor)
	}
	if options.color != nil {
		styleSpinner = styleSpinner.Foreground(*options.color)
	}
	return model[T]{
		base:         app.New(ctx),
		spinner:      variant,
		options:      options,
		styleText:    styleText,
		styleSpinner: styleSpinner,
		frame:        0,
		lastTick:     time.Now(), // Initialize lastTick
	}
}

func (m model[T]) Init() tea.Cmd {
	return nil
}

func (m model[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case app.TickMsg:
		// Only update frame if enough time has passed according to spinner FPS
		now := time.Now()
		if now.Sub(m.lastTick) >= m.spinner.FPS {
			m.frame++
			if m.frame >= len(m.spinner.Frames) {
				m.frame = 0
			}
			m.lastTick = now
		}
		return m, nil
	default:
		return m, nil
	}
}

func (m model[T]) View() string {
	text := m.options.text
	if text != "" {
		text = " " + text
	}
	return m.styleSpinner.Render(m.spinner.Frames[m.frame]) + m.styleText.Render(text)
}

func (m model[T]) Base() *app.Base[T] {
	m.base.Model = m
	return m.base
}

type Spinner struct {
	Frames []string
	FPS    time.Duration
}

var (
	Line = Spinner{
		Frames: []string{"|", "/", "-", "\\"},
		FPS:    time.Second / 6, //nolint:mnd
	}
	Dot = Spinner{
		Frames: []string{"⣾ ", "⣽ ", "⣻ ", "⢿ ", "⡿ ", "⣟ ", "⣯ ", "⣷ "},
		FPS:    time.Second / 12, //nolint:mnd
	}
	MiniDot = Spinner{
		Frames: []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
		FPS:    time.Second / 12, //nolint:mnd
	}
	Jump = Spinner{
		Frames: []string{"⢄", "⢂", "⢁", "⡁", "⡈", "⡐", "⡠"},
		FPS:    time.Second / 12, //nolint:mnd
	}
	Pulse = Spinner{
		Frames: []string{"█", "▓", "▒", "░"},
		FPS:    time.Second / 6, //nolint:mnd
	}
	Points = Spinner{
		Frames: []string{"∙∙∙", "●∙∙", "∙●∙", "∙∙●"},
		FPS:    time.Second / 6, //nolint:mnd
	}
	Meter = Spinner{
		Frames: []string{
			"▰▱▱",
			"▱▰▱",
			"▱▱▰",
			"▱▰▱",
		},
		FPS: time.Second / 6, //nolint:mnd
	}
	Ellipsis = Spinner{
		Frames: []string{"", ".", "..", "..."},
		FPS:    time.Second / 3, //nolint:mnd
	}
)
