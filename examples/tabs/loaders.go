package main

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/divider"
	"github.com/alexanderbh/bubbleapp/component/loader"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"

	tea "github.com/charmbracelet/bubbletea/v2"
)

func NewLoaders(ctx *app.Context) app.UIModel {
	stack := stack.New(ctx)
	stack.AddChildren(
		text.New(ctx, "Loaders share the same tick for performance"),
		divider.New(ctx),
		loader.New(ctx, loader.Dot),
		loader.New(ctx, loader.Ellipsis, loader.WithColor(ctx.Styles.Colors.Secondary)),
		loader.New(ctx, loader.Jump, loader.WithColor(ctx.Styles.Colors.Tertiary)),
		loader.New(ctx, loader.Line, loader.WithColor(ctx.Styles.Colors.Info)),
		loader.New(ctx, loader.Meter, loader.WithColor(ctx.Styles.Colors.Success)),
		loader.New(ctx, loader.MiniDot, loader.WithColor(ctx.Styles.Colors.Warning)),
		loader.New(ctx, loader.Points, loader.WithColor(ctx.Styles.Colors.Danger)),
		loader.New(ctx, loader.Pulse, loader.WithColor(ctx.Styles.Colors.Primary), loader.WithText("Loading...")),
		divider.New(ctx),
		loader.New(ctx, loader.Dot),
		loader.New(ctx, loader.Ellipsis, loader.WithColor(ctx.Styles.Colors.Secondary)),
		loader.New(ctx, loader.Jump, loader.WithColor(ctx.Styles.Colors.Tertiary)),
		loader.New(ctx, loader.Line, loader.WithColor(ctx.Styles.Colors.Info)),
		loader.New(ctx, loader.Meter, loader.WithColor(ctx.Styles.Colors.Success)),
		loader.New(ctx, loader.MiniDot, loader.WithColor(ctx.Styles.Colors.Warning)),
		loader.New(ctx, loader.Points, loader.WithColor(ctx.Styles.Colors.Danger)),
		loader.New(ctx, loader.Pulse, loader.WithColor(ctx.Styles.Colors.Primary), loader.WithText("Loading...")),
		divider.New(ctx),
		loader.New(ctx, loader.Dot),
		loader.New(ctx, loader.Ellipsis, loader.WithColor(ctx.Styles.Colors.Secondary)),
		loader.New(ctx, loader.Jump, loader.WithColor(ctx.Styles.Colors.Tertiary)),
		loader.New(ctx, loader.Line, loader.WithColor(ctx.Styles.Colors.Info)),
		loader.New(ctx, loader.Meter, loader.WithColor(ctx.Styles.Colors.Success)),
		loader.New(ctx, loader.MiniDot, loader.WithColor(ctx.Styles.Colors.Warning)),
		loader.New(ctx, loader.Points, loader.WithColor(ctx.Styles.Colors.Danger)),
		loader.New(ctx, loader.Pulse, loader.WithColor(ctx.Styles.Colors.Primary), loader.WithText("Loading...")),
	)

	base := app.New(ctx)
	base.AddChild(stack)

	return loadersModel{
		base: base,
	}
}

type loadersModel struct {
	base *app.Base
}

func (m loadersModel) Init() tea.Cmd {
	return m.base.Init()
}

func (m loadersModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmd := m.base.Update(msg)

	return m, cmd

}

func (m loadersModel) View() string {
	return m.base.View()
}

func (m loadersModel) Base() *app.Base {
	return m.base
}
