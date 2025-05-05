package app

func (ctx *Context[T]) getAllFocusable(root Fc[T]) []Fc[T] {
	focusableItems := []Fc[T]{}

	var traverse func(Fc[T])
	traverse = func(node Fc[T]) {
		if node == nil {
			return
		}

		if node.Base().Opts.Focusable {
			focusableItems = append(focusableItems, node)
		}
		for _, child := range node.Children(ctx) {
			traverse(child)
		}

	}

	traverse(root)

	return focusableItems
}

func (ctx *Context[T]) FocusFirstCmd(root Fc[T]) {

	focusableItems := ctx.getAllFocusable(root)
	if len(focusableItems) == 0 {
		return
	}
	ctx.UIState.Focused = focusableItems[0].Base().ID
}

func (ctx *Context[T]) FocusNextCmd(root Fc[T]) {

	focusableItems := ctx.getAllFocusable(root)
	if len(focusableItems) == 0 {
		ctx.UIState.Focused = ""
		return
	}
	if len(focusableItems) == 1 {
		if ctx.UIState.Focused == focusableItems[0].Base().ID {
			ctx.UIState.Focused = ""
			return
		}
		ctx.UIState.Focused = focusableItems[0].Base().ID
		return
	}

	currentIndex := -1
	if ctx.UIState.Focused != "" {
		for i, item := range focusableItems {
			if item.Base().ID == ctx.UIState.Focused {
				currentIndex = i
				break
			}
		}
	}

	if currentIndex == -1 {
		ctx.UIState.Focused = focusableItems[0].Base().ID
		return
	}

	nextIndex := (currentIndex + 1) % len(focusableItems)

	ctx.UIState.Focused = focusableItems[nextIndex].Base().ID
}

func (ctx *Context[T]) FocusPrevCmd(root Fc[T]) {

	focusableItems := ctx.getAllFocusable(root)
	if len(focusableItems) == 0 {
		ctx.UIState.Focused = ""
		return
	}
	if len(focusableItems) == 1 {
		ctx.UIState.Focused = focusableItems[0].Base().ID
		return
	}

	currentIndex := -1
	if ctx.UIState.Focused != "" {
		for i, item := range focusableItems {
			if item.Base().ID == ctx.UIState.Focused {
				currentIndex = i
				break
			}
		}
	}

	if currentIndex == -1 {
		ctx.UIState.Focused = focusableItems[len(focusableItems)-1].Base().ID
		return
	}

	prevIndex := (currentIndex - 1 + len(focusableItems)) % len(focusableItems)

	ctx.UIState.Focused = focusableItems[prevIndex].Base().ID
}
