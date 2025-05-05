package shader

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/style"
	"github.com/charmbracelet/lipgloss/v2"
)

type SmallCapsShader struct {
	lastInput string
	output    string
}

func (s *SmallCapsShader) Render(input string, stl *lipgloss.Style) string {
	if s.lastInput == input {
		if stl != nil {
			return stl.Render(s.output)
		}
		return s.output
	}
	s.output = style.ConvertToSmallCaps(input)
	s.lastInput = input
	if stl != nil {
		return stl.Render(s.output)
	}
	return s.output
}
func NewSmallCapsShader() app.Shader {
	return &SmallCapsShader{}
}
