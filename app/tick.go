package app

import (
	"time"

	tea "github.com/charmbracelet/bubbletea/v2"
)

type TickMsg struct{}

func tickCommand(duration time.Duration) tea.Cmd {
	return tea.Tick(duration, func(_ time.Time) tea.Msg {
		return TickMsg{}
	})
}

func tickVisitor[T any](node Fc[T], _ Fc[T], ctx *Context[T]) {
	if node == nil {
		return
	}
	if node.Base() == nil {
		return
	}
	node.Update(ctx, TickMsg{})
	node.Base().Tick()
}
