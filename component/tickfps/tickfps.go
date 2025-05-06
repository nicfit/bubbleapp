package tickfps

import (
	"fmt"
	"time"

	"github.com/alexanderbh/bubbleapp/app"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type uiState struct {
	msgCount  map[string]int64
	tickTimes []time.Time
}

type tickfps[T any] struct {
	base *app.Base
}

// USED FOR DEBUGGING TICK EVENTS.
func New[T any](ctx *app.Context[T], registerTicks time.Duration, baseOptions ...app.BaseOption) *tickfps[T] {
	if baseOptions == nil {
		baseOptions = []app.BaseOption{}
	}
	base, cleanup := app.NewBase(ctx, "tickfps", baseOptions...)
	defer cleanup()

	ctx.Tick.RegisterTickListener(registerTicks, base.ID)

	return &tickfps[T]{
		base: base,
	}
}

func (m *tickfps[T]) Update(ctx *app.Context[T], msg tea.Msg) {
	state := m.getState(ctx)
	switch msgType := msg.(type) {
	case app.TickMsg:
		state.tickTimes = append(state.tickTimes, time.Now())
		cutoff := time.Now().Add(-10 * time.Second)
		idx := 0
		for i, t := range state.tickTimes {
			if t.After(cutoff) {
				idx = i
				break
			}
		}
		state.tickTimes = state.tickTimes[idx:]
		state.msgCount["components.TickMsg"]++
	default:
		state.msgCount[fmt.Sprintf("%T", msgType)]++
	}

}

func (m *tickfps[T]) Render(ctx *app.Context[T]) string {
	state := m.getState(ctx)
	if len(state.tickTimes) < 2 {
		return "Tick FPS: 0.00"
	}
	delta := state.tickTimes[len(state.tickTimes)-1].Sub(state.tickTimes[0]).Seconds()
	if delta == 0 {
		return "Tick FPS: 0.00"
	}
	fps := float64(len(state.tickTimes)-1) / delta
	return fmt.Sprintf("Tick FPS: %.2f (%d)", fps, state.msgCount["components.TickMsg"])
}

func (m *tickfps[T]) Children(ctx *app.Context[T]) []app.Fc[T] {
	return nil
}
func (m *tickfps[T]) Base() *app.Base {
	return m.base
}

func (m *tickfps[T]) getState(ctx *app.Context[T]) *uiState {
	state := app.GetUIState[T, uiState](ctx, m.base.ID)
	if state == nil {
		state = &uiState{
			msgCount:  make(map[string]int64),
			tickTimes: make([]time.Time, 0),
		}
		app.SetUIState(ctx, m.base.ID, state)
	}
	return state
}
