package main

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/divider"
	"github.com/alexanderbh/bubbleapp/component/loader"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
)

func loaders(ctx *app.Ctx, _ app.Props) string {
	return stack.New(ctx, func(ctx *app.Ctx) {
		text.New(ctx, "Loaders share the same tick for performance")
		divider.New(ctx)
		loader.NewWithoutText(ctx, loader.Dots)
		loader.NewWithoutText(ctx, loader.Arrow2, loader.WithColor(ctx.Styles.Colors.Secondary))
		loader.NewWithoutText(ctx, loader.Dots13, loader.WithColor(ctx.Styles.Colors.Tertiary))
		loader.NewWithoutText(ctx, loader.Line, loader.WithColor(ctx.Styles.Colors.Info))
		loader.NewWithoutText(ctx, loader.Balloon2, loader.WithColor(ctx.Styles.Colors.Success))
		loader.NewWithoutText(ctx, loader.Clock, loader.WithColor(ctx.Styles.Colors.Warning))
		loader.NewWithoutText(ctx, loader.Circle, loader.WithColor(ctx.Styles.Colors.Danger))
		loader.New(ctx, loader.CircleQuarters, "Loading...", loader.WithColor(ctx.Styles.Colors.Primary), loader.WithTextColor(ctx.Styles.Colors.Primary))
		divider.New(ctx)
		loader.NewWithoutText(ctx, loader.Dots12)
		loader.NewWithoutText(ctx, loader.Arc, loader.WithColor(ctx.Styles.Colors.Secondary))
		loader.NewWithoutText(ctx, loader.Toggle11, loader.WithColor(ctx.Styles.Colors.Tertiary))
		loader.NewWithoutText(ctx, loader.Line, loader.WithColor(ctx.Styles.Colors.Info))
		loader.NewWithoutText(ctx, loader.Star2, loader.WithColor(ctx.Styles.Colors.Success))
		loader.NewWithoutText(ctx, loader.AestheticSmall, loader.WithColor(ctx.Styles.Colors.Warning))
		loader.NewWithoutText(ctx, loader.Toggle2, loader.WithColor(ctx.Styles.Colors.Danger))
		loader.New(ctx, loader.Dots4, "Loading...", loader.WithColor(ctx.Styles.Colors.Primary), loader.WithTextColor(ctx.Styles.Colors.Primary))
		divider.New(ctx)
		loader.NewWithoutText(ctx, loader.Triangle)
		loader.NewWithoutText(ctx, loader.Arrow2, loader.WithColor(ctx.Styles.Colors.Secondary))
		loader.NewWithoutText(ctx, loader.Star, loader.WithColor(ctx.Styles.Colors.Tertiary))
		loader.NewWithoutText(ctx, loader.Line, loader.WithColor(ctx.Styles.Colors.Info))
		loader.NewWithoutText(ctx, loader.Dots7, loader.WithColor(ctx.Styles.Colors.Success))
		loader.NewWithoutText(ctx, loader.Dots8, loader.WithColor(ctx.Styles.Colors.Warning))
		loader.NewWithoutText(ctx, loader.Dots9, loader.WithColor(ctx.Styles.Colors.Danger))
		loader.New(ctx, loader.Dots10, "Loading...", loader.WithColor(ctx.Styles.Colors.Primary), loader.WithTextColor(ctx.Styles.Colors.Primary))
	})
}
