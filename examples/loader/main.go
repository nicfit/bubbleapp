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

func NewRoot(c *app.Ctx, _ app.Props) app.C {

	stack := stack.New(c, func(c *app.Ctx) []app.C {
		return []app.C{

			text.New(c, "Loaders:"),
			divider.New(c),
			loader.New(c, loader.Dots, "With text...", loader.WithColor(c.Styles.Colors.InfoLight), loader.WithTextColor(c.Styles.Colors.Primary)),
			tickfps.NewAtInterval(c, 100*time.Microsecond), // Used for debugging tick events.
			stack.New(c, func(c *app.Ctx) []app.C {
				return []app.C{

					box.New(c, func(c *app.Ctx) app.C {
						return stack.New(c, func(c *app.Ctx) []app.C {
							return []app.C{
								loader.NewWithoutText(c, loader.Dots),
								loader.NewWithoutText(c, loader.Dots2),
								loader.NewWithoutText(c, loader.Dots3),
								loader.NewWithoutText(c, loader.Dots4),
								loader.NewWithoutText(c, loader.Dots5),
								loader.NewWithoutText(c, loader.Dots6),
								loader.NewWithoutText(c, loader.Dots7),
								loader.NewWithoutText(c, loader.Dots8),
								loader.NewWithoutText(c, loader.Dots9),
								loader.NewWithoutText(c, loader.Dots10),
								loader.NewWithoutText(c, loader.Dots11),
								loader.NewWithoutText(c, loader.Dots12),
								loader.NewWithoutText(c, loader.Dots13),
								loader.NewWithoutText(c, loader.Dots14),
								loader.NewWithoutText(c, loader.Dots8Bit),
								loader.NewWithoutText(c, loader.DotsCircle),
								loader.NewWithoutText(c, loader.Sand),
								loader.NewWithoutText(c, loader.Line),
							}
						})
					}),

					box.New(c, func(c *app.Ctx) app.C {
						return stack.New(c, func(c *app.Ctx) []app.C {
							return []app.C{
								loader.NewWithoutText(c, loader.Line2),
								loader.NewWithoutText(c, loader.Pipe),
								loader.NewWithoutText(c, loader.SimpleDots),
								loader.NewWithoutText(c, loader.SimpleDotsScrolling),
								loader.NewWithoutText(c, loader.Star),
								loader.NewWithoutText(c, loader.Star2),
								loader.NewWithoutText(c, loader.Flip),
								loader.NewWithoutText(c, loader.Hamburger),
								loader.NewWithoutText(c, loader.GrowVertical),
								loader.NewWithoutText(c, loader.GrowHorizontal),
								loader.NewWithoutText(c, loader.Balloon),
								loader.NewWithoutText(c, loader.Balloon2),
								loader.NewWithoutText(c, loader.Noise),
								loader.NewWithoutText(c, loader.Dqpb),
								loader.NewWithoutText(c, loader.Bounce),
								loader.NewWithoutText(c, loader.BoxBounce),
								loader.NewWithoutText(c, loader.BoxBounce2),
								loader.NewWithoutText(c, loader.Triangle),
							}
						})
					}),
					box.New(c, func(c *app.Ctx) app.C {
						return stack.New(c, func(c *app.Ctx) []app.C {
							return []app.C{
								loader.NewWithoutText(c, loader.Binary),
								loader.NewWithoutText(c, loader.Arc),
								loader.NewWithoutText(c, loader.Circle),
								loader.NewWithoutText(c, loader.SquareCorners),
								loader.NewWithoutText(c, loader.CircleQuarters),
								loader.NewWithoutText(c, loader.CircleHalves),
								loader.NewWithoutText(c, loader.Squish),
								loader.NewWithoutText(c, loader.Toggle),
								loader.NewWithoutText(c, loader.Toggle2),
								loader.NewWithoutText(c, loader.Toggle3),
								loader.NewWithoutText(c, loader.Toggle4),
								loader.NewWithoutText(c, loader.Toggle5),
								loader.NewWithoutText(c, loader.Toggle6),
								loader.NewWithoutText(c, loader.Toggle7),
								loader.NewWithoutText(c, loader.Toggle8),
								loader.NewWithoutText(c, loader.Toggle9),
								loader.NewWithoutText(c, loader.Toggle10),
								loader.NewWithoutText(c, loader.Toggle11),
							}
						})
					}),
					box.New(c, func(c *app.Ctx) app.C {
						return stack.New(c, func(c *app.Ctx) []app.C {
							return []app.C{
								loader.NewWithoutText(c, loader.Toggle12),
								loader.NewWithoutText(c, loader.Toggle13),
								loader.NewWithoutText(c, loader.Arrow),
								loader.NewWithoutText(c, loader.Arrow3),
								loader.NewWithoutText(c, loader.BouncingBar),
								loader.NewWithoutText(c, loader.BouncingBall),
								loader.NewWithoutText(c, loader.AestheticSmall),
								loader.NewWithoutText(c, loader.Point),
								loader.NewWithoutText(c, loader.Layer),
								loader.NewWithoutText(c, loader.BetaWave),
								loader.NewWithoutText(c, loader.Monkey),
								loader.NewWithoutText(c, loader.Hearts),
								loader.NewWithoutText(c, loader.Clock),
								loader.NewWithoutText(c, loader.Earth),
								loader.NewWithoutText(c, loader.Moon),
							}
						})
					}),
					box.New(c, func(c *app.Ctx) app.C {
						return stack.New(c, func(c *app.Ctx) []app.C {
							return []app.C{
								loader.NewWithoutText(c, loader.Runner),
								loader.NewWithoutText(c, loader.Pong),
								loader.NewWithoutText(c, loader.Shark),
								loader.NewWithoutText(c, loader.Weather),
								loader.NewWithoutText(c, loader.Christmas),
								loader.NewWithoutText(c, loader.Arrow2),
								loader.NewWithoutText(c, loader.Smiley),
								loader.NewWithoutText(c, loader.FingerDance),
								loader.NewWithoutText(c, loader.FistBump),
								loader.NewWithoutText(c, loader.SoccerHeader),
								loader.NewWithoutText(c, loader.Mindblown),
								loader.NewWithoutText(c, loader.Speaker),
								loader.NewWithoutText(c, loader.OrangePulse),
								loader.NewWithoutText(c, loader.BluePulse),
								loader.NewWithoutText(c, loader.OrangeBluePulse),
								loader.NewWithoutText(c, loader.TimeTravel),
								loader.NewWithoutText(c, loader.Aesthetic),
								loader.NewWithoutText(c, loader.Grenade),
								loader.NewWithoutText(c, loader.DwarfFortress),
							}
						})
					}),
				}

			}, stack.WithDirection(app.Horizontal)),

			text.New(c, "Press [ctrl-c] to quit.", text.WithFg(c.Styles.Colors.Danger)),
		}
	},
	)

	return stack
}

func main() {
	c := app.NewCtx()

	bubbleApp := app.New(c, NewRoot)
	p := tea.NewProgram(bubbleApp, tea.WithAltScreen(), tea.WithMouseAllMotion())
	bubbleApp.SetTeaProgram(p)
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
