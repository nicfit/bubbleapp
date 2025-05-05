package app

import (
	"time"

	"github.com/alexanderbh/bubbleapp/style"
	zone "github.com/alexanderbh/bubblezone/v2"
	tea "github.com/charmbracelet/bubbletea/v2"
)

const (
	// Global FPS for Ticks. Hardcoded for now. Could be dynamic perhaps.
	FPS = time.Second / 12
)

type Fc[T any] interface {
	Render(ctx *Context[T]) string
	Update(ctx *Context[T], msg tea.Msg)
	Base() *Base
	Children(ctx *Context[T]) []Fc[T]
}

type AppOptions struct {
	TickFPS time.Duration
}
type AppOption func(*AppOptions)

func WithTickFPS(fps time.Duration) AppOption {
	return func(o *AppOptions) {
		o.TickFPS = fps
	}
}

type App[T any] struct {
	render  func(ctx *Context[T]) Fc[T]
	ctx     *Context[T]
	tickFPS time.Duration
}

func NewApp[T any](ctx *Context[T], render func(ctx *Context[T]) Fc[T], options ...AppOption) *App[T] {
	if options == nil {
		options = []AppOption{}
	}
	if ctx.ZoneMap == nil {
		ctx.ZoneMap = make(map[string]Fc[T])
	}
	if ctx.Styles == nil {
		ctx.Styles = style.DefaultStyles()
	}
	if ctx.Cmds == nil {
		ctx.Cmds = &[]tea.Cmd{}
	}

	opts := &AppOptions{
		TickFPS: FPS,
	}
	for _, opt := range options {
		opt(opts)
	}

	return &App[T]{
		render:  render,
		ctx:     ctx,
		tickFPS: opts.TickFPS,
	}
}

func (a *App[T]) Init() tea.Cmd {
	var (
		cmds []tea.Cmd
	)

	if a.tickFPS > 0 {
		cmds = append(cmds, tickCommand(a.tickFPS))
	}

	return tea.Batch(cmds...)
}

func (a *App[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	if len(*a.ctx.Cmds) > 0 {
		cmds = *a.ctx.Cmds
		a.ctx.Cmds = &[]tea.Cmd{}
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		// These are global keys. Is this what we want?
		switch msg.String() {
		case "ctrl+c":
			return a, tea.Quit
		case "tab":
			a.ctx.FocusNextCmd(a.ctx.root)
			return a, nil
		case "shift+tab":
			a.ctx.FocusPrevCmd(a.ctx.root)
			return a, nil
		}
		// Propagate key msg to the focused component
		if a.ctx.UIState.Focused != "" {
			focused := a.ctx.IDMap[a.ctx.UIState.Focused]
			if focused != nil {
				focused.Update(a.ctx, msg)
			}
		}
	case tea.WindowSizeMsg:
		a.ctx.Width = msg.Width
		a.ctx.Height = msg.Height
	case tea.MouseMsg:
		a.ctx.Zone.AnyInBounds(a, msg)
	case zone.MsgZoneInBounds:
		foundZone := a.ctx.ZoneMap[a.ctx.Zone.GetReverse(msg.Zone.Id)]
		if foundZone != nil {
			foundZone.Update(a.ctx, msg.Event)
		}
	case TickMsg:
		// Propagate tick to all components (used for Dynamic Shaders for now)
		Visit(a.ctx.root, 0, nil, a.ctx, tickVisitor, PreOrder)
		if a.tickFPS > 0 {
			cmds = append(cmds, tickCommand(a.tickFPS))
		}
	}

	return a, tea.Batch(cmds...)
}

func (a *App[T]) View() string {
	a.ctx.idPath = []string{"root"}
	a.ctx.idPathCount = make(map[string]int)

	rendered := a.render(a.ctx)
	a.ctx.root = rendered
	a.Layout()
	return a.ctx.Zone.Scan(rendered.Render(a.ctx))
}
