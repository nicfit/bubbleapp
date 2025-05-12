package main

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
)

type authProps struct {
	userID string
}

func NewAuthModel(ctx *app.Ctx, props app.Props) string {
	authProps, ok := props.(authProps)
	if !ok {
		panic("NewAuthModel: props must be of type authProps")
	}

	return stack.New(ctx, func(ctx *app.Ctx) {
		text.New(ctx, "You are logged in as: "+authProps.userID)
		text.New(ctx, "Press [ctrl-c] to quit.\n", text.WithFg(ctx.Styles.Colors.Danger))
	})
}
