package app

import (
	"time"

	"github.com/alexanderbh/bubbleapp/style"
	zone "github.com/alexanderbh/bubblezone/v2"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type FCContext struct {
	UIState    *StateStore
	Zone       *zone.Manager
	ZoneMap    map[string]*fcInstance
	teaProgram *tea.Program
	Styles     *style.Styles
	id         *fcIDContext
	Tick       *tickState[any]

	collectorStack []*outputCollector
	isCollecting   bool

	componentContext *fcInstanceContext

	LayoutPhase bool
	Width       int
	Height      int
}

type outputCollector struct {
	outputs []string
}

func NewFCContext() *FCContext {
	return &FCContext{
		UIState:          NewStateStore(),
		Zone:             zone.New(),
		ZoneMap:          make(map[string]*fcInstance),
		Styles:           style.DefaultStyles(),
		id:               newFCIDContext(),
		Tick:             &tickState[any]{},
		collectorStack:   []*outputCollector{},
		componentContext: newFCInstanceContext(),
	}
}

func (c *FCContext) Render(fc FC, props Props) string {
	id := c.id.push(getFuncName(fc))
	defer c.id.pop()

	c.id.ids = append(c.id.ids, id)
	c.componentContext.set(id, fc, props)

	output := fc(c, props)

	// If there is an active output collector, append the output to it
	if len(c.collectorStack) > 0 {
		currentCollector := c.collectorStack[len(c.collectorStack)-1]
		currentCollector.outputs = append(currentCollector.outputs, output)
	}

	return output
}
func (ctx *FCContext) UseID() string {
	return ctx.id.getID()
}

// Registers component as focusable and returns the focus state
func (ctx *FCContext) UseFocus() bool {
	ctx.componentContext.get(ctx.id.getID()).focusable = true
	return ctx.UIState.Focused == ctx.id.getID()
}

func (ctx *FCContext) UseChildren(children Children) []string {
	newCollector := &outputCollector{}
	ctx.collectorStack = append(ctx.collectorStack, newCollector)

	if children != nil {
		children(ctx)
	}

	ctx.collectorStack = ctx.collectorStack[:len(ctx.collectorStack)-1]

	childOutputs := newCollector.outputs
	if len(childOutputs) == 0 {
		return []string{""}
	}

	return childOutputs
}

// _____________________

// Quit signals the application to stop, ensuring cleanup like stopping active timers.
func (ctx *FCContext) Quit() {
	if ctx.Tick != nil {
		ctx.Tick.StopActiveTimer()
	}
	go ctx.teaProgram.Quit()
}

type Props any
type FC = func(ctx *FCContext, props Props) string
type Children func(ctx *FCContext)

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

type app struct {
	root    FC
	ctx     *FCContext
	tickFPS time.Duration
}

func New(ctx *FCContext, root FC, options ...AppOption) *app {
	if options == nil {
		options = []AppOption{}
	}
	if ctx.ZoneMap == nil {
		ctx.ZoneMap = make(map[string]*fcInstance)
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

	return &app{
		root:    root,
		ctx:     ctx,
		tickFPS: opts.TickFPS,
	}
}

func (a *app) SetTeaProgram(p *tea.Program) {
	a.ctx.teaProgram = p
}

func (a *app) Init() tea.Cmd {
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

func (a *app) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		focusedInstance := a.ctx.componentContext.get(a.ctx.UIState.Focused)
		if focusedInstance != nil {
			switch msg.String() {
			case "enter":
				if a.dispatchToHandler(focusedInstance, semanticActionPrimary, "OnKeyPress", KeyEvent{
					Key: msg.Key(),
				}) {
					return a, nil
				}
			}
		}

		//handledKeyMsg := a.propagatedFocused(msg)
		//	if !handledKeyMsg {
		// These are global keys. Is this what we want?
		switch msg.String() {
		case "ctrl+c":
			a.ctx.Quit()
			return a, nil // Use the new Quit method
		case "tab":
			a.ctx.FocusNext()
			return a, nil
		case "shift+tab":
			a.ctx.FocusPrev()
			return a, nil
		}
		//}
		return a, tea.Batch(cmds...)

	case tea.WindowSizeMsg:
		a.ctx.Width = msg.Width
		a.ctx.Height = msg.Height
		return a, tea.Batch(cmds...)
	case tea.MouseMsg:
		a.ctx.Zone.AnyInBounds(a, msg)
		return a, tea.Batch(cmds...)
	case zone.MsgZoneInBounds:
		foundInstance := a.ctx.componentContext.get(a.ctx.Zone.GetReverse(msg.Zone.Id))
		if foundInstance != nil {
			switch msg.Event.(type) {
			case tea.MouseClickMsg:
				a.dispatchToHandler(foundInstance, semanticActionPrimary, "OnClick", MouseEvent{
					X:      msg.Event.Mouse().X,
					Y:      msg.Event.Mouse().Y,
					Button: msg.Event.Mouse().Button,
					Mod:    msg.Event.Mouse().Mod,
				})
			}
		}
		return a, nil
		// case TickMsg:
		// 	now := msg.OccurredAt
		// 	for _, listener := range *a.ctx.Tick.tickListeners {
		// 		l := a.ctx.id.getNode(listener.id)
		// 		if l != nil {
		// 			lastTick, ok := a.ctx.Tick.lastTickTimes[listener.id]
		// 			if !ok || now.Sub(lastTick) >= listener.interval {
		// 				l.Update(a.ctx, msg)
		// 				a.ctx.Tick.lastTickTimes[listener.id] = now
		// 			}
		// 		}
		// 	}
		// 	if a.tickFPS > 0 {
		// 		cmds = append(cmds, tickCommand(a.tickFPS))
		// 	}
		// 	return a, tea.Batch(cmds...)
	}

	//a.propagatedFocused(msg)

	return a, tea.Batch(cmds...)

}

// func (a *app) propagatedFocused(msg tea.Msg) bool {
// 	if a.ctx.UIState.Focused != "" {
// 		focused := a.ctx.id.getNode(a.ctx.UIState.Focused)
// 		if focused != nil {
// 			handled := focused.Update(a.ctx, msg)
// 			return handled
// 		}
// 	}
// 	return false
// }

func (a *app) View() string {
	a.ctx.id.initPath()
	a.ctx.Tick.init()

	a.ctx.LayoutPhase = true
	a.ctx.Render(a.root, nil)
	a.ctx.LayoutPhase = false
	a.ctx.id.initPath()

	//a.ctx.Tick.createTimer(a.ctx)

	//a.Layout()
	return a.ctx.Zone.Scan(a.ctx.Render(a.root, nil))
}
