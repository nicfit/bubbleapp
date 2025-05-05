package app

import (
	"github.com/charmbracelet/lipgloss/v2"
)

// Maybe we want to have a raw string shader and then a lipgloss shader
// that just applies the lipgloss style to the string. Right now it is
// doing both which is not ideal.

type Shader interface {
	Render(input string, style *lipgloss.Style) string
}

type DynamicShader interface {
	Render(input string, style *lipgloss.Style) string
	Tick()
}
