package main

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/box"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
)

func NewScrolling(c *app.Ctx, _ app.Props) string {
	return c.Render(scrolling, nil)
}

func scrolling(c *app.Ctx, _ app.Props) string {
	return stack.New(c, func(c *app.Ctx) {
		box.New(c, func(c *app.Ctx) {
			text.New(c, "This\nis\na\nbox\nwith\na\nbackground\ncolor!")
			//markdown.New(c, "## This is a box with a background color!")
		}, box.WithBg(c.Styles.Colors.HighlightBackground))
		box.New(c, func(c *app.Ctx) {
			//markdown.New(c, `What is Lorem Ipsum?`)
		}, box.WithBg(c.Styles.Colors.Warning))
	})
}
