package app

import (
	"time"

	"github.com/alexanderbh/bubbleapp/style"
	zone "github.com/alexanderbh/bubblezone/v2"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type Fc[T any] interface {
	Render(ctx *Context[T]) string
	Update(ctx *Context[T], msg tea.Msg)
	Base() *Base[T]
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
	root    Fc[T]
	ctx     *Context[T]
	tickFPS time.Duration
}

func NewApp[T any](ctx *Context[T], root Fc[T], options ...AppOption) *App[T] {
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
		TickFPS: time.Second / 12,
	}
	for _, opt := range options {
		opt(opts)
	}

	return &App[T]{
		root:    root,
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
		switch msg.String() {
		case "ctrl+c":
			return a, tea.Quit
		case "tab":
			a.ctx.FocusNextCmd(a.root)
			return a, nil
		case "shift+tab":
			a.ctx.FocusPrevCmd(a.root)
			return a, nil
		}
	case tea.WindowSizeMsg:
		a.root.Base().Width = msg.Width
		a.root.Base().Height = msg.Height
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
		if a.tickFPS > 0 {
			cmds = append(cmds, tickCommand(a.tickFPS))
		}
	}

	return a, tea.Batch(cmds...)
}

func (a *App[T]) View() string {
	a.Layout()
	return a.ctx.Zone.Scan(a.root.Render(a.ctx))
}
