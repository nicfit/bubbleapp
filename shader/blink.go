package shader

import (
	"strings"
	"time"
	"unicode/utf8"
)

type BlinkShader struct {
	blinkCache    string
	lastInput     string
	ticksPerState int
	totalTicks    int
	frame         int
}

func (b *BlinkShader) Render(input string) string {
	if b.frame < b.ticksPerState {
		return input
	}
	if b.lastInput != input || b.blinkCache == "" {
		b.lastInput = input
		runeCount := utf8.RuneCountInString(input)
		b.blinkCache = strings.Repeat(" ", runeCount)
	}
	return b.blinkCache
}
func (b *BlinkShader) Tick() {
	b.frame++
	if b.frame >= b.totalTicks {
		b.frame = 0
	}
}
func NewBlinkShader(fps time.Duration) Shader {
	ticks := int(float64(fps) / float64(FPS))

	if ticks < 1 {
		ticks = 1
	}

	return &BlinkShader{
		ticksPerState: ticks,
		totalTicks:    ticks * 2,
		frame:         0,
	}
}
