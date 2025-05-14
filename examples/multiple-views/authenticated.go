package main

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/context"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
)

func NewAuthModel(ctx *app.Ctx, _ app.Props) string {
	appData := context.UseContext(ctx, AppDataContext)

	return stack.New(ctx, func(ctx *app.Ctx) {
		text.New(ctx, "You are logged in as: "+appData.data.userID)
		text.New(ctx, "Press [ctrl-c] to quit.\n", text.WithFg(ctx.Styles.Colors.Danger))
	})
}
