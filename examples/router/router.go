package main

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/box"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/divider"
	"github.com/alexanderbh/bubbleapp/component/router"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
)

func MainRouter(c *app.Ctx) *app.C {
	return router.NewRouter(c, router.RouterProps{
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

func dashboard(c *app.Ctx) *app.C {
	router := router.UseRouterController(c)

	return stack.New(c, func(c *app.Ctx) []*app.C {
		return []*app.C{
			text.New(c, "Welcome to the dashboard! "),
			text.New(c, "Press [ctrl-c] to quit.", text.WithFg(c.Theme.Colors.DangerFg)),

			divider.New(c),

			button.New(c, "Shop", func() {
				router.Push(c, "/shop")
			}),

			button.New(c, "My Account", func() {
				router.Push(c, "/account/overview")
			}),
		}
	})
}

func account(c *app.Ctx) *app.C {
	r := router.UseRouterController(c)

	return stack.New(c, func(c *app.Ctx) []*app.C {
		return []*app.C{
			text.New(c, "My Account"),
			text.New(c, "Press [ctrl-c] to quit.", text.WithFg(c.Theme.Colors.DangerFg)),
			stack.New(c, func(c *app.Ctx) []*app.C {
				return []*app.C{
					button.New(c, "Overview", func() {
						r.Push(c, "/account/overview")
					}),
					button.New(c, "My Orders", func() {
						r.Push(c, "/account/orders")
					}),
					button.New(c, "Settings", func() {
						r.Push(c, "/account/settings")
					}),
				}
			}, stack.WithDirection(app.Horizontal), stack.WithGap(3), stack.WithGrowY(false)),

			divider.New(c),

			router.NewOutlet(c),

			box.NewEmpty(c),

			button.New(c, "Back to Dashboard", func() {
				r.Push(c, "/")
			}),
		}
	})
}

func accountOverview(c *app.Ctx) *app.C {
	return text.New(c, "Account Overview")
}

func accountSettings(c *app.Ctx) *app.C {
	return text.New(c, "Account Settings")
}
func accountOrders(c *app.Ctx) *app.C {
	return text.New(c, "Account Orders")
}

func shop(c *app.Ctx) *app.C {
	router := router.UseRouterController(c)

	return stack.New(c, func(c *app.Ctx) []*app.C {
		return []*app.C{
			text.New(c, "Welcome to the shop!"),
			text.New(c, "Press [ctrl-c] to quit.", text.WithFg(c.Theme.Colors.Danger)),

			divider.New(c),

			button.New(c, "Back to Dashboard", func() {
				router.Push(c, "/")
			}),

			stack.New(c, func(c *app.Ctx) []*app.C {
				return []*app.C{
					text.New(c, "Shop Item 1 - $10", text.WithFg(c.Theme.Colors.Tertiary)),
					text.New(c, "Shop Item 2 - $12", text.WithFg(c.Theme.Colors.Tertiary)),
					text.New(c, "This is just for show", text.WithFg(c.Theme.Colors.Warning)),
				}
			}),
		}
	})
}
