package main

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/divider"
	"github.com/alexanderbh/bubbleapp/component/loader"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
)

func NewLoaders(ctx *app.Context[CustomData]) app.Fc[CustomData] {
	stack := stack.New(ctx, func(ctx *app.Context[CustomData]) []app.Fc[CustomData] {
		return []app.Fc[CustomData]{
			text.New(ctx, "Loaders share the same tick for performance", nil),
			divider.New(ctx),
			loader.NewWithoutText(ctx, loader.Dots, nil),
			loader.NewWithoutText(ctx, loader.Arrow2, &loader.Options{Color: ctx.Styles.Colors.Secondary}),
			loader.NewWithoutText(ctx, loader.Dots13, &loader.Options{Color: ctx.Styles.Colors.Tertiary}),
			loader.NewWithoutText(ctx, loader.Line, &loader.Options{Color: ctx.Styles.Colors.Info}),
			loader.NewWithoutText(ctx, loader.Balloon2, &loader.Options{Color: ctx.Styles.Colors.Success}),
			loader.NewWithoutText(ctx, loader.Clock, &loader.Options{Color: ctx.Styles.Colors.Warning}),
			loader.NewWithoutText(ctx, loader.Circle, &loader.Options{Color: ctx.Styles.Colors.Danger}),
			loader.New(ctx, loader.CircleQuarters, "Loading...", &loader.Options{Color: ctx.Styles.Colors.Primary, TextColor: ctx.Styles.Colors.Primary}),
			divider.New(ctx),
			loader.NewWithoutText(ctx, loader.Dots12, nil),
			loader.NewWithoutText(ctx, loader.Arc, &loader.Options{Color: ctx.Styles.Colors.Secondary}),
			loader.NewWithoutText(ctx, loader.Toggle11, &loader.Options{Color: ctx.Styles.Colors.Tertiary}),
			loader.NewWithoutText(ctx, loader.Line, &loader.Options{Color: ctx.Styles.Colors.Info}),
			loader.NewWithoutText(ctx, loader.Star2, &loader.Options{Color: ctx.Styles.Colors.Success}),
			loader.NewWithoutText(ctx, loader.AestheticSmall, &loader.Options{Color: ctx.Styles.Colors.Warning}),
			loader.NewWithoutText(ctx, loader.Toggle2, &loader.Options{Color: ctx.Styles.Colors.Danger}),
			loader.New(ctx, loader.Dots4, "Loading...", &loader.Options{Color: ctx.Styles.Colors.Primary, TextColor: ctx.Styles.Colors.Primary}),
			divider.New(ctx),
			loader.NewWithoutText(ctx, loader.Triangle, nil),
			loader.NewWithoutText(ctx, loader.Arrow2, &loader.Options{Color: ctx.Styles.Colors.Secondary}),
			loader.NewWithoutText(ctx, loader.Star, &loader.Options{Color: ctx.Styles.Colors.Tertiary}),
			loader.NewWithoutText(ctx, loader.Line, &loader.Options{Color: ctx.Styles.Colors.Info}),
			loader.NewWithoutText(ctx, loader.Dots7, &loader.Options{Color: ctx.Styles.Colors.Success}),
			loader.NewWithoutText(ctx, loader.Dots8, &loader.Options{Color: ctx.Styles.Colors.Warning}),
			loader.NewWithoutText(ctx, loader.Dots9, &loader.Options{Color: ctx.Styles.Colors.Danger}),
			loader.New(ctx, loader.Dots10, "Loading...", &loader.Options{Color: ctx.Styles.Colors.Primary, TextColor: ctx.Styles.Colors.Primary}),
		}
	}, nil)

	return stack
}
