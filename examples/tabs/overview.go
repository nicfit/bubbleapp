package main

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
	"github.com/alexanderbh/bubbleapp/style"
)

func overview(c *app.Ctx) *app.C {
	return stack.New(c, func(c *app.Ctx) []*app.C {
		return []*app.C{
			text.New(c, "\nFor now you navigate tabs with arrow keys.\nThey should have shortcuts probably. And perhaps navigate with tab? Or vim keys?\n\n"),
			button.New(c, "Quit", c.Quit, button.WithVariant(style.Danger)),
			button.New(c, "Primary", c.Quit, button.WithVariant(style.Primary), button.WithMT(1)),
			button.New(c, "Secondary", c.Quit, button.WithVariant(style.Secondary), button.WithMT(1)),
			button.New(c, "Tertiary", c.Quit, button.WithVariant(style.Tertiary), button.WithMT(1)),
			button.New(c, "Warning", c.Quit, button.WithVariant(style.Warning), button.WithMT(1)),
			button.New(c, "Success", c.Quit, button.WithVariant(style.Success), button.WithMT(1)),
			button.New(c, "Info", c.Quit, button.WithVariant(style.Info), button.WithMT(1)),
			button.New(c, "Base", c.Quit, button.WithVariant(style.Base), button.WithMT(1)),
		}
	})
}
