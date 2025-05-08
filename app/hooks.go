package app

// Returns the component ID
func UseID(c *FCContext) string {
	return c.id.getID()
}

// Registers component as focusable and returns the focus state
func UseFocus(c *FCContext) bool {
	instance, instanceExists := c.componentContext.get(c.id.getID())
	if instanceExists {
		instance.focusable = true
	}
	return c.UIState.Focused == c.id.getID()
}

func UseChildren(c *FCContext, children Children) []string {
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
func UseState[T any](c *FCContext, initialValue T) (T, func(newValue T)) {
	instanceID := c.id.getID()
	// FCContext.Render guarantees that the instance exists by calling componentContext.set
	instance, _ := c.componentContext.get(instanceID)

	hookIndex := c.UseStateCounter
	c.UseStateCounter++

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
		// Trigger a re-render. Sending a nil message is a common way to do this in BubbleTea.
		// The main app's Update method should handle `nil` msg to just re-render.

	}

	return currentValue, setter
}
