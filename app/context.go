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
	contextValues    map[uint64][]any // Added for Context API

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
		contextValues:    make(map[uint64][]any), // Initialize contextValues
	}
}

func (c *Ctx) initView() {
	c.id.initIDCollections()
	c.id.initPath()
	c.tick.init()
	c.layoutManager.componentTree = newComponentTree()
	c.collectorStack = []*outputCollector{}
	c.zoneMap = make(map[string]*instanceContext)
	// Is this the right way to clean up here?
	for _, cs := range c.componentContext.ctxs {
		cs.keyHandlers = make([]KeyHandler, 0)
		cs.mouseHandlers = make([]MouseHandler, 0)
		cs.messageHandlers = make([]MsgHandler, 0)
		cs.onFocused = nil
	}
	// c.contextValues = make(map[uint64][]any) // No need to reset here, should persist across renders unless specifically cleared
}

// Render a functional component with the given props.
// This function is responsible for managing the lifecycle of the component,
// including state management, effect handling, and ID management.
func (c *Ctx) Render(fc FC, props Props) string {
	id := c.id.push(getKeyName(fc, props))
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
				if output == "" {
					c.UIState.setHeight(id, 0)
				} else {
					c.UIState.setHeight(id, lipgloss.Height(output))
				}
			}
		}
	}

	return output
}

// MouseZone creates a mouse zone for the given content.
// The ID of the zone is the components ID.
func (c *Ctx) MouseZone(content string) string {
	id := c.id.getID()
	instance, _ := c.componentContext.get(id)
	c.zoneMap[id] = instance
	return c.zone.Mark(id, content)
}

// MouseZoneChild creates a mouse zone for child (sub part) of a component.
// The ID of the zone is the components ID + "###" + childID.
// MouseHandlers will receive the childID extracted from the ID mentioned above.
func (c *Ctx) MouseZoneChild(childID string, content string) string {
	id := c.id.getID()
	return c.zone.Mark(id+"###"+childID, content)
}

type outputCollector struct {
	outputs []string
}
type InvalidateMsg struct{}

// Helper to get function name (can be fragile - consider other approach)
func getKeyName(fn interface{}, props Props) string {
	// Ensure fn is a function
	v := reflect.ValueOf(fn)
	if v.Kind() != reflect.Func {
		panic("fn is not a function")
	}
	fullName := runtime.FuncForPC(v.Pointer()).Name()
	parts := strings.Split(fullName, ".")
	name := parts[len(parts)-1]
	// Clean up common anonymous function suffixes like ".func1"
	name = strings.Split(name, ".")[0]
	if name == "" {
		panic("function name is empty")
	}
	var (
		key    string
		hasKey bool
	)
	if props != nil {
		vProps := reflect.ValueOf(props)
		// Dereference pointer if it's a pointer
		if vProps.Kind() == reflect.Ptr {
			vProps = vProps.Elem()
		}

		switch vProps.Kind() {
		case reflect.Struct:
			// Handle structs
			keyField := vProps.FieldByName("Key")
			if keyField.IsValid() && keyField.Kind() == reflect.String {
				key = keyField.String()
				hasKey = key != ""
			}
		}
	}
	if hasKey {
		return name + "{" + key + "}"
	} else {
		return name
	}
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

func (c *Ctx) ExecuteCmd(cmd tea.Cmd) {
	if c.teaProgram == nil {
		panic("teaProgram is nil. Cannot execute command.")
	}
	if cmd != nil {
		go c.teaProgram.Send(cmd())
	}
}

// Quit signals the application to stop, ensuring cleanup like stopping active timers.
func (ctx *Ctx) Quit() {
	if ctx.tick != nil {
		ctx.tick.StopActiveTimer()
	}
	go ctx.teaProgram.Quit()
}

// pushContextValue adds a value to the stack for a given context ID.
func (c *Ctx) PushContextValue(contextID uint64, value any) {
	c.contextValues[contextID] = append(c.contextValues[contextID], value)
}

// popContextValue removes the top value from the stack for a given context ID.
func (c *Ctx) PopContextValue(contextID uint64) {
	stack, ok := c.contextValues[contextID]
	if !ok || len(stack) == 0 {
		// This should ideally not happen if push/pop are balanced.
		// Consider logging an error or panicking if it's a critical issue.
		return
	}
	c.contextValues[contextID] = stack[:len(stack)-1]
	if len(c.contextValues[contextID]) == 0 {
		delete(c.contextValues, contextID)
	}
}

// getContextValue retrieves the current value for a given context ID from the top of its stack.
func (c *Ctx) GetContextValue(contextID uint64) (any, bool) {
	stack, ok := c.contextValues[contextID]
	if !ok || len(stack) == 0 {
		return nil, false
	}
	return stack[len(stack)-1], true
}
