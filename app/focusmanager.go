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

func (fm *Context[T]) getAllFocusable(root *Base[T]) []*Base[T] {
	focusableItems := []*Base[T]{}

	var traverse func(*Base[T])
	traverse = func(node *Base[T]) {
		if node == nil {
			return
		}

		if node.Opts.Focusable {
			focusableItems = append(focusableItems, node)
		}

		if node.Children != nil {
			for _, child := range node.Children {
				traverse(child)
			}
		}
	}

	traverse(root)

	return focusableItems
}

func (fm *Context[T]) FocusFirstCmd(root *Base[T]) tea.Cmd {

	focusableItems := fm.getAllFocusable(root)
	if len(focusableItems) == 0 {
		return sendFocusMsg("")
	}
	firstID := focusableItems[0].ID
	fm.FocusedID = firstID
	return sendFocusMsg(firstID)
}

func (fm *Context[T]) FocusNextCmd(root *Base[T]) tea.Cmd {

	focusableItems := fm.getAllFocusable(root)
	if len(focusableItems) == 0 {
		return sendFocusMsg("")
	}
	if len(focusableItems) == 1 {
		return sendFocusMsg(focusableItems[0].ID)
	}

	currentIndex := -1
	if fm.FocusedID != "" {
		for i, item := range focusableItems {
			if item.ID == fm.FocusedID {
				currentIndex = i
				break
			}
		}
	}

	if currentIndex == -1 {
		return sendFocusMsg(focusableItems[0].ID)
	}

	nextIndex := (currentIndex + 1) % len(focusableItems)
	nextID := focusableItems[nextIndex].ID

	fm.FocusedID = nextID
	return sendFocusMsg(nextID)
}

func (fm *Context[T]) FocusPrevCmd(root *Base[T]) tea.Cmd {

	focusableItems := fm.getAllFocusable(root)
	if len(focusableItems) == 0 {
		return sendFocusMsg("")
	}
	if len(focusableItems) == 1 {
		return sendFocusMsg(focusableItems[0].ID)
	}

	currentIndex := -1
	if fm.FocusedID != "" {
		for i, item := range focusableItems {
			if item.ID == fm.FocusedID {
				currentIndex = i
				break
			}
		}
	}

	if currentIndex == -1 {
		return sendFocusMsg(focusableItems[len(focusableItems)-1].ID)
	}

	prevIndex := (currentIndex - 1 + len(focusableItems)) % len(focusableItems)
	prevID := focusableItems[prevIndex].ID

	fm.FocusedID = prevID
	return sendFocusMsg(prevID)
}
