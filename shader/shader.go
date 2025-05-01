package shader

import (
	"time"
)

const (
	FPS = time.Second / 12
)

type Shader interface {
	Render(input string) string
}

type DynamicShader interface {
	Render(input string) string
	Tick()
}
