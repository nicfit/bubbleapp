package loader

import (
	"image/color"
	"time"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/charmbracelet/lipgloss/v2"

	tea "github.com/charmbracelet/bubbletea/v2"
)

type uiState struct {
	frame    int
	lastTick time.Time
}

type loader[T any] struct {
	base         *app.Base
	render       func(ctx *app.Context[T]) string
	options      *Options
	styleSpinner lipgloss.Style
	styleText    lipgloss.Style
	spinner      Spinner
}

type Options struct {
	TextColor           color.Color
	TextBackgroundColor color.Color
	Color               color.Color
}

func New[T any](ctx *app.Context[T], variant Spinner, text string, options *Options, baseOptions ...app.BaseOption) *loader[T] {
	return NewDynamic(ctx, variant, func(ctx *app.Context[T]) string {
		return text
	}, options, baseOptions...)
}

func NewWithoutText[T any](ctx *app.Context[T], variant Spinner, options *Options, baseOptions ...app.BaseOption) *loader[T] {
	return NewDynamic(ctx, variant, func(ctx *app.Context[T]) string {
		return ""
	}, options, baseOptions...)
}

func NewDynamic[T any](ctx *app.Context[T], variant Spinner, render func(ctx *app.Context[T]) string, options *Options, baseOptions ...app.BaseOption) *loader[T] {
	if options == nil {
		options = &Options{}
	}
	if baseOptions == nil {
		baseOptions = []app.BaseOption{}
	}
	if options.Color == nil {
		options.Color = ctx.Styles.Colors.WarningLight
	}

	styleText := lipgloss.NewStyle().Padding(0, 0, 0, 1) // Padding left 1
	styleSpinner := lipgloss.NewStyle()

	if options.TextColor != nil {
		styleText = styleText.Foreground(options.TextColor)
	}
	if options.TextBackgroundColor != nil {
		styleText = styleText.Background(options.TextBackgroundColor)
	}
	if options.Color != nil {
		styleSpinner = styleSpinner.Foreground(options.Color)
	}

	base, cleanup := app.NewBase(ctx, "loader", baseOptions...)
	defer cleanup()

	ctx.Tick.RegisterTickListener(variant.Interval, base.ID)

	return &loader[T]{
		base:         base,
		render:       render,
		spinner:      variant,
		options:      options,
		styleText:    styleText,
		styleSpinner: styleSpinner,
	}
}

func (m *loader[T]) Update(ctx *app.Context[T], msg tea.Msg) {
	switch msg.(type) {
	case app.TickMsg:
		// Only update frame if enough time has passed according to spinner FPS
		// Not sure if this logic is correct.
		now := time.Now()
		uiState := m.getState(ctx)
		if now.Sub(uiState.lastTick) >= m.spinner.Interval {
			uiState.frame++
			if uiState.frame >= len(m.spinner.Frames) {
				uiState.frame = 0
			}
			uiState.lastTick = now
		}
		return
	}
}

func (m *loader[T]) Render(ctx *app.Context[T]) string {
	text := m.render(ctx)
	return m.styleSpinner.Render(m.spinner.Frames[m.getState(ctx).frame]) + m.styleText.Render(text)
}

func (m *loader[T]) Children(ctx *app.Context[T]) []app.Fc[T] {
	return nil
}
func (m *loader[T]) Base() *app.Base {
	return m.base
}

func (m *loader[T]) getState(ctx *app.Context[T]) *uiState {
	state := app.GetUIState[T, uiState](ctx, m.base.ID)
	if state == nil {
		state = &uiState{
			frame:    0,
			lastTick: time.Now(),
		}
		app.SetUIState(ctx, m.base.ID, state)
	}
	return state
}
