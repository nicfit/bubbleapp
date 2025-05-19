package main

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/context"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
)

func NewAuthModel(c *app.Ctx) app.C {
	appData := context.UseContext(c, AppDataContext)

	return stack.New(c, func(c *app.Ctx) []app.C {
		return []app.C{
			text.New(c, "You are logged in as: "+appData.data.userID),
			text.New(c, "Press [ctrl-c] to quit.\n", text.WithFg(c.Theme.Colors.DangerFg)),
		}
	})
}
