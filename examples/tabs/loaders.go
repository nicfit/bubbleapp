package main

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/divider"
	"github.com/alexanderbh/bubbleapp/component/loader"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
)

func loaders(c *app.Ctx, _ app.Props) app.C {
	return stack.New(c, func(c *app.Ctx) []app.C {
		return []app.C{
			text.New(c, "Loaders share the same tick for performance"),
			divider.New(c),
			loader.NewWithoutText(c, loader.Dots),
			loader.NewWithoutText(c, loader.Arrow2, loader.WithColor(c.Styles.Colors.Secondary)),
			loader.NewWithoutText(c, loader.Dots13, loader.WithColor(c.Styles.Colors.Tertiary)),
			loader.NewWithoutText(c, loader.Line, loader.WithColor(c.Styles.Colors.Info)),
			loader.NewWithoutText(c, loader.Balloon2, loader.WithColor(c.Styles.Colors.Success)),
			loader.NewWithoutText(c, loader.Clock, loader.WithColor(c.Styles.Colors.Warning)),
			loader.NewWithoutText(c, loader.Circle, loader.WithColor(c.Styles.Colors.Danger)),
			loader.New(c, loader.CircleQuarters, "Loading...", loader.WithColor(c.Styles.Colors.Primary), loader.WithTextColor(c.Styles.Colors.Primary)),
			divider.New(c),
			loader.NewWithoutText(c, loader.Dots12),
			loader.NewWithoutText(c, loader.Arc, loader.WithColor(c.Styles.Colors.Secondary)),
			loader.NewWithoutText(c, loader.Toggle11, loader.WithColor(c.Styles.Colors.Tertiary)),
			loader.NewWithoutText(c, loader.Line, loader.WithColor(c.Styles.Colors.Info)),
			loader.NewWithoutText(c, loader.Star2, loader.WithColor(c.Styles.Colors.Success)),
			loader.NewWithoutText(c, loader.AestheticSmall, loader.WithColor(c.Styles.Colors.Warning)),
			loader.NewWithoutText(c, loader.Toggle2, loader.WithColor(c.Styles.Colors.Danger)),
			loader.New(c, loader.Dots4, "Loading...", loader.WithColor(c.Styles.Colors.Primary), loader.WithTextColor(c.Styles.Colors.Primary)),
			divider.New(c),
			loader.NewWithoutText(c, loader.Triangle),
			loader.NewWithoutText(c, loader.Arrow2, loader.WithColor(c.Styles.Colors.Secondary)),
			loader.NewWithoutText(c, loader.Star, loader.WithColor(c.Styles.Colors.Tertiary)),
			loader.NewWithoutText(c, loader.Line, loader.WithColor(c.Styles.Colors.Info)),
			loader.NewWithoutText(c, loader.Dots7, loader.WithColor(c.Styles.Colors.Success)),
			loader.NewWithoutText(c, loader.Dots8, loader.WithColor(c.Styles.Colors.Warning)),
			loader.NewWithoutText(c, loader.Dots9, loader.WithColor(c.Styles.Colors.Danger)),
			loader.New(c, loader.Dots10, "Loading...", loader.WithColor(c.Styles.Colors.Primary), loader.WithTextColor(c.Styles.Colors.Primary)),
		}
	})
}
