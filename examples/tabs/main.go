package main

import (
	"os"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/tabs"
	"github.com/alexanderbh/bubbleapp/style"

	zone "github.com/alexanderbh/bubblezone/v2"
	tea "github.com/charmbracelet/bubbletea/v2"
)

var tabsData = []tabs.TabElement[CustomData]{
	{Title: "Overview", Content: NewOverview},
	{Title: "Loaders", Content: NewLoaders},
	{Title: "Scolling", Content: NewScrolling},
}

type CustomData struct {
	HowCoolIsThis string
}

func NewRoot() model[CustomData] {
	ctx := &app.Context[CustomData]{
		Styles: style.DefaultStyles(),
		Zone:   zone.New(),
		Data: &CustomData{
			HowCoolIsThis: "Very cool!",
		},
	}

	tabs := tabs.New(ctx, tabsData)

	base := app.New(ctx, app.AsRoot())
	base.AddChild(tabs)

	return model[CustomData]{
		base: base,
	}
}

type model[T CustomData] struct {
	base *app.Base[T]
}

func (m model[T]) Init() tea.Cmd {
	return m.base.Init()
}

func (m model[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmd := m.base.Update(msg)
	return m, cmd
}

func (m model[T]) View() string {
	return m.base.Render()
}

func main() {
	p := tea.NewProgram(NewRoot(), tea.WithAltScreen(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
