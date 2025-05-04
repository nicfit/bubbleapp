package app

import tea "github.com/charmbracelet/bubbletea/v2"

type FocusComponentMsg struct {
	TargetID string
}
type BlurAllMsg struct{}

func sendFocusMsg(targetID string) tea.Cmd {
	return func() tea.Msg {
		return FocusComponentMsg{TargetID: targetID}
	}
}

func (fm *Context[T]) getAllFocusable(root Fc[T]) []Fc[T] {
	focusableItems := []Fc[T]{}

	var traverse func(Fc[T])
	traverse = func(node Fc[T]) {
		if node == nil {
			return
		}

		if node.Base().Opts.Focusable {
			focusableItems = append(focusableItems, node)
		}
		for _, child := range node.Children(fm) {
			traverse(child)
		}

	}

	traverse(root)

	return focusableItems
}

func (fm *Context[T]) FocusFirstCmd(root Fc[T]) {

	focusableItems := fm.getAllFocusable(root)
	if len(focusableItems) == 0 {
		return
	}
	fm.Focused = focusableItems[0]
}

func (fm *Context[T]) FocusNextCmd(root Fc[T]) {

	focusableItems := fm.getAllFocusable(root)
	if len(focusableItems) == 0 {
		fm.Focused = nil
		return
	}
	if len(focusableItems) == 1 {
		if fm.Focused == focusableItems[0] {
			fm.Focused = nil
			return
		}
		fm.Focused = focusableItems[0]
		return
	}

	currentIndex := -1
	if fm.Focused != nil {
		for i, item := range focusableItems {
			if item == fm.Focused {
				currentIndex = i
				break
			}
		}
	}

	if currentIndex == -1 {
		fm.Focused = focusableItems[0]
		return
	}

	nextIndex := (currentIndex + 1) % len(focusableItems)

	fm.Focused = focusableItems[nextIndex]
}

func (fm *Context[T]) FocusPrevCmd(root Fc[T]) {

	focusableItems := fm.getAllFocusable(root)
	if len(focusableItems) == 0 {
		fm.Focused = nil
		return
	}
	if len(focusableItems) == 1 {
		fm.Focused = focusableItems[0]
		return
	}

	currentIndex := -1
	if fm.Focused != nil {
		for i, item := range focusableItems {
			if item == fm.Focused {
				currentIndex = i
				break
			}
		}
	}

	if currentIndex == -1 {
		fm.Focused = focusableItems[len(focusableItems)-1]
		return
	}

	prevIndex := (currentIndex - 1 + len(focusableItems)) % len(focusableItems)

	fm.Focused = focusableItems[prevIndex]
}
