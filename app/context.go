package app

import (
	"reflect"
	"runtime"
	"strings"

	"github.com/alexanderbh/bubbleapp/style"
	zone "github.com/alexanderbh/bubblezone/v2"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type Ctx struct {
	UIState          *uiStateContext
	Zone             *zone.Manager
	ZoneMap          map[string]*instanceContext
	teaProgram       *tea.Program
	Styles           *style.Styles
	id               *idContext
	Tick             *tickState[any]
	collectorStack   []*outputCollector
	componentContext *fcInstanceContext
	useEffectCounter int
	useStateCounter  int

	// Layout
	layoutPhase   layoutPhase
	layoutManager *layoutManager
}

func NewCtx() *Ctx {
	return &Ctx{
		UIState:          NewUIStateContext(),
		Zone:             zone.New(),
		ZoneMap:          make(map[string]*instanceContext),
		Styles:           style.DefaultStyles(),
		id:               newIdContext(),
		Tick:             &tickState[any]{},
		collectorStack:   []*outputCollector{},
		componentContext: newInstanceContext(),
		layoutManager:    newLayoutManager(),
	}
}

// Render a functional component with the given props.
// This function is responsible for managing the lifecycle of the component,
// including state management, effect handling, and ID management.
func (c *Ctx) Render(fc FC, props Props) string {
	id := c.id.push(getFuncName(fc))
	defer c.id.pop()

	c.layoutManager.addComponent(id, fc, props)
	defer c.layoutManager.pop()

	c.id.ids = append(c.id.ids, id)

	c.componentContext.set(id, fc, props)

	c.useStateCounter = 0
	c.useEffectCounter = 0

	output := fc(c, props)

	// If there is an active output collector, append the output to it
	if len(c.collectorStack) > 0 {
		currentCollector := c.collectorStack[len(c.collectorStack)-1]
		currentCollector.outputs = append(currentCollector.outputs, output)
	}

	return output
}

type outputCollector struct {
	outputs []string
}
type InvalidateMsg struct{}

// Helper to get function name (can be fragile - consider other approach)
func getFuncName(fn interface{}) string {
	// Ensure fn is a function
	v := reflect.ValueOf(fn)
	if v.Kind() != reflect.Func {
		return "unknownComponent"
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
func (ctx *Ctx) Update() {
	if ctx.teaProgram == nil {
		panic("teaProgram is nil. Cannot update manually.")
	}
	go ctx.teaProgram.Send(InvalidateMsg{})
}

// Quit signals the application to stop, ensuring cleanup like stopping active timers.
func (ctx *Ctx) Quit() {
	if ctx.Tick != nil {
		ctx.Tick.StopActiveTimer()
	}
	go ctx.teaProgram.Quit()
}
