package app

import (
	"image/color"
	"reflect"
	"runtime"
	"strings"

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

// Helper to get function name (can be fragile, use with caution)
func getFuncName(fn interface{}) string {
	// Ensure fn is a function
	v := reflect.ValueOf(fn)
	if v.Kind() != reflect.Func {
		return "unknownComponent" // Or panic, or handle error
	}
	fullName := runtime.FuncForPC(v.Pointer()).Name()
	parts := strings.Split(fullName, ".")
	name := parts[len(parts)-1]
	// Clean up common anonymous function suffixes like ".func1"
	name = strings.Split(name, ".")[0]
	if name == "" {
		return "anonymousComponent"
	}
	return name
}

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
