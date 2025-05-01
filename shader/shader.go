package shader

import "github.com/alexanderbh/bubbleapp/style"

type Shader interface {
	Render(input string) string
}

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

type CombinatorSharder struct {
	shaders []Shader
}

func (c *CombinatorSharder) Render(input string) string {
	var output string
	for _, shader := range c.shaders {
		output = shader.Render(input)
	}
	return output
}
func NewCombinatorShader(shaders ...Shader) Shader {
	return &CombinatorSharder{
		shaders: shaders,
	}
}
