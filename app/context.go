package app

import (
	"reflect"
	"runtime"
	"strings"

	"github.com/alexanderbh/bubbleapp/style"
	zone "github.com/alexanderbh/bubblezone/v2"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type Ctx struct {
	UIState          *uiStateContext
	zone             *zone.Manager
	zoneMap          map[string]*instanceContext
	teaProgram       *tea.Program
	Styles           *style.Styles
	id               *idContext
	tick             *tickState[any]
	invalidate       bool
	collectorStack   []*outputCollector
	componentContext *fcInstanceContext
	useEffectCounter map[string]int
	useStateCounter  map[string]int

	// Layout
	LayoutPhase   layoutPhase
	layoutManager *layoutManager
}

func NewCtx() *Ctx {
	return &Ctx{
		UIState:          NewUIStateContext(),
		zone:             zone.New(),
		zoneMap:          make(map[string]*instanceContext),
		Styles:           style.DefaultStyles(),
		id:               newIdContext(),
		tick:             &tickState[any]{},
		collectorStack:   []*outputCollector{},
		componentContext: newInstanceContext(),
		layoutManager:    newLayoutManager(),
		useEffectCounter: make(map[string]int),
		useStateCounter:  make(map[string]int),
	}
}

func (c *Ctx) initView() {
	c.id.initIDCollections()
	c.id.initPath()
	c.tick.init()
	c.zoneMap = make(map[string]*instanceContext)
	for _, cs := range c.componentContext.ctxs {
		cs.keyHandlers = make([]KeyHandler, 0)
		cs.mouseHandlers = make([]MouseHandler, 0)
	}
}

// Render a functional component with the given props.
// This function is responsible for managing the lifecycle of the component,
// including state management, effect handling, and ID management.
func (c *Ctx) Render(fc FC, props Props) string {
	id := c.id.push(getFuncName(fc))
	defer c.id.pop()

	var node *ComponentNode
	if c.LayoutPhase == LayoutPhaseIntrincintWidth {
		node = c.layoutManager.addComponent(id, fc, props)
		defer c.layoutManager.pop()
	} else {
		node = c.layoutManager.getComponent(id)
	}

	c.id.ids = append(c.id.ids, id)

	c.componentContext.set(id, fc, props)

	c.useStateCounter[id] = 0
	c.useEffectCounter[id] = 0

	output := fc(c, props)

	// If there is an active output collector, append the output to it
	if len(c.collectorStack) > 0 {
		currentCollector := c.collectorStack[len(c.collectorStack)-1]
		currentCollector.outputs = append(currentCollector.outputs, output)
	}

	if node != nil {
		node.LastRender = output
		if c.LayoutPhase == LayoutPhaseIntrincintWidth {
			if node.Parent == nil {
				c.UIState.setWidth(id, c.layoutManager.width)
			} else if !node.Layout.GrowX {
				c.UIState.setWidth(id, lipgloss.Width(output))
			}
		}
		if c.LayoutPhase == LayoutPhaseIntrincintHeight {
			if node.Parent == nil {
				c.UIState.setHeight(id, c.layoutManager.height)
			} else if !node.Layout.GrowY {
				c.UIState.setHeight(id, lipgloss.Height(output))
			}
		}
	}

	return output
}

func (c *Ctx) MouseZone(id string, content string) string {
	instance, _ := c.componentContext.get(id)
	c.zoneMap[id] = instance
	return c.zone.Mark(id, content)
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
func (c *Ctx) Update() {
	if c.teaProgram == nil {
		panic("teaProgram is nil. Cannot update manually.")
	}
	if !c.invalidate {
		if c.teaProgram != nil {
			go c.teaProgram.Send(InvalidateMsg{})
		}
	}
	c.invalidate = true
}

// Quit signals the application to stop, ensuring cleanup like stopping active timers.
func (ctx *Ctx) Quit() {
	if ctx.tick != nil {
		ctx.tick.StopActiveTimer()
	}
	go ctx.teaProgram.Quit()
}
