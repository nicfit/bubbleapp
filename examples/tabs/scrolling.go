package main

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/box"
	"github.com/alexanderbh/bubbleapp/component/markdown"
	"github.com/alexanderbh/bubbleapp/component/stack"

	tea "github.com/charmbracelet/bubbletea/v2"
)

func NewScrolling(ctx *app.Context) app.UIModel {

	box1 := box.New(ctx, box.WithBg(ctx.Styles.Colors.HighlightBackground))
	box1.AddChild(markdown.New(ctx, "## This does not work yet. It will soon I hope!"))

	box2 := box.New(ctx)
	box2.AddChild(markdown.New(ctx, `What is Lorem Ipsum?
Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.

Why do we use it?
It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout. The point of using Lorem Ipsum is that it has a more-or-less normal distribution of letters, as opposed to using 'Content here, content here', making it look like readable English. Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text, and a search for 'lorem ipsum' will uncover many web sites still in their infancy. Various versions have evolved over the years, sometimes by accident, sometimes on purpose (injected humour and the like).`))

	stackView := stack.New(ctx)
	stackView.AddChildren(
		box1, box2,
	)

	base := app.New(ctx)
	base.AddChild(stackView)

	return markdownModel{
		base: base,
	}
}

type markdownModel struct {
	base *app.Base
}

func (m markdownModel) Init() tea.Cmd {
	return m.base.Init()
}

func (m markdownModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmd := m.base.Update(msg)

	return m, cmd

}

func (m markdownModel) View() string {
	return m.base.View()
}

func (m markdownModel) Base() *app.Base {
	return m.base
}
