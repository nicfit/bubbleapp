package app

import (
	"time"

	tea "github.com/charmbracelet/bubbletea/v2"
)

type TickMsg struct {
	OccurredAt time.Time
}

type tickState[T any] struct {
	tickListeners *[]tickListener
	timers        *map[time.Duration]*time.Timer
	lastTickTimes map[string]time.Time
}

type tickListener struct {
	interval time.Duration
	id       string
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
func (tick *tickState[T]) RegisterTickListener(interval time.Duration, id string) {
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

func (tick *tickState[T]) createTimer(ctx *Context[T]) {
	if len(*tick.tickListeners) == 0 {
		return
	}

	// Collect all intervals
	intervals := make([]time.Duration, 0, len(*tick.tickListeners))
	for _, listener := range *tick.tickListeners {
		// Convert interval to integer milliseconds (loss of precision is fine)
		ms := listener.interval.Milliseconds()
		intervals = append(intervals, time.Duration(ms)*time.Millisecond)
	}

	gcdInterval := gcdSlice(intervals)
	if gcdInterval == 0 {
		return
	}

	if tick.timers == nil {
		tick.timers = &map[time.Duration]*time.Timer{}
	}

	// Stop and remove all existing timers except the gcdInterval
	for interval, timer := range *tick.timers {
		if interval != gcdInterval {
			if timer != nil {
				timer.Stop()
				go func(t *time.Timer) {
					for len(t.C) > 0 {
						<-t.C
					}
				}(timer)
			}
			delete(*tick.timers, interval)
		}
	}

	if timer, ok := (*tick.timers)[gcdInterval]; ok && timer != nil {
		return // Timer already running
	}

	t := time.NewTimer(gcdInterval)
	(*tick.timers)[gcdInterval] = t

	go func(gcdInterval time.Duration, t *time.Timer, ctx *Context[T], listeners []tickListener) {
		var elapsed time.Duration
		for {
			<-t.C
			elapsed += gcdInterval
			ctx.teaProgram.Send(TickMsg{OccurredAt: time.Now()})
			t.Reset(gcdInterval)
		}
	}(gcdInterval, t, ctx, *tick.tickListeners)
}

func tickCommand(duration time.Duration) tea.Cmd {
	return tea.Tick(duration, func(_ time.Time) tea.Msg {
		return TickMsg{OccurredAt: time.Now()}
	})
}
