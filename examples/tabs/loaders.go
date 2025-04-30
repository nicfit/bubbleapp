package main

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/divider"
	"github.com/alexanderbh/bubbleapp/component/loader"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"

	tea "github.com/charmbracelet/bubbletea/v2"
)

func NewLoaders(ctx *app.Context[CustomData]) *app.Base[CustomData] {
	stack := stack.New(ctx, &stack.Options[CustomData]{
		Children: []*app.Base[CustomData]{
			text.New(ctx, "Loaders share the same tick for performance", nil),
			divider.New(ctx),
			loader.New(ctx, loader.Dot, nil),
			loader.New(ctx, loader.Ellipsis, &loader.Options{Color: ctx.Styles.Colors.Secondary}),
			loader.New(ctx, loader.Jump, &loader.Options{Color: ctx.Styles.Colors.Tertiary}),
			loader.New(ctx, loader.Line, &loader.Options{Color: ctx.Styles.Colors.Info}),
			loader.New(ctx, loader.Meter, &loader.Options{Color: ctx.Styles.Colors.Success}),
			loader.New(ctx, loader.MiniDot, &loader.Options{Color: ctx.Styles.Colors.Warning}),
			loader.New(ctx, loader.Points, &loader.Options{Color: ctx.Styles.Colors.Danger}),
			loader.New(ctx, loader.Pulse, &loader.Options{Color: ctx.Styles.Colors.Primary, Text: "Loading..."}),
			divider.New(ctx),
			loader.New(ctx, loader.Dot, nil),
			loader.New(ctx, loader.Ellipsis, &loader.Options{Color: ctx.Styles.Colors.Secondary}),
			loader.New(ctx, loader.Jump, &loader.Options{Color: ctx.Styles.Colors.Tertiary}),
			loader.New(ctx, loader.Line, &loader.Options{Color: ctx.Styles.Colors.Info}),
			loader.New(ctx, loader.Meter, &loader.Options{Color: ctx.Styles.Colors.Success}),
			loader.New(ctx, loader.MiniDot, &loader.Options{Color: ctx.Styles.Colors.Warning}),
			loader.New(ctx, loader.Points, &loader.Options{Color: ctx.Styles.Colors.Danger}),
			loader.New(ctx, loader.Pulse, &loader.Options{Color: ctx.Styles.Colors.Primary, Text: "Loading..."}),
			divider.New(ctx),
			loader.New(ctx, loader.Dot, nil),
			loader.New(ctx, loader.Ellipsis, &loader.Options{Color: ctx.Styles.Colors.Secondary}),
			loader.New(ctx, loader.Jump, &loader.Options{Color: ctx.Styles.Colors.Tertiary}),
			loader.New(ctx, loader.Line, &loader.Options{Color: ctx.Styles.Colors.Info}),
			loader.New(ctx, loader.Meter, &loader.Options{Color: ctx.Styles.Colors.Success}),
			loader.New(ctx, loader.MiniDot, &loader.Options{Color: ctx.Styles.Colors.Warning}),
			loader.New(ctx, loader.Points, &loader.Options{Color: ctx.Styles.Colors.Danger}),
			loader.New(ctx, loader.Pulse, &loader.Options{Color: ctx.Styles.Colors.Primary, Text: "Loading..."}),
		}},
	)

	base := app.New(ctx)
	base.AddChild(stack)

	return loadersModel[CustomData]{
		base: base,
	}.Base()
}

type loadersModel[T CustomData] struct {
	base *app.Base[CustomData]
}

func (m loadersModel[T]) Init() tea.Cmd {
	return m.base.Init()
}

func (m loadersModel[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmd := m.base.Update(msg)

	return m, cmd

}

func (m loadersModel[T]) View() string {
	return m.base.Render()
}

func (m loadersModel[T]) Base() *app.Base[CustomData] {
	m.base.Model = m
	return m.base
}
