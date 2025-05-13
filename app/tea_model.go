package app

import (
	"time"

	"github.com/alexanderbh/bubbleapp/style"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type Props any
type rootProps struct {
	Layout Layout
}
type FC = func(ctx *Ctx, props Props) string
type Children func(ctx *Ctx)

// _____________________

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
	ctx     *Ctx
	tickFPS time.Duration
}

func New(ctx *Ctx, root FC, options ...AppOption) *app {
	if options == nil {
		options = []AppOption{}
	}
	if ctx.zoneMap == nil {
		ctx.zoneMap = make(map[string]*instanceContext)
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

	switch msg := msg.(type) {
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
		for _, id := range idsInBounds {
			foundInstance, found := a.ctx.componentContext.get(id)
			if found {
				for _, handler := range foundInstance.mouseHandlers {
					if handler(msg) {
						return a, nil // Mouse event was handled by the component's mouse handler
					}
				}
			}
		}
		return a, nil

	}

	return a, nil

}

func (a *app) View() string {
	// Get all component IDs before rendering (current state)
	prevIDs := a.ctx.componentContext.getAllIDs()

	a.ctx.initView()

	defaultRootProps := rootProps{
		Layout: Layout{
			GrowX: true,
			GrowY: true,
		},
	}

	a.ctx.UIState.resetSizes()
	a.ctx.LayoutPhase = LayoutPhaseIntrincintWidth
	a.ctx.Render(a.root, defaultRootProps)
	a.ctx.layoutManager.distributeWidth(a.ctx)
	// TODO: CONTENT WRAPPING PHASE HERE!!!! ************************************
	a.ctx.LayoutPhase = LayoutPhaseIntrincintHeight
	a.ctx.id.initIDCollections()
	a.ctx.id.initPath()
	a.ctx.Render(a.root, defaultRootProps)
	a.ctx.layoutManager.distributeHeight(a.ctx)

	a.ctx.LayoutPhase = LayoutPhaseFinalRender
	a.ctx.invalidate = false
	a.ctx.id.initIDCollections()
	a.ctx.id.initPath()
	renderedView := a.ctx.zone.Scan(a.ctx.Render(a.root, defaultRootProps))

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
