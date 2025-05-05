package shader

import (
	"time"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/charmbracelet/lipgloss/v2"
)

type BlinkShader struct {
	blinkStyle    lipgloss.Style
	ticksPerState int
	totalTicks    int
	frame         int
}

func (b *BlinkShader) Render(input string, stl *lipgloss.Style) string {
	if b.frame < b.ticksPerState {
		if stl != nil {
			return stl.Render(input)
		}
		return input
	}
	if stl != nil {
		return b.blinkStyle.Inherit(*stl).Render(input)
	}
	return b.blinkStyle.Render(input)
}
func (b *BlinkShader) Tick() {
	b.frame++
	if b.frame >= b.totalTicks {
		b.frame = 0
	}
}
func NewBlinkShader(fps time.Duration, blinkStyle lipgloss.Style) app.Shader {
	ticks := max(int(float64(fps)/float64(app.FPS)), 1)

	return &BlinkShader{
		ticksPerState: ticks,
		totalTicks:    ticks * 2,
		frame:         0,
		blinkStyle:    blinkStyle,
	}
}
