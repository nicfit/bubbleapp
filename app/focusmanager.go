package app

import tea "github.com/charmbracelet/bubbletea/v2"

type FocusManager struct {
	FocusedID string
}

type FocusComponentMsg struct {
	TargetID string
}
type BlurAllMsg struct{}

func NewFocusManager() *FocusManager {
	fm := &FocusManager{}

	return fm
}

func sendFocusMsg(targetID string) tea.Cmd {
	return func() tea.Msg {
		return FocusComponentMsg{TargetID: targetID}
	}
}

func (fm *FocusManager) getAllFocusable(root UIModel) []UIModel {
	focusableItems := []UIModel{}

	var traverse func(UIModel)
	traverse = func(node UIModel) {
		if node == nil {
			return
		}

		if node.Base().Opts.Focusable {
			focusableItems = append(focusableItems, node)
		}

		if node.Base().Children != nil {
			for _, child := range node.Base().Children {
				traverse(child)
			}
		}
	}

	traverse(root)

	return focusableItems
}

func (fm *FocusManager) FocusFirstCmd(root UIModel) tea.Cmd {

	focusableItems := fm.getAllFocusable(root)
	if len(focusableItems) == 0 {
		return sendFocusMsg("")
	}
	firstID := focusableItems[0].Base().ID
	fm.FocusedID = firstID
	return sendFocusMsg(firstID)
}

func (fm *FocusManager) FocusNextCmd(root UIModel) tea.Cmd {

	focusableItems := fm.getAllFocusable(root)
	if len(focusableItems) == 0 {
		return sendFocusMsg("")
	}
	if len(focusableItems) == 1 {
		return sendFocusMsg(focusableItems[0].Base().ID)
	}

	currentIndex := -1
	if fm.FocusedID != "" {
		for i, item := range focusableItems {
			if item.Base().ID == fm.FocusedID {
				currentIndex = i
				break
			}
		}
	}

	if currentIndex == -1 {
		return sendFocusMsg(focusableItems[0].Base().ID)
	}

	nextIndex := (currentIndex + 1) % len(focusableItems)
	nextID := focusableItems[nextIndex].Base().ID

	fm.FocusedID = nextID
	return sendFocusMsg(nextID)
}

func (fm *FocusManager) FocusPrevCmd(root UIModel) tea.Cmd {

	focusableItems := fm.getAllFocusable(root)
	if len(focusableItems) == 0 {
		return sendFocusMsg("")
	}
	if len(focusableItems) == 1 {
		return sendFocusMsg(focusableItems[0].Base().ID)
	}

	currentIndex := -1
	if fm.FocusedID != "" {
		for i, item := range focusableItems {
			if item.Base().ID == fm.FocusedID {
				currentIndex = i
				break
			}
		}
	}

	if currentIndex == -1 {
		return sendFocusMsg(focusableItems[len(focusableItems)-1].Base().ID)
	}

	prevIndex := (currentIndex - 1 + len(focusableItems)) % len(focusableItems)
	prevID := focusableItems[prevIndex].Base().ID

	fm.FocusedID = prevID
	return sendFocusMsg(prevID)
}
