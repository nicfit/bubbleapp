package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/box"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/divider"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"

	tea "github.com/charmbracelet/bubbletea/v2"
)

type CustomData struct {
	presses int
	log     []string
}

func NewRoot(ctx *app.Context[CustomData]) app.Fc[CustomData] {

	addButton := button.New(ctx, "Button 1", func(ctx *app.Context[CustomData]) {
		ctx.Data.log = append(ctx.Data.log, "["+strconv.Itoa(ctx.Data.presses)+"] "+"Button 1 pressed")
		ctx.Data.presses++
	}, &button.Options{Variant: button.Primary, Type: button.Compact})

	logMessages := box.New(ctx, func(ctx *app.Context[CustomData]) app.Fc[CustomData] {
		return text.NewDynamic(ctx, func(ctx *app.Context[CustomData]) (log string) {
			return strings.Join(ctx.Data.log, "\n")
		}, nil)
	}, nil)

	stack := stack.New(ctx, func(ctx *app.Context[CustomData]) []app.Fc[CustomData] {
		return []app.Fc[CustomData]{
			text.New(ctx, "Tab through the buttons to see focus state!", nil),
			addButton,
			divider.New(ctx),
			logMessages,
			divider.New(ctx),
			button.New(ctx, "Quit App", app.Quit, &button.Options{Variant: button.Danger, Type: button.Compact}),
		}
	}, nil)

	return stack
}

func main() {
	ctx := app.NewContext(&CustomData{})

	app := app.NewApp(ctx, NewRoot)
	p := tea.NewProgram(app, tea.WithAltScreen(), tea.WithMouseAllMotion())
	app.SetTeaProgram(p)
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
