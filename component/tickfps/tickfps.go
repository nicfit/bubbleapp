package tickfps

import (
	"fmt"
	"time"

	"github.com/alexanderbh/bubbleapp/app"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type Model struct {
	base      *app.Base
	msgCount  map[string]int64
	tickTimes []time.Time
}

func New(ctx *app.Context) Model {
	return Model{
		base:      app.New(ctx),
		msgCount:  make(map[string]int64),
		tickTimes: make([]time.Time, 0),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msgType := msg.(type) {
	case app.TickMsg:
		m.tickTimes = append(m.tickTimes, time.Now())
		cutoff := time.Now().Add(-10 * time.Second)
		idx := 0
		for i, t := range m.tickTimes {
			if t.After(cutoff) {
				idx = i
				break
			}
		}
		m.tickTimes = m.tickTimes[idx:]
		m.msgCount["components.TickMsg"]++
	default:
		m.msgCount[fmt.Sprintf("%T", msgType)]++
	}
	return m, nil
}

func (m Model) View() string {
	if len(m.tickTimes) < 2 {
		return "Tick FPS: 0.00"
	}
	delta := m.tickTimes[len(m.tickTimes)-1].Sub(m.tickTimes[0]).Seconds()
	if delta == 0 {
		return "Tick FPS: 0.00"
	}
	fps := float64(len(m.tickTimes)-1) / delta
	return fmt.Sprintf("Tick FPS: %.2f", fps)
}

func (m Model) Base() *app.Base {
	return m.base
}
