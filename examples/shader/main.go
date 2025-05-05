package main

import (
	"os"
	"time"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
	"github.com/alexanderbh/bubbleapp/shader"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type CustomData struct{}

func NewRoot(ctx *app.Context[CustomData]) app.Fc[CustomData] {

	blinkShader := shader.NewBlinkShader(time.Second/3, lipgloss.NewStyle().
		Foreground(ctx.Styles.Colors.Success).
		BorderForeground(ctx.Styles.Colors.Success))

	stack := stack.New(ctx, []app.Fc[CustomData]{
		text.New(ctx, "Shader examples:", nil),
		text.New(ctx, "Small Caps Shader", &text.Options{
			Foreground: ctx.Styles.Colors.Primary,
		}, app.WithShader(shader.NewSmallCapsShader())),
		button.New(ctx, " Blink ", app.Quit, &button.Options{
			Variant: button.Danger,
		}, app.WithShader(blinkShader)),
	}, nil)

	return stack
}

func main() {
	ctx := app.NewContext(&CustomData{})

	p := tea.NewProgram(app.NewApp(ctx, NewRoot), tea.WithAltScreen(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
