package app

import (
	"fmt" // Added import for error formatting
	"reflect"
	"time"

	tea "github.com/charmbracelet/bubbletea/v2"
)

// ChildMissingRenderError is the error thrown when a component
// in a layout isn't properly registered with c.Render
type ChildMissingRenderError struct {
	ComponentPath string
}

func (e ChildMissingRenderError) Error() string {
	return fmt.Sprintf("Component at path '%s' was not properly registered with c.Render - always use c.Render for components in layouts", e.ComponentPath)
}

// Returns the component ID
func UseID(c *Ctx) string {
	return c.id.getID()
}

// Registers component as focusable and returns the focus state
func UseIsFocused(c *Ctx) bool {
	instance, instanceExists := c.componentContext.get(c.id.getID())
	if instanceExists {
		instance.focusable = true
	}
	return c.UIState.Focused == c.id.getID()
}

func UseOnFocused(c *Ctx, onFocused func(isReverse bool)) {
	if c.LayoutPhase != LayoutPhaseFinalRender {
		return
	}
	instanceID := c.id.getID()
	instance, exists := c.componentContext.get(instanceID)
	if !exists {
		panic("UseOnFocused: component instance not found")
	}
	instance.focusable = true
	instance.onFocused = onFocused
}

// Returns the ID of the component that is hovered.
// If the component is hovered, it returns true and any potential child ID.
func UseIsHovered(c *Ctx) (bool, string) {
	hoveredID := c.UIState.Hovered == c.id.getID()
	if hoveredID {
		return true, c.UIState.HoveredChild
	}
	return false, ""
}

func UseSize(c *Ctx) (int, int) {
	if c.LayoutPhase == LayoutPhaseIntrincintWidth {
		return c.layoutManager.width, c.layoutManager.height
	}
	id := c.id.getID()
	if c.LayoutPhase == LayoutPhaseIntrincintHeight {
		return c.UIState.GetWidth(id), c.layoutManager.height
	}
	return c.UIState.GetWidth(id), c.UIState.GetHeight(id)
}

// UseFCs executes the children function to get pre-rendered Components
// and returns their string contents for layout component consumption
func UseFCs(c *Ctx, fcs FCs) []string {
	if fcs == nil {
		return []string{}
	}
	// Execute the children function to get the components
	components := fcs(c)

	// Allocate space for outputs
	outputs := make([]string, 0, len(components))

	// Extract the content from each pre-rendered Component
	for _, component := range components {
		outputs = append(outputs, component.String())
	}

	if len(outputs) == 0 {
		return []string{""}
	}

	return outputs
}

// UseState provides stateful value and a function to update it.
// It's analogous to React's useState hook.
// IMPORTANT: Hooks must be called in the same order on every render,
// and they must not be called conditionally.
func UseState[T any](c *Ctx, initialValue T) (T, func(valueOrUpdater interface{})) {
	instanceID := c.id.getID()
	instance, _ := c.componentContext.get(instanceID)

	hookIndex := c.useStateCounter[instanceID]
	c.useStateCounter[instanceID]++

	if hookIndex >= len(instance.states) {
		instance.states = append(instance.states, initialValue)
	}

	currentValue := instance.states[hookIndex].(T)

	setter := func(valueOrUpdater interface{}) {
		var newValue T
		previousValueFromState := instance.states[hookIndex].(T)

		switch updater := valueOrUpdater.(type) {
		case func(prevValue T) T:
			newValue = updater(previousValueFromState)
		case T:
			newValue = updater
		default:
			typeName := reflect.TypeOf(initialValue).String()
			panic(fmt.Sprintf("UseState setter: unexpected type %T passed. Expected %s or func(prevValue %s) %s",
				valueOrUpdater, typeName, typeName, typeName))
		}
		instance.states[hookIndex] = newValue
		c.Update()
	}

	return currentValue, setter
}

// UseTick schedules a function to be called at a specified interval.
// The callback will be invoked in a separate goroutine managed by the tick system.
func UseTick(c *Ctx, interval time.Duration, callback func()) {
	if c.LayoutPhase != LayoutPhaseFinalRender {
		return
	}
	instanceID := c.id.getID()
	c.tick.RegisterTickListener(interval, instanceID, callback)
	UseEffectWithCleanup(c, func() func() {
		// Return the cleanup function.
		return func() {
			if c.tick != nil {
				c.tick.UnregisterTickListener(instanceID)
			}
		}
	}, []any{})
}

type effectRecord struct {
	cleanupFn   func() // The cleanup function returned by the effect.
	deps        []any  // Dependencies for the effect.
	hasExecuted bool   // Tracks if the effect has executed at least once.
}

var RunOnceDeps = []any{}

// UseEffect is the same as UseEffectWithCleanup but without a cleanup function.
func UseEffect(c *Ctx, effect func(), deps []any) {
	UseEffectWithCleanup(c, func() func() {
		effect()
		return nil
	}, deps)
}

// UseEffect schedules a function to run after render, and optionally clean up.
// Dependencies (deps) are checked to see if the effect should re-run.
// If deps is nil, the effect runs after every render.
// If deps is an empty slice, it runs only once after the initial render and on unmount.
func UseEffectWithCleanup(c *Ctx, effect func() func(), deps []any) {
	if c.LayoutPhase != LayoutPhaseFinalRender {
		return
	}
	instanceID := c.id.getID()
	instance, _ := c.componentContext.get(instanceID)

	hookIndex := c.useEffectCounter[instanceID]
	c.useEffectCounter[instanceID]++

	if hookIndex >= len(instance.effects) {
		instance.effects = append(instance.effects, effectRecord{})
	}

	record := &instance.effects[hookIndex]

	depsChanged := true // Assume changed for nil deps (run every time) or first run
	if record.hasExecuted && deps != nil {
		if len(deps) == 0 && len(record.deps) == 0 { // Both empty (e.g. RunOnceDeps), no change
			depsChanged = false
		} else if len(deps) == len(record.deps) { // Same length, check elements
			depsChanged = false // Assume no change until a difference is found
			for i, currentDep := range deps {
				oldDep := record.deps[i]

				// 1. Handle nil cases for individual dependencies
				if currentDep == nil && oldDep == nil {
					continue // Both nil, considered same for this element
				}
				if currentDep == nil || oldDep == nil {
					depsChanged = true // One is nil, the other isn't, so different
					break
				}

				// 2. Use reflection for actual comparison
				valCurrent := reflect.ValueOf(currentDep)
				valOld := reflect.ValueOf(oldDep)

				// 3. If types are different, dependencies have changed
				if valCurrent.Type() != valOld.Type() {
					depsChanged = true
					break
				}

				// 4. Compare values based on comparability
				if valCurrent.Type().Comparable() {
					// For comparable types, direct value comparison
					if currentDep != oldDep {
						depsChanged = true
						break
					}
				} else {
					// For non-comparable types (e.g., slice, map, func, or struct with non-comparable field)
					// Compare by pointer for types where it's meaningful (slice, map, func, chan, ptr, unsafeptr)
					kind := valCurrent.Kind()
					if kind == reflect.Chan || kind == reflect.Func || kind == reflect.Map || kind == reflect.Ptr || kind == reflect.Slice || kind == reflect.UnsafePointer {
						if valCurrent.Pointer() != valOld.Pointer() {
							depsChanged = true
							break
						}
					} else {
						// For other non-comparable types (e.g., a struct passed by value that contains a slice).
						// Treat as changed, as new instances won't be pointer-equal.
						// This mimics React's behavior for new object/array literals in deps.
						depsChanged = true
						break
					}
				}
			}
		} else {
			// Lengths are different, depsChanged remains true (initial assumption)
			depsChanged = true
		}
	}

	if depsChanged {
		// If a cleanup function exists from a previous run, execute it
		if record.cleanupFn != nil {
			record.cleanupFn()
		}

		// Execute the effect and store any returned cleanup function
		record.cleanupFn = effect()

		// Store a snapshot of the dependencies
		record.deps = deps
		record.hasExecuted = true
		c.Update()
	}
}

// UseKeyHandler registers a function to handle key presses within a component.
// This handler is only called if the component is focused.
// The handler function should return true if it handled the key, false otherwise.
func UseKeyHandler(c *Ctx, handler KeyHandler) {
	if c.LayoutPhase != LayoutPhaseFinalRender {
		return
	}
	instanceID := c.id.getID()
	instance, exists := c.componentContext.get(instanceID)
	if !exists {
		panic("UseKeyHandler: component instance not found")
	}
	instance.focusable = true
	instance.keyHandlers = append(instance.keyHandlers, handler)
}

// UseMouseHandler registers a function to handle mouse events within a component.
func UseMouseHandler(c *Ctx, handler MouseHandler) {
	if c.LayoutPhase != LayoutPhaseFinalRender {
		return
	}
	instanceID := c.id.getID()
	instance, exists := c.componentContext.get(instanceID)
	if !exists {
		panic("UseMouseHandler: component instance not found")
	}
	instance.focusable = true
	instance.mouseHandlers = append(instance.mouseHandlers, handler)
}

// UseAction registers a function to be called when the component is clicked
// with left mouse button or Enter is pressed while the component is focused.
func UseAction(c *Ctx, handler func(childID string)) {
	if c.LayoutPhase != LayoutPhaseFinalRender {
		return
	}
	instanceID := c.id.getID()
	instance, exists := c.componentContext.get(instanceID)
	if !exists {
		panic("UseAction: component instance not found")
	}
	instance.focusable = true
	instance.mouseHandlers = append(instance.mouseHandlers, func(msg tea.MouseMsg, childID string) bool {
		if releaseMsg, ok := msg.(tea.MouseReleaseMsg); ok && releaseMsg.Mouse().Button == tea.MouseLeft {
			handler(childID)
			return true
		}
		return false
	})
	instance.keyHandlers = append(instance.keyHandlers, func(keyMsg tea.KeyMsg) bool {
		if keyMsg.String() == "enter" {
			handler("")
			return true
		}
		return false
	})
}

func UseMsgHandler(c *Ctx, handler MsgHandler) {
	if c.LayoutPhase != LayoutPhaseFinalRender {
		return
	}
	instanceID := c.id.getID()
	instance, exists := c.componentContext.get(instanceID)
	if !exists {
		panic("UseMsgHandler: component instance not found")
	}
	instance.messageHandlers = append(instance.messageHandlers, handler)
}
