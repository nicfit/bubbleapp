package app

import (
	"time"

	"github.com/alexanderbh/bubbleapp/style"
	zone "github.com/alexanderbh/bubblezone/v2"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type Fc[T any] interface {
	Render(ctx *Context[T]) string
	Update(ctx *Context[T], msg tea.Msg) bool
	Base() *Base
	Children(ctx *Context[T]) []Fc[T]
}

type AppOptions struct {
	TickFPS time.Duration
}
type AppOption func(*AppOptions)

// This registers a constant global tick
// Use with caution. Most components should have registered
// their own tick listener.
// Use ctx.Update() invalidate the UI form the outside.
func WithTickFPS(fps time.Duration) AppOption {
	return func(o *AppOptions) {
		o.TickFPS = fps
	}
}

type App[T any] struct {
	scaffold func(ctx *Context[T]) Fc[T]
	ctx      *Context[T]
	tickFPS  time.Duration
}

func NewApp[T any](ctx *Context[T], scaffold func(ctx *Context[T]) Fc[T], options ...AppOption) *App[T] {
	if options == nil {
		options = []AppOption{}
	}
	if ctx.ZoneMap == nil {
		ctx.ZoneMap = make(map[string]Fc[T])
	}
	if ctx.Styles == nil {
		ctx.Styles = style.DefaultStyles()
	}

	opts := &AppOptions{
		TickFPS: 0,
	}
	for _, opt := range options {
		opt(opts)
	}

	return &App[T]{
		scaffold: scaffold,
		ctx:      ctx,
		tickFPS:  opts.TickFPS,
	}
}

func (a *App[T]) SetTeaProgram(p *tea.Program) {
	a.ctx.teaProgram = p
}

func (a *App[T]) Init() tea.Cmd {
	if a.ctx.teaProgram == nil {
		panic("teaProgram is nil. Set the tea.Program with app.SetTeaProgram(p).")
	}
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

	switch msg := msg.(type) {
	case tea.KeyMsg:
		handledKeyMsg := a.propagatedFocused(msg)
		if !handledKeyMsg {
			// These are global keys. Is this what we want?
			switch msg.String() {
			case "ctrl+c":
				a.ctx.Quit()
				return a, nil // Use the new Quit method
			case "tab":
				a.ctx.FocusNextCmd(a.ctx.root)
				return a, nil
			case "shift+tab":
				a.ctx.FocusPrevCmd(a.ctx.root)
				return a, nil
			}
		}
		return a, tea.Batch(cmds...)

	case tea.WindowSizeMsg:
		a.ctx.Width = msg.Width
		a.ctx.Height = msg.Height
		return a, tea.Batch(cmds...)
	case tea.MouseMsg:
		a.ctx.Zone.AnyInBounds(a, msg)
		return a, tea.Batch(cmds...)
	case zone.MsgZoneInBounds:
		foundZone := a.ctx.ZoneMap[a.ctx.Zone.GetReverse(msg.Zone.Id)]
		if foundZone != nil {
			foundZone.Update(a.ctx, msg.Event)
		}
		return a, tea.Batch(cmds...)
	case TickMsg:
		now := msg.OccurredAt
		for _, listener := range *a.ctx.Tick.tickListeners {
			l := a.ctx.id.getNode(listener.id)
			if l != nil {
				lastTick, ok := a.ctx.Tick.lastTickTimes[listener.id]
				if !ok || now.Sub(lastTick) >= listener.interval {
					l.Update(a.ctx, msg)
					a.ctx.Tick.lastTickTimes[listener.id] = now
				}
			}
		}
		if a.tickFPS > 0 {
			cmds = append(cmds, tickCommand(a.tickFPS))
		}
		return a, tea.Batch(cmds...)
	}

	a.propagatedFocused(msg)

	return a, tea.Batch(cmds...)
}

func (a *App[T]) propagatedFocused(msg tea.Msg) bool {
	if a.ctx.UIState.Focused != "" {
		focused := a.ctx.id.getNode(a.ctx.UIState.Focused)
		if focused != nil {
			handled := focused.Update(a.ctx, msg)
			return handled
		}
	}
	return false
}

func (a *App[T]) View() string {
	a.ctx.id.initPath()
	a.ctx.Tick.init()

	root := a.scaffold(a.ctx)

	a.ctx.Tick.createTimer(a.ctx)

	a.ctx.root = root
	a.Layout()
	return a.ctx.Zone.Scan(root.Render(a.ctx))
}
