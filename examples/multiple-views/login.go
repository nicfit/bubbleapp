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

func NewLoginRoot(c *app.Ctx, _ app.Props) app.C {
	data, setData := app.UseState(c, authData{})

	mainApp := MainApp{
		data:    data,
		setData: func(ad authData) { setData(ad) },
	}

	return context.NewProvider(c, AppDataContext, mainApp, func(c *app.Ctx) app.C {
		return router.NewRouter(c, router.RouterProps{
			Routes: []router.Route{
				{Path: "/", Component: mainRoute},
				{Path: "/login", Component: loginRoute},
			},
		})
	})

}

func mainRoute(c *app.Ctx, _ app.Props) app.C {
	// UseContext returns AppContextValue.
	contextValue := context.UseContext(c, AppDataContext)
	appAuthData := contextValue.data // This is *authData
	router := router.UseRouterController(c)

	if appAuthData.userID == "" {
		router.ReplaceRoot(c, "/login")
		return text.New(c, "No user logged in")
	}

	return NewAuthModel(c, nil)
}

func loginRoute(c *app.Ctx, _ app.Props) app.C {
	appData := context.UseContext(c, AppDataContext)

	loggingIn, setLogginIn := app.UseState(c, false)
	loginError, setLoginError := app.UseState(c, "")
	router := router.UseRouterController(c)

	if loggingIn {
		return stack.New(c, func(c *app.Ctx) []app.C {
			return []app.C{
				text.New(c, "Please wait..."),
				loader.New(c, loader.Binary, "Logging in..."),
			}
		})
	}

	loginFunc := func(fail bool) {
		userID, err := LoginSuperSecure(c, fail)
		if err != nil {
			setLogginIn(false)
			setLoginError("Login failed: " + err.Error())
			return
		}

		appData.data.userID = userID
		appData.setData(appData.data)

		router.ReplaceRoot(c, "/")
	}

	return stack.New(c, func(c *app.Ctx) []app.C {
		views := []app.C{
			text.New(c, "██       ██████   ██████  ██ ███    ██\n██      ██    ██ ██       ██ ████   ██\n██      ██    ██ ██   ███ ██ ██ ██  ██\n██      ██    ██ ██    ██ ██ ██  ██ ██\n███████  ██████   ██████  ██ ██   ████\n\n"),
			text.New(c, "Log in or fail! Up to you!"),

			button.New(c, "Log in", func() {
				setLoginError("")
				setLogginIn(true)
				go loginFunc(false)
			}, button.WithVariant(button.Primary)),

			button.New(c, "Fail log in", func() {
				setLoginError("")
				setLogginIn(true)
				go loginFunc(true)
			}, button.WithVariant(button.Warning)),

			button.New(c, "Quit App", c.Quit, button.WithVariant(button.Danger)),
		}
		if loginError != "" {
			views = append(views, text.New(c, "\n"+loginError, text.WithFg(c.Styles.Colors.Danger)))
		}

		return views
	})
}

func LoginSuperSecure(c *app.Ctx, fail bool) (string, error) {
	time.Sleep(2 * time.Second)
	if fail {
		return "", errors.New("yikes")
	}

	return "1234abcd", nil
}
