package main

import (
	"os"
	"time"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/box"
	"github.com/alexanderbh/bubbleapp/component/divider"
	"github.com/alexanderbh/bubbleapp/component/loader"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
	"github.com/alexanderbh/bubbleapp/component/tickfps"

	tea "github.com/charmbracelet/bubbletea/v2"
)

type CustomData struct{}

func NewRoot(ctx *app.Context[CustomData]) app.Fc[CustomData] {

	stack := stack.New(ctx, func(ctx *app.Context[CustomData]) []app.Fc[CustomData] {
		return []app.Fc[CustomData]{
			text.New(ctx, "Loaders:", nil),
			divider.New(ctx),
			loader.New(ctx, loader.Dots, "With text...", &loader.Options{Color: ctx.Styles.Colors.InfoLight, TextColor: ctx.Styles.Colors.Primary}),
			tickfps.New(ctx, time.Second), // Used for debugging tick events.
			stack.New(ctx, func(ctx *app.Context[CustomData]) []app.Fc[CustomData] {
				return []app.Fc[CustomData]{
					box.New(ctx, func(ctx *app.Context[CustomData]) app.Fc[CustomData] {
						return stack.New(ctx, func(ctx *app.Context[CustomData]) []app.Fc[CustomData] {
							return []app.Fc[CustomData]{
								loader.NewWithoutText(ctx, loader.Dots, nil),
								loader.NewWithoutText(ctx, loader.Dots2, nil),
								loader.NewWithoutText(ctx, loader.Dots3, nil),
								loader.NewWithoutText(ctx, loader.Dots4, nil),
								loader.NewWithoutText(ctx, loader.Dots5, nil),
								loader.NewWithoutText(ctx, loader.Dots6, nil),
								loader.NewWithoutText(ctx, loader.Dots7, nil),
								loader.NewWithoutText(ctx, loader.Dots8, nil),
								loader.NewWithoutText(ctx, loader.Dots9, nil),
								loader.NewWithoutText(ctx, loader.Dots10, nil),
								loader.NewWithoutText(ctx, loader.Dots11, nil),
								loader.NewWithoutText(ctx, loader.Dots12, nil),
								loader.NewWithoutText(ctx, loader.Dots13, nil),
								loader.NewWithoutText(ctx, loader.Dots14, nil),
								loader.NewWithoutText(ctx, loader.Dots8Bit, nil),
								loader.NewWithoutText(ctx, loader.DotsCircle, nil),
								loader.NewWithoutText(ctx, loader.Sand, nil),
								loader.NewWithoutText(ctx, loader.Line, nil),
							}
						}, nil)
					}, nil),
					box.New(ctx, func(ctx *app.Context[CustomData]) app.Fc[CustomData] {
						return stack.New(ctx, func(ctx *app.Context[CustomData]) []app.Fc[CustomData] {
							return []app.Fc[CustomData]{
								loader.NewWithoutText(ctx, loader.Line2, nil),
								loader.NewWithoutText(ctx, loader.Pipe, nil),
								loader.NewWithoutText(ctx, loader.SimpleDots, nil),
								loader.NewWithoutText(ctx, loader.SimpleDotsScrolling, nil),
								loader.NewWithoutText(ctx, loader.Star, nil),
								loader.NewWithoutText(ctx, loader.Star2, nil),
								loader.NewWithoutText(ctx, loader.Flip, nil),
								loader.NewWithoutText(ctx, loader.Hamburger, nil),
								loader.NewWithoutText(ctx, loader.GrowVertical, nil),
								loader.NewWithoutText(ctx, loader.GrowHorizontal, nil),
								loader.NewWithoutText(ctx, loader.Balloon, nil),
								loader.NewWithoutText(ctx, loader.Balloon2, nil),
								loader.NewWithoutText(ctx, loader.Noise, nil),
								loader.NewWithoutText(ctx, loader.Dqpb, nil),
								loader.NewWithoutText(ctx, loader.Bounce, nil),
								loader.NewWithoutText(ctx, loader.BoxBounce, nil),
								loader.NewWithoutText(ctx, loader.BoxBounce2, nil),
								loader.NewWithoutText(ctx, loader.Triangle, nil),
							}
						}, nil)
					}, nil),
					box.New(ctx, func(ctx *app.Context[CustomData]) app.Fc[CustomData] {
						return stack.New(ctx, func(ctx *app.Context[CustomData]) []app.Fc[CustomData] {
							return []app.Fc[CustomData]{
								loader.NewWithoutText(ctx, loader.Binary, nil),
								loader.NewWithoutText(ctx, loader.Arc, nil),
								loader.NewWithoutText(ctx, loader.Circle, nil),
								loader.NewWithoutText(ctx, loader.SquareCorners, nil),
								loader.NewWithoutText(ctx, loader.CircleQuarters, nil),
								loader.NewWithoutText(ctx, loader.CircleHalves, nil),
								loader.NewWithoutText(ctx, loader.Squish, nil),
								loader.NewWithoutText(ctx, loader.Toggle, nil),
								loader.NewWithoutText(ctx, loader.Toggle2, nil),
								loader.NewWithoutText(ctx, loader.Toggle3, nil),
								loader.NewWithoutText(ctx, loader.Toggle4, nil),
								loader.NewWithoutText(ctx, loader.Toggle5, nil),
								loader.NewWithoutText(ctx, loader.Toggle6, nil),
								loader.NewWithoutText(ctx, loader.Toggle7, nil),
								loader.NewWithoutText(ctx, loader.Toggle8, nil),
								loader.NewWithoutText(ctx, loader.Toggle9, nil),
								loader.NewWithoutText(ctx, loader.Toggle10, nil),
								loader.NewWithoutText(ctx, loader.Toggle11, nil),
							}
						}, nil)
					}, nil),
					box.New(ctx, func(ctx *app.Context[CustomData]) app.Fc[CustomData] {
						return stack.New(ctx, func(ctx *app.Context[CustomData]) []app.Fc[CustomData] {
							return []app.Fc[CustomData]{
								loader.NewWithoutText(ctx, loader.Toggle12, nil),
								loader.NewWithoutText(ctx, loader.Toggle13, nil),
								loader.NewWithoutText(ctx, loader.Arrow, nil),
								loader.NewWithoutText(ctx, loader.Arrow3, nil),
								loader.NewWithoutText(ctx, loader.BouncingBar, nil),
								loader.NewWithoutText(ctx, loader.BouncingBall, nil),
								loader.NewWithoutText(ctx, loader.AestheticSmall, nil),
								loader.NewWithoutText(ctx, loader.Point, nil),
								loader.NewWithoutText(ctx, loader.Layer, nil),
								loader.NewWithoutText(ctx, loader.BetaWave, nil),
								loader.NewWithoutText(ctx, loader.Monkey, nil),
								loader.NewWithoutText(ctx, loader.Hearts, nil),
								loader.NewWithoutText(ctx, loader.Clock, nil),
								loader.NewWithoutText(ctx, loader.Earth, nil),
								loader.NewWithoutText(ctx, loader.Moon, nil),
							}
						}, nil)
					}, nil),
					box.New(ctx, func(ctx *app.Context[CustomData]) app.Fc[CustomData] {
						return stack.New(ctx, func(ctx *app.Context[CustomData]) []app.Fc[CustomData] {
							return []app.Fc[CustomData]{
								loader.NewWithoutText(ctx, loader.Runner, nil),
								loader.NewWithoutText(ctx, loader.Pong, nil),
								loader.NewWithoutText(ctx, loader.Shark, nil),
								loader.NewWithoutText(ctx, loader.Weather, nil),
								loader.NewWithoutText(ctx, loader.Christmas, nil),
								loader.NewWithoutText(ctx, loader.Arrow2, nil),
								loader.NewWithoutText(ctx, loader.Smiley, nil),
								loader.NewWithoutText(ctx, loader.FingerDance, nil),
								loader.NewWithoutText(ctx, loader.FistBump, nil),
								loader.NewWithoutText(ctx, loader.SoccerHeader, nil),
								loader.NewWithoutText(ctx, loader.Mindblown, nil),
								loader.NewWithoutText(ctx, loader.Speaker, nil),
								loader.NewWithoutText(ctx, loader.OrangePulse, nil),
								loader.NewWithoutText(ctx, loader.BluePulse, nil),
								loader.NewWithoutText(ctx, loader.OrangeBluePulse, nil),
								loader.NewWithoutText(ctx, loader.TimeTravel, nil),
								loader.NewWithoutText(ctx, loader.Aesthetic, nil),
								loader.NewWithoutText(ctx, loader.Grenade, nil),
								loader.NewWithoutText(ctx, loader.DwarfFortress, nil),
							}
						}, nil)
					}, nil),
				}
			}, &stack.Options{Horizontal: true}),
			text.New(ctx, "Press [ctrl-c] to quit.", &text.Options{Foreground: ctx.Styles.Colors.Danger}),
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
