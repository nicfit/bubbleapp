package app

import (
	"sync" // Added for mutex
	"time"

	tea "github.com/charmbracelet/bubbletea/v2"
)

type TickMsg struct {
	OccurredAt time.Time
}

type tickState[T any] struct {
	tickListeners       *[]tickListener
	timers              *map[time.Duration]*time.Timer
	lastTickTimes       map[string]time.Time
	activeTimer         *time.Timer
	activeTimerDone     chan struct{}
	activeTimerInterval time.Duration
	mu                  sync.Mutex
}

type tickListener struct {
	interval time.Duration
	id       string
	callback func()
}

func (tick *tickState[T]) init() {
	tick.tickListeners = &[]tickListener{}
	tick.lastTickTimes = make(map[string]time.Time)
}

// Tell BubbleApp that the component with this ID wants to receive tick events at the given interval.
// No matter the internal tick interval, the component will receive tick events at the given interval.
//
// IMPORTANT: Use intervals with a large common divisor.
// Ticks will happen internally at the greatest common divisor of all intervals.
// This means that if you register a tick listener with 1 second and another with 2 seconds,
// the internal tick will be happen every 1 second.
// But if you register a tick listener with 80ms and another with 100ms,
// the internal tick will be happen every 20ms, which might hurt performance.
func (tick *tickState[T]) RegisterTickListener(interval time.Duration, id string, callback func()) {
	tick.mu.Lock()
	defer tick.mu.Unlock()
	if tick.tickListeners == nil {
		tick.tickListeners = &[]tickListener{}
	}
	if tick.timers == nil {
		tick.timers = &map[time.Duration]*time.Timer{}
	}
	if tick.lastTickTimes == nil {
		tick.lastTickTimes = make(map[string]time.Time)
	}

	*tick.tickListeners = append(*tick.tickListeners, tickListener{
		interval: interval,
		id:       id,
		callback: callback, // Store the callback
	})
}

func gcd(a, b time.Duration) time.Duration {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcdSlice(durations []time.Duration) time.Duration {
	if len(durations) == 0 {
		return 0
	}
	result := durations[0]
	for _, d := range durations[1:] {
		result = gcd(result, d)
	}
	return result
}

func (tick *tickState[T]) createTimer(ctx *Ctx) {
	tick.mu.Lock()
	defer tick.mu.Unlock()

	if tick.tickListeners == nil || len(*tick.tickListeners) == 0 {
		if tick.activeTimer != nil {
			tick.activeTimer.Stop()
			tick.activeTimer = nil
		}
		if tick.activeTimerDone != nil {
			select {
			case <-tick.activeTimerDone:
			default:
				close(tick.activeTimerDone)
			}
			tick.activeTimerDone = nil
		}
		tick.activeTimerInterval = 0
		if tick.timers != nil {
			*tick.timers = make(map[time.Duration]*time.Timer)
		}
		return
	}

	intervals := make([]time.Duration, 0, len(*tick.tickListeners))
	for _, listener := range *tick.tickListeners {
		ms := listener.interval.Milliseconds()
		intervals = append(intervals, time.Duration(ms)*time.Millisecond)
	}

	gcdInterval := max(12*time.Millisecond, gcdSlice(intervals)) // 12ms low limit. Maybe too low.

	if gcdInterval == 0 {
		if tick.activeTimer != nil {
			tick.activeTimer.Stop()
			tick.activeTimer = nil
		}
		if tick.activeTimerDone != nil {
			select {
			case <-tick.activeTimerDone:
			default:
				close(tick.activeTimerDone)
			}
			tick.activeTimerDone = nil
		}
		tick.activeTimerInterval = 0
		if tick.timers != nil {
			*tick.timers = make(map[time.Duration]*time.Timer)
		}
		return
	}

	if tick.activeTimer != nil && tick.activeTimerInterval == gcdInterval && tick.activeTimerDone != nil {
		return
	}

	if tick.activeTimer != nil {
		tick.activeTimer.Stop()
	}
	if tick.activeTimerDone != nil {
		select {
		case <-tick.activeTimerDone:
		default:
			close(tick.activeTimerDone) // Signal old goroutine to stop
		}
	}

	tick.activeTimerDone = make(chan struct{})
	newTimer := time.NewTimer(gcdInterval)
	tick.activeTimer = newTimer
	tick.activeTimerInterval = gcdInterval

	newTimersMap := make(map[time.Duration]*time.Timer)
	newTimersMap[gcdInterval] = newTimer
	if tick.timers == nil {
		tick.timers = &newTimersMap
	} else {
		*tick.timers = newTimersMap
	}

	go func(timer *time.Timer, done <-chan struct{}, program *tea.Program, interval time.Duration) {
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				tick.mu.Lock()
				listeners := *tick.tickListeners
				now := time.Now()
				for _, listener := range listeners {
					lastTick, ok := tick.lastTickTimes[listener.id]
					if !ok || now.Sub(lastTick) >= listener.interval {
						if listener.callback != nil {
							listener.callback()
						}
						tick.lastTickTimes[listener.id] = now
					}
				}
				tick.mu.Unlock()
				select {
				case <-done:
					return
				default:
					if timer != nil {
						timer.Reset(interval)
					}
				}
			case <-done:
				return
			}
		}
	}(newTimer, tick.activeTimerDone, ctx.teaProgram, gcdInterval)
}

func (tick *tickState[T]) StopActiveTimer() {
	tick.mu.Lock()
	defer tick.mu.Unlock()

	if tick.activeTimerDone != nil {
		select {
		case <-tick.activeTimerDone:
			// Channel already closed
		default:
			close(tick.activeTimerDone)
		}
		tick.activeTimerDone = nil
	}

	if tick.activeTimer != nil {
		tick.activeTimer.Stop()
		tick.activeTimer = nil
	}
}

func (tick *tickState[T]) UnregisterTickListener(id string) {
	tick.mu.Lock()
	defer tick.mu.Unlock()

	if tick.tickListeners == nil {
		return
	}

	found := false
	newListeners := []tickListener{}
	for _, listener := range *tick.tickListeners {
		if listener.id != id {
			newListeners = append(newListeners, listener)
		} else {
			found = true
		}
	}

	if found {
		*tick.tickListeners = newListeners
		delete(tick.lastTickTimes, id)

		if len(*tick.tickListeners) == 0 {
			tick.StopActiveTimer()
		}
	}
}

func tickCommand(duration time.Duration) tea.Cmd {
	return tea.Tick(duration, func(_ time.Time) tea.Msg {
		return TickMsg{OccurredAt: time.Now()}
	})
}
