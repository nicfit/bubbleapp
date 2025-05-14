package main

import (
	"time"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/context"
	"github.com/alexanderbh/bubbleapp/component/loader"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
)

type appData struct {
	loggingIn   bool
	loginFailed string
	userID      string
}

var AppDataContext = context.Create(&appData{})

func NewLoginRoot(ctx *app.Ctx, _ app.Props) string {
	return context.NewProvider(ctx, AppDataContext, func(ctx *app.Ctx) {
		base(ctx, nil)
	})
}

func base(ctx *app.Ctx, _ app.Props) string {
	appData := context.UseContext(ctx, AppDataContext)

	if appData.userID != "" {
		return NewAuthModel(ctx, authProps{userID: appData.userID})
	}

	if appData.loggingIn {
		return stack.New(ctx, func(ctx *app.Ctx) {
			text.New(ctx, "Please wait...")
			loader.New(ctx, loader.Binary, "Logging in...")
		})
	}

	return stack.New(ctx, func(ctx *app.Ctx) {

		text.New(ctx, "██       ██████   ██████  ██ ███    ██\n██      ██    ██ ██       ██ ████   ██\n██      ██    ██ ██   ███ ██ ██ ██  ██\n██      ██    ██ ██    ██ ██ ██  ██ ██\n███████  ██████   ██████  ██ ██   ████\n\n")
		text.New(ctx, "Log in or fail! Up to you!")

		button.New(ctx, "Log in", func() {
			go LoginSuperSecure(ctx, appData, false)
		}, button.WithVariant(button.Primary))

		button.New(ctx, "Fail log in", func() {
			go LoginSuperSecure(ctx, appData, true)
		}, button.WithVariant(button.Warning))

		button.New(ctx, "Quit App", ctx.Quit, button.WithVariant(button.Danger))

		if appData.loginFailed != "" {
			text.New(ctx, "\n"+appData.loginFailed, text.WithFg(ctx.Styles.Colors.Danger))
		}

	})
}

func LoginSuperSecure(c *app.Ctx, data *appData, fail bool) {
	data.loginFailed = ""
	data.loggingIn = true
	c.Update()
	time.Sleep(2 * time.Second)
	if fail {
		data.loggingIn = false
		data.loginFailed = "Login failed! Ouch!"
		c.Update()
		return
	}

	// Setting global state here. Could be from DB or something else.
	data.userID = "1234abc"
	c.Update()
}
