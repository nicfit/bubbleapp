package tickfps

import (
	"fmt"
	"time"

	"github.com/alexanderbh/bubbleapp/app"
)

// Props for the TickFPS functional component.
type Props struct {
	// RegisterTicks specifies the interval at which the component's internal logic
	// (like updating tick times) should be triggered via app.UseTick.
	RegisterTicks time.Duration
}

// TickFPS is a functional component that displays the Tick FPS and the count of processed ticks.
// It is primarily used for debugging tick events.
func TickFPS(c *app.Ctx, props app.Props) string {
	p, ok := props.(Props)
	if !ok {
		panic("TickFPS: props must be of type tickfps.Props")
	}

	// tickTimes stores the timestamps of recent ticks processed by this component.
	tickTimesVal, setTickTimes := app.UseState(c, []time.Time{})

	// tickMsgCount stores the count of tick messages processed by this component.
	tickMsgCountVal, setTickMsgCount := app.UseState(c, int64(0))

	// app.UseTick registers a function to be called at the interval specified by p.RegisterTicks.
	app.UseTick(c, p.RegisterTicks, func() {
		now := time.Now()
		currentTimes := tickTimesVal // Access state value directly
		newTimes := append(currentTimes, now)

		var keptTimes []time.Time
		pruneTimeLimit := now.Add(-10 * time.Second)
		for _, t := range newTimes {
			if t.After(pruneTimeLimit) {
				keptTimes = append(keptTimes, t)
			}
		}

		setTickTimes(keptTimes)
		setTickMsgCount(tickMsgCountVal + 1) // Access state value directly and update
	})

	currentTickTimes := tickTimesVal       // Access state value directly
	currentTickMsgCount := tickMsgCountVal // Access state value directly

	if len(currentTickTimes) < 2 {
		return fmt.Sprintf("Tick FPS: 0.00 (%d)", currentTickMsgCount)
	}

	delta := currentTickTimes[len(currentTickTimes)-1].Sub(currentTickTimes[0]).Seconds()
	if delta <= 0 {
		return fmt.Sprintf("Tick FPS: 0.00 (%d)", currentTickMsgCount)
	}

	fps := float64(len(currentTickTimes)-1) / delta
	return fmt.Sprintf("Tick FPS: %.2f (%d)", fps, currentTickMsgCount)
}

// New creates an instance of the TickFPS that registers its own tick listener
// with the specified interval. This is useful for debugging tick events.
func NewAtInterval(c *app.Ctx, registerTicks time.Duration) string {
	componentProps := Props{
		RegisterTicks: registerTicks,
	}
	return c.Render(TickFPS, componentProps)
}

// New creates an instance of the TickFPS component with a default interval of 1s.
func New(c *app.Ctx) string {
	return NewAtInterval(c, 1000*time.Millisecond)
}
