package main

import (
	"os"
	"time"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
	"github.com/alexanderbh/bubbleapp/shader"
	"github.com/alexanderbh/bubbleapp/style"

	zone "github.com/alexanderbh/bubblezone/v2"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

func NewRoot() model {
	ctx := &app.Context[struct{}]{
		Styles: style.DefaultStyles(),
		Zone:   zone.New(),
	}

	stack := stack.New(ctx, &stack.Options[struct{}]{
		Children: []*app.Base[struct{}]{
			text.New(ctx, "Shader examples:", nil),
			text.New(ctx, "Small Caps Shader", &text.Options{
				Foreground: ctx.Styles.Colors.Primary,
			}, app.WithShader(shader.NewSmallCapsShader())),
			button.New(ctx, " Blink ", &button.Options{
				Variant: button.Danger,
			}, app.WithShader(shader.NewBlinkShader(time.Second/3, lipgloss.NewStyle().Foreground(ctx.Styles.Colors.Success).BorderForeground(ctx.Styles.Colors.Success)))),
		},
	}, app.AsRoot())

	return model{
		base: stack,
	}
}

type model struct {
	base *app.Base[struct{}]
}

func (m model) Init() tea.Cmd {
	return m.base.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	cmd := m.base.Update(msg)

	return m, cmd

}

func (m model) View() string {
	return m.base.Render()
}

func main() {
	p := tea.NewProgram(NewRoot(), tea.WithAltScreen(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
