package app

import (
	"image/color"

	"github.com/alexanderbh/bubbleapp/style"
	zone "github.com/alexanderbh/bubblezone/v2"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type Context[T any] struct {
	root    Fc[T]
	UIState *StateStore
	Zone    *zone.Manager
	ZoneMap map[string]Fc[T]

	Styles          *style.Styles
	BackgroundColor color.Color
	Width           int
	Height          int
	LayoutPhase     bool
	Data            *T
	teaProgram      *tea.Program

	id   *idContext[T]
	Tick *tickState[T]
}

func NewContext[T any](data *T) *Context[T] {
	return &Context[T]{
		Zone:    zone.New(),
		ZoneMap: make(map[string]Fc[T]),
		UIState: NewStateStore(),
		Styles:  style.DefaultStyles(),
		Data:    data,
		id:      newIDContext[T](),
		Tick:    &tickState[T]{},
	}
}

type InvalidateMsg struct{}

// Invalidates the UI and forces a re-render.
// Requires a tea.Program to be set with app.SetTeaProgram.
// This is useful for performance optimizations where a tick
// is too expensive.
func (ctx *Context[T]) Update() {
	if ctx.teaProgram == nil {
		panic("teaProgram is nil. Cannot update manually.")
	}
	go ctx.teaProgram.Send(InvalidateMsg{})
}

func (ctx *Context[T]) AddCmd(cmd tea.Cmd) {
	if cmd == nil {
		return
	}

	go ctx.teaProgram.Send(cmd())
}

// Quit signals the application to stop, ensuring cleanup like stopping active timers.
func (ctx *Context[T]) Quit() {
	if ctx.Tick != nil {
		ctx.Tick.StopActiveTimer()
	}
	go ctx.teaProgram.Quit()
}
