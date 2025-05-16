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

func NewRoot(ctx *app.Ctx, _ app.Props) app.C {

	stack := stack.New(ctx, func(ctx *app.Ctx) []app.C {
		return []app.C{

			text.New(ctx, "Loaders:"),
			divider.New(ctx),
			loader.New(ctx, loader.Dots, "With text...", loader.WithColor(ctx.Styles.Colors.InfoLight), loader.WithTextColor(ctx.Styles.Colors.Primary)),
			tickfps.NewAtInterval(ctx, 100*time.Microsecond), // Used for debugging tick events.
			stack.New(ctx, func(ctx *app.Ctx) []app.C {
				return []app.C{

					box.New(ctx, func(ctx *app.Ctx) app.C {
						return stack.New(ctx, func(ctx *app.Ctx) []app.C {
							return []app.C{
								loader.NewWithoutText(ctx, loader.Dots),
								loader.NewWithoutText(ctx, loader.Dots2),
								loader.NewWithoutText(ctx, loader.Dots3),
								loader.NewWithoutText(ctx, loader.Dots4),
								loader.NewWithoutText(ctx, loader.Dots5),
								loader.NewWithoutText(ctx, loader.Dots6),
								loader.NewWithoutText(ctx, loader.Dots7),
								loader.NewWithoutText(ctx, loader.Dots8),
								loader.NewWithoutText(ctx, loader.Dots9),
								loader.NewWithoutText(ctx, loader.Dots10),
								loader.NewWithoutText(ctx, loader.Dots11),
								loader.NewWithoutText(ctx, loader.Dots12),
								loader.NewWithoutText(ctx, loader.Dots13),
								loader.NewWithoutText(ctx, loader.Dots14),
								loader.NewWithoutText(ctx, loader.Dots8Bit),
								loader.NewWithoutText(ctx, loader.DotsCircle),
								loader.NewWithoutText(ctx, loader.Sand),
								loader.NewWithoutText(ctx, loader.Line),
							}
						})
					}),

					box.New(ctx, func(ctx *app.Ctx) app.C {
						return stack.New(ctx, func(ctx *app.Ctx) []app.C {
							return []app.C{
								loader.NewWithoutText(ctx, loader.Line2),
								loader.NewWithoutText(ctx, loader.Pipe),
								loader.NewWithoutText(ctx, loader.SimpleDots),
								loader.NewWithoutText(ctx, loader.SimpleDotsScrolling),
								loader.NewWithoutText(ctx, loader.Star),
								loader.NewWithoutText(ctx, loader.Star2),
								loader.NewWithoutText(ctx, loader.Flip),
								loader.NewWithoutText(ctx, loader.Hamburger),
								loader.NewWithoutText(ctx, loader.GrowVertical),
								loader.NewWithoutText(ctx, loader.GrowHorizontal),
								loader.NewWithoutText(ctx, loader.Balloon),
								loader.NewWithoutText(ctx, loader.Balloon2),
								loader.NewWithoutText(ctx, loader.Noise),
								loader.NewWithoutText(ctx, loader.Dqpb),
								loader.NewWithoutText(ctx, loader.Bounce),
								loader.NewWithoutText(ctx, loader.BoxBounce),
								loader.NewWithoutText(ctx, loader.BoxBounce2),
								loader.NewWithoutText(ctx, loader.Triangle),
							}
						})
					}),
					box.New(ctx, func(ctx *app.Ctx) app.C {
						return stack.New(ctx, func(ctx *app.Ctx) []app.C {
							return []app.C{
								loader.NewWithoutText(ctx, loader.Binary),
								loader.NewWithoutText(ctx, loader.Arc),
								loader.NewWithoutText(ctx, loader.Circle),
								loader.NewWithoutText(ctx, loader.SquareCorners),
								loader.NewWithoutText(ctx, loader.CircleQuarters),
								loader.NewWithoutText(ctx, loader.CircleHalves),
								loader.NewWithoutText(ctx, loader.Squish),
								loader.NewWithoutText(ctx, loader.Toggle),
								loader.NewWithoutText(ctx, loader.Toggle2),
								loader.NewWithoutText(ctx, loader.Toggle3),
								loader.NewWithoutText(ctx, loader.Toggle4),
								loader.NewWithoutText(ctx, loader.Toggle5),
								loader.NewWithoutText(ctx, loader.Toggle6),
								loader.NewWithoutText(ctx, loader.Toggle7),
								loader.NewWithoutText(ctx, loader.Toggle8),
								loader.NewWithoutText(ctx, loader.Toggle9),
								loader.NewWithoutText(ctx, loader.Toggle10),
								loader.NewWithoutText(ctx, loader.Toggle11),
							}
						})
					}),
					box.New(ctx, func(ctx *app.Ctx) app.C {
						return stack.New(ctx, func(ctx *app.Ctx) []app.C {
							return []app.C{
								loader.NewWithoutText(ctx, loader.Toggle12),
								loader.NewWithoutText(ctx, loader.Toggle13),
								loader.NewWithoutText(ctx, loader.Arrow),
								loader.NewWithoutText(ctx, loader.Arrow3),
								loader.NewWithoutText(ctx, loader.BouncingBar),
								loader.NewWithoutText(ctx, loader.BouncingBall),
								loader.NewWithoutText(ctx, loader.AestheticSmall),
								loader.NewWithoutText(ctx, loader.Point),
								loader.NewWithoutText(ctx, loader.Layer),
								loader.NewWithoutText(ctx, loader.BetaWave),
								loader.NewWithoutText(ctx, loader.Monkey),
								loader.NewWithoutText(ctx, loader.Hearts),
								loader.NewWithoutText(ctx, loader.Clock),
								loader.NewWithoutText(ctx, loader.Earth),
								loader.NewWithoutText(ctx, loader.Moon),
							}
						})
					}),
					box.New(ctx, func(ctx *app.Ctx) app.C {
						return stack.New(ctx, func(ctx *app.Ctx) []app.C {
							return []app.C{
								loader.NewWithoutText(ctx, loader.Runner),
								loader.NewWithoutText(ctx, loader.Pong),
								loader.NewWithoutText(ctx, loader.Shark),
								loader.NewWithoutText(ctx, loader.Weather),
								loader.NewWithoutText(ctx, loader.Christmas),
								loader.NewWithoutText(ctx, loader.Arrow2),
								loader.NewWithoutText(ctx, loader.Smiley),
								loader.NewWithoutText(ctx, loader.FingerDance),
								loader.NewWithoutText(ctx, loader.FistBump),
								loader.NewWithoutText(ctx, loader.SoccerHeader),
								loader.NewWithoutText(ctx, loader.Mindblown),
								loader.NewWithoutText(ctx, loader.Speaker),
								loader.NewWithoutText(ctx, loader.OrangePulse),
								loader.NewWithoutText(ctx, loader.BluePulse),
								loader.NewWithoutText(ctx, loader.OrangeBluePulse),
								loader.NewWithoutText(ctx, loader.TimeTravel),
								loader.NewWithoutText(ctx, loader.Aesthetic),
								loader.NewWithoutText(ctx, loader.Grenade),
								loader.NewWithoutText(ctx, loader.DwarfFortress),
							}
						})
					}),
				}

			}, stack.WithDirection(app.Horizontal)),

			text.New(ctx, "Press [ctrl-c] to quit.", text.WithFg(ctx.Styles.Colors.Danger)),
		}
	},
	)

	return stack
}

func main() {
	ctx := app.NewCtx()

	bubbleApp := app.New(ctx, NewRoot)
	p := tea.NewProgram(bubbleApp, tea.WithAltScreen(), tea.WithMouseAllMotion())
	bubbleApp.SetTeaProgram(p)
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
