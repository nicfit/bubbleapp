package app

import (
	"image/color"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/alexanderbh/bubbleapp/style"
	zone "github.com/alexanderbh/bubblezone/v2"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type Ctx struct {
	UIState       *uiStateContext
	zone          *zone.Manager
	zoneMap       map[string]*C
	teaProgram    *tea.Program
	Theme         *style.AppTheme
	id            *idContext
	tick          *tickState[any]
	invalidate    bool
	components    map[string]*C
	ids           []string
	contextValues map[uint64][]any // Added for Context API

	CurrentBg color.Color
	// Layout
	LayoutPhase   layoutPhase
	root          *C
	layoutManager *layoutManager

	Cursor *tea.Cursor
}

func NewCtx() *Ctx {
	return &Ctx{
		UIState:       NewUIStateContext(),
		zone:          zone.New(),
		zoneMap:       make(map[string]*C),
		Theme:         style.NewDefaultAppTheme(),
		id:            newIdContext(),
		tick:          &tickState[any]{},
		components:    make(map[string]*C),
		ids:           make([]string, 0),
		layoutManager: newLayoutManager(),
		contextValues: make(map[uint64][]any),
	}
}

func (c *Ctx) RenderWithName(fn func(c *Ctx, props Props) string, props Props, name string) *C {
	id := c.id.push(name)
	defer c.id.pop()

	var comp *C = c.components[id]
	if comp == nil || c.LayoutPhase == LayoutPhaseIntrincintWidth {
		comp = c.initComponent(id, props)
	}

	c.ids = append(c.ids, id)

	c.layoutManager.addComponent(comp)
	defer c.layoutManager.pop(c, comp)
	if c.root == nil {
		c.root = comp
	}

	comp.useStateCounter = 0
	comp.useEffectCounter = 0

	// FC now returns a string, not Component
	outputStr := fn(c, props)

	comp.content = outputStr

	return comp
}

// Render a functional component with the given props.
// This function is responsible for managing the lifecycle of the component,
// including state management, effect handling, and ID management.
func (c *Ctx) Render(fn func(c *Ctx, props Props) string, props Props) *C {
	return c.RenderWithName(fn, props, getKeyName(fn, props))
}

func (c *Ctx) initView() {
	c.root = nil
	c.id.initPath()
	c.tick.init()
	c.zoneMap = make(map[string]*C)
	c.Cursor = nil

	c.ids = []string{}
	for _, cs := range c.components {
		cs.keyHandlers = make([]KeyHandler, 0)
		cs.mouseHandlers = make([]MouseHandler, 0)
		cs.messageHandlers = make([]MsgHandler, 0)
		cs.onFocused = nil
		cs.height = 0
		cs.width = 0

		cs.children = make([]*C, 0)
		cs.parent = nil
		cs.layout = Layout{}
	}
}

func (c *Ctx) initPhase() {
	c.ids = []string{}
	for _, cs := range c.components {

		cs.children = make([]*C, 0)
		cs.parent = nil
	}
}

// MouseZone creates a mouse zone for the given content.
// The ID of the zone is the components ID.
func (c *Ctx) MouseZone(content string) string {
	id := c.id.getID()
	instance, _ := c.getComponent(id)
	c.zoneMap[id] = instance
	markedContent := c.zone.Mark(id, content)
	return markedContent
}

// MouseZoneChild creates a mouse zone for child (sub part) of a component.
// The ID of the zone is the components ID + "###" + childID.
// MouseHandlers will receive the childID extracted from the ID mentioned above.
func (c *Ctx) MouseZoneChild(childID string, content string) string {
	id := c.id.getID()
	markedContent := c.zone.Mark(id+"###"+childID, content)
	return markedContent
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

func (c *Ctx) UpdateInMs(ms int) {
	if c.teaProgram == nil {
		panic("teaProgram is nil. Cannot update manually.")
	}

	if c.teaProgram != nil {
		go func() {
			<-time.After(time.Duration(ms) * time.Millisecond)
			c.teaProgram.Send(InvalidateMsg{})
		}()
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
