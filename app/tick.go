package app

import (
	"time"

	tea "github.com/charmbracelet/bubbletea/v2"
)

type TickMsg time.Time

func tickCommand(duration time.Duration) tea.Cmd {
	return tea.Tick(duration, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func tickVisitor[T any](node Fc[T], _ Fc[T], _ *Context[T]) {
	if node == nil {
		return
	}
	if node.Base() == nil {
		return
	}
	node.Base().Tick()
}
