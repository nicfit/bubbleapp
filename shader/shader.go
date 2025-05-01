package shader

import (
	"time"

	"github.com/charmbracelet/lipgloss/v2"
)

const (
	FPS = time.Second / 12
)

type Shader interface {
	Render(input string, style *lipgloss.Style) string
}

type DynamicShader interface {
	Render(input string, style *lipgloss.Style) string
	Tick()
}
