package app

// Returns the component ID
func UseID(c *Ctx) string {
	return c.id.getID()
}

// Registers component as focusable and returns the focus state
func UseFocus(c *Ctx) bool {
	instance, instanceExists := c.componentContext.get(c.id.getID())
	if instanceExists {
		instance.focusable = true
	}
	return c.UIState.Focused == c.id.getID()
}

func UseChildren(c *Ctx, children Children) []string {
	newCollector := &outputCollector{}
	c.collectorStack = append(c.collectorStack, newCollector)

	if children != nil {
		children(c)
	}

	c.collectorStack = c.collectorStack[:len(c.collectorStack)-1]

	childOutputs := newCollector.outputs
	if len(childOutputs) == 0 {
		return []string{""}
	}

	return childOutputs
}

// UseState provides stateful value and a function to update it.
// It's analogous to React's useState hook.
// IMPORTANT: Hooks must be called in the same order on every render,
// and they must not be called conditionally.
func UseState[T any](c *Ctx, initialValue T) (T, func(newValue T)) {
	instanceID := c.id.getID()
	// FCContext.Render guarantees that the instance exists by calling componentContext.set
	instance, _ := c.componentContext.get(instanceID)

	hookIndex := c.useStateCounter[instanceID]
	c.useStateCounter[instanceID]++

	if hookIndex >= len(instance.States) {
		// This is the first render for this hook in this component instance,
		// or the States slice needs to grow.
		instance.States = append(instance.States, initialValue)
	}
	// If hookIndex < len(instance.States), the state already exists from a previous render.

	// Type assertion. This could panic if the type T changes for a given hookIndex
	// between renders, which would be a misuse of the hook.
	currentValue := instance.States[hookIndex].(T)

	setter := func(newValue T) {
		instance.States[hookIndex] = newValue
		if c.teaProgram != nil {
			go c.teaProgram.Send(InvalidateMsg{})
		}
	}

	return currentValue, setter
}

type effectRecord struct {
	cleanupFn   func() // The cleanup function returned by the effect.
	deps        []any  // Dependencies for the effect.
	hasExecuted bool   // Tracks if the effect has executed at least once.
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

	if hookIndex >= len(instance.Effects) {
		instance.Effects = append(instance.Effects, effectRecord{})
	}

	record := &instance.Effects[hookIndex]

	// Check if dependencies have changed
	depsChanged := true                    // Assume changed for nil deps (run every time) or first run
	if record.hasExecuted && deps != nil { // Only check deps if not first run and deps are provided
		if len(deps) == 0 && len(record.deps) == 0 { // Both empty, no change
			depsChanged = false
		} else if len(deps) == len(record.deps) {
			depsChanged = false // Assume no change until a difference is found
			for i, d := range deps {
				if d != record.deps[i] {
					depsChanged = true
					break
				}
			}
		}
		// If lengths are different, depsChanged remains true
	}

	if depsChanged {
		// If a cleanup function exists from a previous run, execute it
		if record.cleanupFn != nil {
			record.cleanupFn()
		}

		// Execute the effect and store any returned cleanup function
		record.cleanupFn = effect()
		record.deps = deps
		record.hasExecuted = true
		if c.teaProgram != nil {
			go c.teaProgram.Send(InvalidateMsg{})
		}
	}
}

// UseKeyHandler registers a function to handle key presses internally within a component.
// This handler is only called if the component is focused and the key event is not
// handled by a more specific semantic handler (like OnKeyPress for Enter).
// The handler function should return true if it handled the key, false otherwise.
func UseKeyHandler(c *Ctx, handler InternalKeyHandler) {
	instanceID := c.id.getID()
	instance, exists := c.componentContext.get(instanceID)
	if !exists {
		// This should ideally not happen if hooks are called correctly within a component's lifecycle
		return
	}
	instance.internalKeyHandler = handler
}

func UseEffect(c *Ctx, effect func(), deps []any) {
	UseEffectWithCleanup(c, func() func() {
		effect()
		return nil
	}, deps)
}
