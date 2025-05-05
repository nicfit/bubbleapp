package main

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
)

func NewOverview(ctx *app.Context[CustomData]) app.Fc[CustomData] {
	quitButton := button.New(ctx, "Quit", app.Quit, &button.Options{Variant: button.Danger})

	stack := stack.New(ctx, func(ctx *app.Context[CustomData]) []app.Fc[CustomData] {
		return []app.Fc[CustomData]{
			text.New(ctx, "\nFor now you navigate tabs with arrow keys.\nThey should have shortcuts probably. And perhaps navigate with tab? Or vim keys?\n\n", nil),
			text.New(ctx, "From global data: "+ctx.Data.HowCoolIsThis, nil),
			quitButton,
		}
	}, nil)

	return stack
}
