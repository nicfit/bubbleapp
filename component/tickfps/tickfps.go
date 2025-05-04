package tickfps

import (
	"fmt"
	"time"

	"github.com/alexanderbh/bubbleapp/app"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type model[T any] struct {
	base      *app.Base[T]
	msgCount  map[string]int64
	tickTimes []time.Time
}

func New[T any](ctx *app.Context[T], baseOptions ...app.BaseOption) *app.Base[T] {
	if baseOptions == nil {
		baseOptions = []app.BaseOption{}
	}
	return model[T]{
		base:      app.NewBase(ctx, baseOptions...),
		msgCount:  make(map[string]int64),
		tickTimes: make([]time.Time, 0),
	}.Base()
}

func (m model[T]) Init() tea.Cmd {
	return nil
}

func (m model[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m model[T]) View() string {
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

func (m model[T]) Base() *app.Base[T] {
	m.base.Model = m
	return m.base
}
