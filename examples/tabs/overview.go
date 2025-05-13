package main

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
)

func overview(c *app.Ctx, _ app.Props) string {
	return stack.New(c, func(ctx *app.Ctx) {
		text.New(ctx, "\nFor now you navigate tabs with arrow keys.\nThey should have shortcuts probably. And perhaps navigate with tab? Or vim keys?\n\n")
		button.New(c, "Quit", c.Quit, button.WithVariant(button.Danger))
	})
}
