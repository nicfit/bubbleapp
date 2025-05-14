package main

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/divider"
	"github.com/alexanderbh/bubbleapp/component/router"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
)

func MainRouter(ctx *app.Ctx, _ app.Props) string {
	return router.NewRouter(ctx, router.RouterProps{
		Routes: []router.Route{
			{Path: "/", Component: dashboard},
			{Path: "/shop", Component: shop},

			{Path: "/account", Component: account, Children: []router.Route{
				{Path: "/overview", Component: accountOverview},
				{Path: "/settings", Component: accountSettings},
				{Path: "/orders", Component: accountOrders},
			}},
		},
	})
}

func dashboard(ctx *app.Ctx, _ app.Props) string {
	router := router.UseRouterController(ctx)

	return stack.New(ctx, func(ctx *app.Ctx) {
		text.New(ctx, "Welcome to the dashboard! ")
		text.New(ctx, "Press [ctrl-c] to quit.", text.WithFg(ctx.Styles.Colors.Danger))

		divider.New(ctx)

		button.New(ctx, "Shop", func() {
			router.Push("/shop")
		})

		button.New(ctx, "My Account", func() {
			router.Push("/account/overview")
		})

	})
}

func account(ctx *app.Ctx, _ app.Props) string {
	r := router.UseRouterController(ctx)

	return stack.New(ctx, func(ctx *app.Ctx) {
		text.New(ctx, "My Account")
		text.New(ctx, "Press [ctrl-c] to quit.", text.WithFg(ctx.Styles.Colors.Danger))
		stack.New(ctx, func(ctx *app.Ctx) {
			button.New(ctx, "Overview", func() {
				r.Push("/account/overview")
			})
			button.New(ctx, "My Orders", func() {
				r.Push("/account/orders")
			})
			button.New(ctx, "Settings", func() {
				r.Push("/account/settings")
			})
		}, stack.WithDirection(app.Horizontal), stack.WithGap(3), stack.WithGrow(false))

		divider.New(ctx)

		router.NewOutlet(ctx)

		//box.NewEmpty(ctx, box.WithBg(ctx.Styles.Colors.Tertiary))

		button.New(ctx, "Back to Dashboard", func() {
			r.Push("/")
		})
	})
}

func accountOverview(ctx *app.Ctx, _ app.Props) string {
	return text.New(ctx, "Account Overview")
}

func accountSettings(ctx *app.Ctx, _ app.Props) string {
	return text.New(ctx, "Account Settings")
}
func accountOrders(ctx *app.Ctx, _ app.Props) string {
	return text.New(ctx, "Account Orders")
}

func shop(ctx *app.Ctx, _ app.Props) string {
	router := router.UseRouterController(ctx)

	return stack.New(ctx, func(ctx *app.Ctx) {
		text.New(ctx, "Welcome to the shop!")
		text.New(ctx, "Press [ctrl-c] to quit.", text.WithFg(ctx.Styles.Colors.Danger))

		divider.New(ctx)

		button.New(ctx, "Back to Dashboard", func() {
			router.Push("/")
		})

		stack.New(ctx, func(ctx *app.Ctx) {
			text.New(ctx, "Shop Item 1 - $10", text.WithFg(ctx.Styles.Colors.Tertiary))
			text.New(ctx, "Shop Item 2 - $12", text.WithFg(ctx.Styles.Colors.Tertiary))
			text.New(ctx, "This is just for show", text.WithFg(ctx.Styles.Colors.Warning))
		})
	})
}
