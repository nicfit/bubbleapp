package app

// _getFocusableComponentIDs retrieves all focusable component IDs.
func _getFocusableComponentIDs(c *FCContext) []string {
	var focusableIDs []string
	if c.id == nil || c.componentContext == nil {
		// Should not happen in a healthy context
		return focusableIDs
	}

	for _, id := range c.id.ids {
		componentInfo := c.componentContext.get(id)
		if componentInfo != nil && componentInfo.focusable {
			focusableIDs = append(focusableIDs, id)
		}
	}
	return focusableIDs
}

// _findCurrentFocusIndex finds the index of the currentFocusID in the focusableIDs list.
// Returns -1 if not found.
func _findCurrentFocusIndex(focusableIDs []string, currentFocusID string) int {
	for i, id := range focusableIDs {
		if id == currentFocusID {
			return i
		}
	}
	return -1
}

func (c *FCContext) FocusNext() string {
	focusableIDs := _getFocusableComponentIDs(c)
	if len(focusableIDs) == 0 {
		c.UIState.Focused = ""
		return "" // No items to focus
	}

	currentFocusID := c.UIState.Focused
	currentIndex := _findCurrentFocusIndex(focusableIDs, currentFocusID)

	nextIndex := 0
	if currentIndex != -1 {
		// Current focused item is in the list, move to the next
		nextIndex = (currentIndex + 1) % len(focusableIDs)
	} else {
		// Current focused item is not in the list (e.g. initially empty, or item disappeared)
		// or no item was focused; focus the first available item.
		nextIndex = 0
	}

	nextIDToFocus := focusableIDs[nextIndex]
	c.UIState.Focused = nextIDToFocus
	return nextIDToFocus
}

func (c *FCContext) FocusPrev() string {
	focusableIDs := _getFocusableComponentIDs(c)
	if len(focusableIDs) == 0 {
		c.UIState.Focused = ""
		return "" // No items to focus
	}

	currentFocusID := c.UIState.Focused
	currentIndex := _findCurrentFocusIndex(focusableIDs, currentFocusID)

	prevIndex := 0
	if currentIndex != -1 {
		// Current focused item is in the list, move to the previous
		prevIndex = (currentIndex - 1 + len(focusableIDs)) % len(focusableIDs)
	} else {
		// Current focused item is not in the list (e.g. initially empty, or item disappeared)
		// or no item was focused; focus the last available item.
		prevIndex = len(focusableIDs) - 1
	}

	prevIDToFocus := focusableIDs[prevIndex]
	c.UIState.Focused = prevIDToFocus
	return prevIDToFocus
}
