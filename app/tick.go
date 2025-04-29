package app

import (
	"time"

	tea "github.com/charmbracelet/bubbletea/v2"
)

type TickMsg time.Time

func (m Base) tick() tea.Cmd {
	return tea.Tick(m.Opts.TickFPS, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}
