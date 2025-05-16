package main

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
)

func overview(c *app.Ctx) app.C {
	return stack.New(c, func(c *app.Ctx) []app.C {
		return []app.C{
			text.New(c, "\nFor now you navigate tabs with arrow keys.\nThey should have shortcuts probably. And perhaps navigate with tab? Or vim keys?\n\n"),
			button.New(c, "Quit", c.Quit, button.WithVariant(button.Danger)),
		}
	})
}
