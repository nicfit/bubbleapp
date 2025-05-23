package main

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/box"
	"github.com/alexanderbh/bubbleapp/component/markdown"
	"github.com/alexanderbh/bubbleapp/component/stack"
)

func tabtab(c *app.Ctx) *app.C {
	return stack.New(c, func(c *app.Ctx) []*app.C {
		return []*app.C{
			box.New(c, func(c *app.Ctx) *app.C {
				return markdown.New(c, "## This is a box with a background color!")
			}),
			box.New(c, func(c *app.Ctx) *app.C {
				return markdown.New(c, `What is Lorem Ipsum?`)
			}),
		}
	})
}
