package main

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
)

func NewAuthModel(ctx *app.Context[CustomData]) app.Fc[CustomData] {
	root := stack.New(ctx, []app.Fc[CustomData]{
		text.New(ctx, "You are logged in as: "+ctx.Data.UserID, nil),
		text.New(ctx, "Press [ctrl-c] to quit.\n", &text.Options{Foreground: ctx.Styles.Colors.Danger}),
	}, nil)

	return root
}
