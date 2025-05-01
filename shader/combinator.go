package shader

import "github.com/charmbracelet/lipgloss/v2"

type CombinatorSharder struct {
	shaders []Shader
}

func (c *CombinatorSharder) Render(input string, stl *lipgloss.Style) string {
	var output string = input
	for _, shader := range c.shaders {
		output = shader.Render(output, stl)
	}
	return output
}
func (b *CombinatorSharder) Tick() {
	for _, shader := range b.shaders {
		if dynamicShader, ok := shader.(DynamicShader); ok {
			dynamicShader.Tick()
		}
	}
}
func NewCombinatorShader(shaders ...Shader) Shader {
	return &CombinatorSharder{
		shaders: shaders,
	}
}
