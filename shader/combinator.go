package shader

type CombinatorSharder struct {
	shaders []Shader
}

func (c *CombinatorSharder) Render(input string) string {
	var output string = input
	for _, shader := range c.shaders {
		output = shader.Render(output)
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
