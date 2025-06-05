package app

import (
	"strings"

	"github.com/alexanderbh/bubbleapp/style"
	tea "github.com/charmbracelet/bubbletea/v2"
)

// Props is the generic type for component properties
type Props any

// FCs is a function that returns a slice of pre-rendered Components.
// This ensures that all child elements must be created via c.Render while
// allowing for conditional logic when rendering children.
type FCs = func(c *Ctx) []*C
type FC = func(c *Ctx) *C

type AppOptions struct {
	Theme *style.AppTheme
}
type AppOption func(*AppOptions)

func WithTheme(theme *style.AppTheme) AppOption {
	return func(opts *AppOptions) {
		opts.Theme = theme
	}
}

type app struct {
	root FC
	ctx  *Ctx
}

func New(ctx *Ctx, root FC, options ...AppOption) *app {
	opts := &AppOptions{}
	for _, opt := range options {
		opt(opts)
	}

	if opts.Theme != nil {
		ctx.Theme = opts.Theme
	}

	return &app{
		root: root,
		ctx:  ctx,
	}
}

func (a *app) SetTeaProgram(p *tea.Program) {
	a.ctx.teaProgram = p
}

func (a *app) Init() tea.Cmd {
	if a.ctx.teaProgram == nil {
		panic("teaProgram is nil. Set the tea.Program with app.SetTeaProgram(p).")
	}
	var cmds []tea.Cmd
	if a.ctx.Theme.BackgroundColor != nil {
		cmds = append(cmds, tea.SetBackgroundColor(a.ctx.Theme.BackgroundColor))
	}
	if a.ctx.Theme.ForegroundColor != nil {
		cmds = append(cmds, tea.SetForegroundColor(a.ctx.Theme.ForegroundColor))
	}
	return tea.Batch(cmds...)
}

func (a *app) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// TODO: Add debug logging flag
	//log.Println("UPDATE", msg)
	switch msg := msg.(type) {
	case InvalidateMsg:
		return a, nil
	case tea.KeyMsg:
		focusedInstance, focusedInstanceExists := a.ctx.getComponent(a.ctx.UIState.Focused)
		if focusedInstanceExists {
			for _, handler := range focusedInstance.keyHandlers {
				if handler(msg) {
					return a, nil // Key was handled by the component's internal handler
				}
			}
		}

		// If the key was not handled by the focused component, check global key handlers.
		for _, handler := range a.ctx.getAllGlobalKeyHandlers() {
			if handler(msg) {
				return a, nil // Key was handled by a global key handler
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

		for i := range idsInBounds {
			id := idsInBounds[i]
			splitID := strings.Split(id, "###")
			id = splitID[0]
			childID := ""
			if len(splitID) > 1 {
				childID = splitID[1]
			}
			if isMotionMsg {
				// Longer ID means it is more specific here and thus hovered.
				// Not sure if that is valid always
				if a.ctx.UIState.Hovered == "" || len(id) > len(a.ctx.UIState.Hovered) {
					a.ctx.UIState.Hovered = id // This will run for each meaning the last one will be the hovered one
					a.ctx.UIState.HoveredChild = childID
				}
			}
			foundInstance, found := a.ctx.getComponent(id)
			if found {
				for _, handler := range foundInstance.mouseHandlers {
					if handler(msg, childID) {
						return a, nil // Mouse event was handled by the component's mouse handler
					}
				}
			}
		}
		// Nothing was clicked with the mouse so remove focus
		if releaseMsg, ok := msg.(tea.MouseReleaseMsg); ok {
			if releaseMsg.Button == tea.MouseLeft {
				a.ctx.UIState.Focused = ""
			}
		}
		return a, nil
	default:
		if a.ctx.UIState.Focused != "" {
			foundInstance, found := a.ctx.getComponent(a.ctx.UIState.Focused)
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

func (a *app) View() (string, *tea.Cursor) {
	// Get all component IDs before rendering (current state)
	prevIDs := a.ctx.ids

	a.ctx.initView()

	// Intrinsic width phase
	a.ctx.LayoutPhase = LayoutPhaseIntrincintWidth
	a.ctx.RenderWithName(func(c *Ctx, props Props) string {
		return a.root(c).String()
	}, nil, "Root")
	a.ctx.layoutManager.distributeWidth(a.ctx)
	// TODO: CONTENT WRAPPING PHASE HERE!!!! ************************************

	// Intrinsic height phase
	a.ctx.LayoutPhase = LayoutPhaseIntrincintHeight
	a.ctx.initPhase()
	a.ctx.id.initPath()
	a.ctx.RenderWithName(func(c *Ctx, props Props) string {
		return a.root(c).String()
	}, nil, "Root")
	a.ctx.layoutManager.distributeHeight(a.ctx)

	// Absolute positioning phase
	a.ctx.LayoutPhase = LayoutPhaseAbsolutePositions
	a.ctx.invalidate = false
	a.ctx.initPhase()
	a.ctx.id.initPath()
	a.ctx.RenderWithName(func(c *Ctx, props Props) string {
		return a.root(c).String()
	}, nil, "Root")
	a.ctx.layoutManager.calculatePositions(a.ctx)

	// Final render phase
	a.ctx.LayoutPhase = LayoutPhaseFinalRender
	a.ctx.invalidate = false
	a.ctx.initPhase()
	a.ctx.id.initPath()

	rootComponent := a.ctx.RenderWithName(func(c *Ctx, props Props) string {
		return a.root(c).String()
	}, nil, "Root")
	renderedView := a.ctx.zone.Scan(rootComponent.String())

	// Create or update the timer based on the current set of tick listeners
	a.ctx.tick.createTimer(a.ctx)

	// Get all component IDs after rendering (new state)
	currentIDs := a.ctx.ids

	// Determine removed IDs
	removedIDs := findRemovedIDs(prevIDs, currentIDs)

	// Cleanup effects for removed components
	a.ctx.cleanupEffects(removedIDs)

	for _, removedID := range removedIDs {
		delete(a.ctx.components, removedID)
	}

	return renderedView, a.ctx.Cursor
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
