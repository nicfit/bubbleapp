package main

import (
	"errors"
	"time"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/context"
	"github.com/alexanderbh/bubbleapp/component/loader"
	"github.com/alexanderbh/bubbleapp/component/router"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
)

type authData struct {
	userID string
}

type MainApp struct {
	data    authData
	setData func(authData)
}

var AppDataContext = context.Create(MainApp{})

func NewLoginRoot(ctx *app.Ctx, _ app.Props) string {
	data, setData := app.UseState(ctx, authData{})

	mainApp := MainApp{
		data:    data,
		setData: func(ad authData) { setData(ad) },
	}

	return context.NewProvider(ctx, AppDataContext, mainApp, func(ctx *app.Ctx) {
		router.NewRouter(ctx, router.RouterProps{
			Routes: []router.Route{
				{Path: "/", Component: mainRoute},
				{Path: "/login", Component: loginRoute},
			},
		})
	})
}

func mainRoute(ctx *app.Ctx, _ app.Props) string {
	// UseContext returns AppContextValue.
	contextValue := context.UseContext(ctx, AppDataContext)
	appAuthData := contextValue.data // This is *authData
	router := router.UseRouterController(ctx)

	if appAuthData.userID == "" {
		router.ReplaceRoot("/login")
		return ""
	}

	return NewAuthModel(ctx, nil)
}

func loginRoute(ctx *app.Ctx, _ app.Props) string {
	appData := context.UseContext(ctx, AppDataContext)

	loggingIn, setLogginIn := app.UseState(ctx, false)
	loginError, setLoginError := app.UseState(ctx, "")
	router := router.UseRouterController(ctx)

	if loggingIn {
		return stack.New(ctx, func(ctx *app.Ctx) {
			text.New(ctx, "Please wait...")
			loader.New(ctx, loader.Binary, "Logging in...")
		})
	}

	loginFunc := func(fail bool) {
		userID, err := LoginSuperSecure(ctx, fail)
		if err != nil {
			setLogginIn(false)
			setLoginError("Login failed: " + err.Error())
			return
		}

		appData.data.userID = userID
		appData.setData(appData.data)

		router.ReplaceRoot("/")
	}

	return stack.New(ctx, func(ctx *app.Ctx) {

		text.New(ctx, "██       ██████   ██████  ██ ███    ██\n██      ██    ██ ██       ██ ████   ██\n██      ██    ██ ██   ███ ██ ██ ██  ██\n██      ██    ██ ██    ██ ██ ██  ██ ██\n███████  ██████   ██████  ██ ██   ████\n\n")
		text.New(ctx, "Log in or fail! Up to you!")

		button.New(ctx, "Log in", func() {
			setLoginError("")
			setLogginIn(true)
			go loginFunc(false)
		}, button.WithVariant(button.Primary))

		button.New(ctx, "Fail log in", func() {
			setLoginError("")
			setLogginIn(true)
			go loginFunc(true)
		}, button.WithVariant(button.Warning))

		button.New(ctx, "Quit App", ctx.Quit, button.WithVariant(button.Danger))

		if loginError != "" {
			text.New(ctx, "\n"+loginError, text.WithFg(ctx.Styles.Colors.Danger))
		}

	})
}

func LoginSuperSecure(c *app.Ctx, fail bool) (string, error) {
	time.Sleep(2 * time.Second)
	if fail {
		return "", errors.New("yikes")
	}

	return "1234abcd", nil
}
