package app

import (
	"image/color"
	"sync"

	"github.com/alexanderbh/bubbleapp/style"
	zone "github.com/alexanderbh/bubblezone/v2"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type Context[T any] struct {
	root            Fc[T]
	UIState         *StateStore
	Zone            *zone.Manager
	ZoneMap         map[string]Fc[T]
	cmds            *[]tea.Cmd
	cmdMutex        sync.Mutex
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
		cmds:    &[]tea.Cmd{},
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
	ctx.teaProgram.Send(InvalidateMsg{})
}

func (ctx *Context[T]) AddCmd(cmd tea.Cmd) {
	ctx.cmdMutex.Lock()
	defer ctx.cmdMutex.Unlock()
	if ctx.cmds == nil {
		ctx.cmds = &[]tea.Cmd{}
	}
	*ctx.cmds = append(*ctx.cmds, cmd)
}

func (ctx *Context[T]) Quit() {
	ctx.AddCmd(tea.Quit)
}
