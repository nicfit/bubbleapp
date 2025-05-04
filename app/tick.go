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
