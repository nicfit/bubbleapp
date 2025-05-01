package shader

import "github.com/alexanderbh/bubbleapp/style"

type SmallCapsShader struct {
	lastInput string
	output    string
}

func (s *SmallCapsShader) Render(input string) string {
	if s.lastInput == input {
		return s.output
	}
	s.output = style.ConvertToSmallCaps(input)
	s.lastInput = input
	return s.output
}
func NewSmallCapsShader() Shader {
	return &SmallCapsShader{}
}
