package main

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/divider"
	"github.com/alexanderbh/bubbleapp/component/loader"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"

	tea "github.com/charmbracelet/bubbletea/v2"
)

func NewLoaders(ctx *app.Context[CustomData]) app.UIModel[CustomData] {
	stack := stack.New(ctx, stack.Options[CustomData]{
		Children: []*app.Base[CustomData]{
			text.New(ctx, "Loaders share the same tick for performance").Base(),
			divider.New(ctx).Base(),
			loader.New(ctx, loader.Dot).Base(),
			loader.New(ctx, loader.Ellipsis, loader.WithColor(ctx.Styles.Colors.Secondary)).Base(),
			loader.New(ctx, loader.Jump, loader.WithColor(ctx.Styles.Colors.Tertiary)).Base(),
			loader.New(ctx, loader.Line, loader.WithColor(ctx.Styles.Colors.Info)).Base(),
			loader.New(ctx, loader.Meter, loader.WithColor(ctx.Styles.Colors.Success)).Base(),
			loader.New(ctx, loader.MiniDot, loader.WithColor(ctx.Styles.Colors.Warning)).Base(),
			loader.New(ctx, loader.Points, loader.WithColor(ctx.Styles.Colors.Danger)).Base(),
			loader.New(ctx, loader.Pulse, loader.WithColor(ctx.Styles.Colors.Primary), loader.WithText("Loading...")).Base(),
			divider.New(ctx).Base(),
			loader.New(ctx, loader.Dot).Base(),
			loader.New(ctx, loader.Ellipsis, loader.WithColor(ctx.Styles.Colors.Secondary)).Base(),
			loader.New(ctx, loader.Jump, loader.WithColor(ctx.Styles.Colors.Tertiary)).Base(),
			loader.New(ctx, loader.Line, loader.WithColor(ctx.Styles.Colors.Info)).Base(),
			loader.New(ctx, loader.Meter, loader.WithColor(ctx.Styles.Colors.Success)).Base(),
			loader.New(ctx, loader.MiniDot, loader.WithColor(ctx.Styles.Colors.Warning)).Base(),
			loader.New(ctx, loader.Points, loader.WithColor(ctx.Styles.Colors.Danger)).Base(),
			loader.New(ctx, loader.Pulse, loader.WithColor(ctx.Styles.Colors.Primary), loader.WithText("Loading...")).Base(),
			divider.New(ctx).Base(),
			loader.New(ctx, loader.Dot).Base(),
			loader.New(ctx, loader.Ellipsis, loader.WithColor(ctx.Styles.Colors.Secondary)).Base(),
			loader.New(ctx, loader.Jump, loader.WithColor(ctx.Styles.Colors.Tertiary)).Base(),
			loader.New(ctx, loader.Line, loader.WithColor(ctx.Styles.Colors.Info)).Base(),
			loader.New(ctx, loader.Meter, loader.WithColor(ctx.Styles.Colors.Success)).Base(),
			loader.New(ctx, loader.MiniDot, loader.WithColor(ctx.Styles.Colors.Warning)).Base(),
			loader.New(ctx, loader.Points, loader.WithColor(ctx.Styles.Colors.Danger)).Base(),
			loader.New(ctx, loader.Pulse, loader.WithColor(ctx.Styles.Colors.Primary), loader.WithText("Loading...")).Base(),
		}},
	)

	base := app.New(ctx)
	base.AddChild(stack.Base())

	return loadersModel[CustomData]{
		base: base,
	}
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
