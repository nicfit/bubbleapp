package main

import (
	"time"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/loader"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
)

type CustomData struct {
	LoginingIn  bool
	LoginFailed string
	UserID      string
}

func NewLoginRoot(ctx *app.Context[CustomData]) app.Fc[CustomData] {
	if ctx.Data.UserID != "" {
		return NewAuthModel(ctx)
	}

	if ctx.Data.LoginingIn {
		return stack.New(ctx, []app.Fc[CustomData]{
			text.New(ctx, "Please wait...", nil),
			loader.New(ctx, loader.Binary, "Logging in...", nil),
		}, nil)
	}

	loginButton := button.New(ctx, "Log in", func(ctx *app.Context[CustomData]) {
		go LoginSuperSecure(ctx.Data, false)
	}, &button.Options{Variant: button.Primary, Type: button.Compact})

	failButton := button.New(ctx, "Fail log in", func(ctx *app.Context[CustomData]) {
		go LoginSuperSecure(ctx.Data, true)
	}, &button.Options{Variant: button.Warning, Type: button.Compact})

	quitButton := button.New(ctx, "Quit App", app.Quit, &button.Options{Variant: button.Danger, Type: button.Compact})

	views := []app.Fc[CustomData]{
		text.New(ctx, "██       ██████   ██████  ██ ███    ██\n██      ██    ██ ██       ██ ████   ██\n██      ██    ██ ██   ███ ██ ██ ██  ██\n██      ██    ██ ██    ██ ██ ██  ██ ██\n███████  ██████   ██████  ██ ██   ████\n\n", nil),
		text.New(ctx, "Log in or fail! Up to you!", nil),
		loginButton,
		failButton,
		quitButton,
	}
	if ctx.Data.LoginFailed != "" {
		views = append(views, text.New(ctx, "\n"+ctx.Data.LoginFailed, &text.Options{Foreground: ctx.Styles.Colors.Danger}))
	}

	root := stack.New(ctx, views, nil)

	return root
}

func LoginSuperSecure(data *CustomData, fail bool) {
	data.LoginFailed = ""
	data.LoginingIn = true
	time.Sleep(2 * time.Second)
	if fail {
		data.LoginingIn = false
		data.LoginFailed = "Login failed! Ouch!"
		return
	}

	// Setting global state here. Could be from DB or something else.
	data.UserID = "1234abc"
}
