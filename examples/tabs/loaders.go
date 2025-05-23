package main

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/divider"
	"github.com/alexanderbh/bubbleapp/component/loader"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
)

func loaders(c *app.Ctx) *app.C {
	return stack.New(c, func(c *app.Ctx) []*app.C {
		return []*app.C{
			text.New(c, "Loaders share the same tick for performance"),
			divider.New(c),
			loader.NewWithoutText(c, loader.Dots),
			loader.NewWithoutText(c, loader.Binary, loader.WithColor(c.Theme.Colors.SecondaryFg)),
			loader.NewWithoutText(c, loader.Dots13, loader.WithColor(c.Theme.Colors.TertiaryFg)),
			loader.NewWithoutText(c, loader.Line, loader.WithColor(c.Theme.Colors.InfoFg)),
			loader.NewWithoutText(c, loader.Balloon2, loader.WithColor(c.Theme.Colors.SuccessFg)),
			loader.NewWithoutText(c, loader.CircleHalves, loader.WithColor(c.Theme.Colors.WarningFg)),
			loader.NewWithoutText(c, loader.Circle, loader.WithColor(c.Theme.Colors.DangerFg)),
			loader.New(c, loader.CircleQuarters, "Loading...", loader.WithColor(c.Theme.Colors.PrimaryFg), loader.WithTextColor(c.Theme.Colors.PrimaryFg)),
			divider.New(c),
			loader.NewWithoutText(c, loader.Dots12),
			loader.NewWithoutText(c, loader.Arc, loader.WithColor(c.Theme.Colors.PrimaryFg)),
			loader.NewWithoutText(c, loader.Toggle11, loader.WithColor(c.Theme.Colors.SecondaryFg)),
			loader.NewWithoutText(c, loader.Line, loader.WithColor(c.Theme.Colors.TertiaryFg)),
			loader.NewWithoutText(c, loader.Star2, loader.WithColor(c.Theme.Colors.SuccessFg)),
			loader.NewWithoutText(c, loader.AestheticSmall, loader.WithColor(c.Theme.Colors.WarningFg)),
			loader.NewWithoutText(c, loader.Toggle2, loader.WithColor(c.Theme.Colors.DangerFg)),
			loader.New(c, loader.Dots4, "Loading...", loader.WithColor(c.Theme.Colors.SecondaryFg), loader.WithTextColor(c.Theme.Colors.SecondaryFg)),
			divider.New(c),
			loader.NewWithoutText(c, loader.Triangle),
			loader.NewWithoutText(c, loader.Dots3, loader.WithColor(c.Theme.Colors.Base200)),
			loader.NewWithoutText(c, loader.Star, loader.WithColor(c.Theme.Colors.Base400)),
			loader.NewWithoutText(c, loader.Line, loader.WithColor(c.Theme.Colors.InfoFg)),
			loader.NewWithoutText(c, loader.Dots7, loader.WithColor(c.Theme.Colors.SuccessFg)),
			loader.NewWithoutText(c, loader.Dots8, loader.WithColor(c.Theme.Colors.WarningFg)),
			loader.NewWithoutText(c, loader.Dots9, loader.WithColor(c.Theme.Colors.DangerFg)),
			loader.New(c, loader.Dots10, "Loading...", loader.WithColor(c.Theme.Colors.TertiaryFg), loader.WithTextColor(c.Theme.Colors.TertiaryFg)),
		}
	})
}
