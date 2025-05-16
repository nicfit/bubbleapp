package app

import (
	"strings"

	"github.com/alexanderbh/bubbleapp/style"
	tea "github.com/charmbracelet/bubbletea/v2"
)

// C (Component) is a special type that can only be created via c.Render.
// This ensures all components are properly registered with the rendering system.
type C struct {
	content string
	id      string
}

// String allows the Component to be used in contexts where a string is expected
func (c C) String() string {
	return c.content
}

// Props is the generic type for component properties
type Props any

// FC (Function Component) returns a Component.
// This ensures that component functions MUST be wrapped with c.Render
// to create proper Components that can be used in layouts.
//
// Example:
//
//	func MyComponent(c *Ctx, props Props) string {  // Returns string
//	  return "Hello world"
//	}
//
//	// Usage always requires c.Render to create a Component:
//	result := c.Render(MyComponent, props)
type FC = func(ctx *Ctx, props Props) C

// Children is a function that returns a slice of pre-rendered Components.
// This ensures that all child elements must be created via c.Render while
// allowing for conditional logic when rendering children.
type Children = func(c *Ctx) []C
type Child = func(c *Ctx) C

type AppOptions struct {
	rootProps Props
}
type AppOption func(*AppOptions)

type app struct {
	root      FC
	rootProps Props
	ctx       *Ctx
}

func New(ctx *Ctx, root FC, options ...AppOption) *app {
	if ctx.zoneMap == nil {
		ctx.zoneMap = make(map[string]*instanceContext)
	}
	if ctx.Styles == nil {
		ctx.Styles = style.DefaultStyles()
	}

	opts := &AppOptions{}
	for _, opt := range options {
		opt(opts)
	}

	return &app{
		root:      root,
		rootProps: opts.rootProps,
		ctx:       ctx,
	}
}

func WithRootProps(props any) AppOption {
	return func(opts *AppOptions) {
		opts.rootProps = props
	}
}

func (a *app) SetTeaProgram(p *tea.Program) {
	a.ctx.teaProgram = p
}

func (a *app) Init() tea.Cmd {
	if a.ctx.teaProgram == nil {
		panic("teaProgram is nil. Set the tea.Program with app.SetTeaProgram(p).")
	}
	return nil
}

func (a *app) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case InvalidateMsg:
		return a, nil
	case tea.KeyMsg:
		focusedInstance, focusedInstanceExists := a.ctx.componentContext.get(a.ctx.UIState.Focused)
		if focusedInstanceExists {
			for _, handler := range focusedInstance.keyHandlers {
				if handler(msg) {
					return a, nil // Key was handled by the component's internal handler
				}
			}
		}

		// If no focused component handled the key, or there's no focused component,
		// handle global key bindings.
		switch msg.String() {
		case "ctrl+c":
			a.ctx.Quit()
			return a, nil
		case "tab":
			a.ctx.FocusNext()
			return a, nil
		case "shift+tab":
			a.ctx.FocusPrev()
			return a, nil
		}
		return a, nil
	case tea.WindowSizeMsg:
		a.ctx.layoutManager.width = msg.Width
		a.ctx.layoutManager.height = msg.Height
		return a, nil
	case tea.MouseMsg:
		idsInBounds := a.ctx.zone.IDsInBounds(msg)
		_, isMotionMsg := msg.(tea.MouseMotionMsg)
		if isMotionMsg {
			a.ctx.UIState.Hovered = ""
			a.ctx.UIState.HoveredChild = ""
		}
		for _, id := range idsInBounds {
			splitID := strings.Split(id, "###")
			id = splitID[0]
			childID := ""
			if len(splitID) > 1 {
				childID = splitID[1]
			}
			if isMotionMsg {
				a.ctx.UIState.Hovered = id // This will run for each meaning the last one will be the hovered one
				a.ctx.UIState.HoveredChild = childID
			}
			foundInstance, found := a.ctx.componentContext.get(id)
			if found {
				for _, handler := range foundInstance.mouseHandlers {
					if handler(msg, childID) {
						return a, nil // Mouse event was handled by the component's mouse handler
					}
				}
			}
		}
		return a, nil
	default:
		if a.ctx.UIState.Focused != "" {
			foundInstance, found := a.ctx.componentContext.get(a.ctx.UIState.Focused)
			if found {
				for _, handler := range foundInstance.messageHandlers {
					cmd := handler(msg)
					if cmd != nil {
						return a, cmd
					}
				}
			}
		}

	}

	return a, nil

}

func (a *app) View() string {
	// Get all component IDs before rendering (current state)
	prevIDs := a.ctx.componentContext.getAllIDs()

	a.ctx.initView()

	a.ctx.UIState.resetSizes()
	a.ctx.LayoutPhase = LayoutPhaseIntrincintWidth
	a.ctx.RenderWithName(func(c *Ctx, props Props) string { return a.root(a.ctx, props).String() }, a.rootProps, "Root")
	a.ctx.layoutManager.distributeWidth(a.ctx)
	// TODO: CONTENT WRAPPING PHASE HERE!!!! ************************************
	a.ctx.LayoutPhase = LayoutPhaseIntrincintHeight
	a.ctx.id.initIDCollections()
	a.ctx.id.initPath()
	a.ctx.RenderWithName(func(c *Ctx, props Props) string { return a.root(a.ctx, props).String() }, a.rootProps, "Root")
	a.ctx.layoutManager.distributeHeight(a.ctx)

	a.ctx.LayoutPhase = LayoutPhaseFinalRender
	a.ctx.invalidate = false
	a.ctx.id.initIDCollections()
	a.ctx.id.initPath()
	component := a.ctx.RenderWithName(func(c *Ctx, props Props) string { return a.root(a.ctx, props).String() }, a.rootProps, "Root")
	renderedView := a.ctx.zone.Scan(component.String())

	// Create or update the timer based on the current set of tick listeners
	a.ctx.tick.createTimer(a.ctx)

	// Get all component IDs after rendering (new state)
	currentIDs := a.ctx.id.ids // These are the IDs that were actually rendered

	// Determine removed IDs
	removedIDs := findRemovedIDs(prevIDs, currentIDs)

	// Cleanup effects for removed components
	a.ctx.componentContext.cleanupEffects(removedIDs)

	// Cleanup general state for removed components (from app/state.go logic)
	a.ctx.UIState.cleanup(currentIDs)

	return renderedView
}

// findRemovedIDs returns the IDs that are present in prevIDs but not in currentIDs.
func findRemovedIDs(prevIDs, currentIDs []string) []string {
	currentSet := make(map[string]struct{}, len(currentIDs))
	for _, id := range currentIDs {
		currentSet[id] = struct{}{}
	}
	var removed []string
	for _, id := range prevIDs {
		if _, found := currentSet[id]; !found {
			removed = append(removed, id)
		}
	}
	return removed
}
