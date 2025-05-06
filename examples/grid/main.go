package main

// import (
// 	"os"

// 	"github.com/alexanderbh/bubbleapp/app"
// 	"github.com/alexanderbh/bubbleapp/component/box"
// 	"github.com/alexanderbh/bubbleapp/component/button"
// 	"github.com/alexanderbh/bubbleapp/component/grid"
// 	"github.com/alexanderbh/bubbleapp/component/loader"
// 	"github.com/alexanderbh/bubbleapp/component/stack"
// 	"github.com/alexanderbh/bubbleapp/component/text"
// 	"github.com/alexanderbh/bubbleapp/style"

// 	zone "github.com/alexanderbh/bubblezone/v2"
// 	tea "github.com/charmbracelet/bubbletea/v2"
// )

// type CustomData struct{}

// func NewRoot() model {
// 	ctx := &app.Context[CustomData]{
// 		Styles: style.DefaultStyles(),
// 		Zone:   zone.New(),
// 	}

// 	gridView := grid.New(ctx, &grid.Options[CustomData]{
// 		Items: []grid.Item[CustomData]{
// 			{Xs: 6, Lg: 3,
// 				Item: box.New(ctx, &box.Options[CustomData]{Bg: ctx.Styles.Colors.DangerDark,
// 					Child: text.New(ctx, "I wish I could center text! Some day...", nil),
// 				}),
// 			},
// 			{Xs: 6, Lg: 3,
// 				Item: box.New(ctx, &box.Options[CustomData]{Bg: ctx.Styles.Colors.Success}),
// 			},
// 			{Xs: 6, Lg: 3,
// 				Item: stack.New(ctx, &stack.Options[CustomData]{
// 					Children: []*app.Base[CustomData]{
// 						text.New(ctx, "Background mess up if this text has foreground style.", nil),
// 						text.New(ctx, "Otherwise pretty nice", nil),
// 						button.New(ctx, "BUTTON 1", &button.Options{Type: button.Compact}),
// 					},
// 				}),
// 			},
// 			{Xs: 6, Lg: 3,
// 				Item: button.New(ctx, "BUTTON 2", &button.Options{Variant: button.Danger, Type: button.Compact}),
// 			},
// 			{Xs: 6,
// 				Item: box.New(ctx, &box.Options[CustomData]{Bg: ctx.Styles.Colors.InfoDark,
// 					Child: stack.New(ctx, &stack.Options[CustomData]{
// 						Children: []*app.Base[CustomData]{
// 							text.New(ctx, "I am in a stack!", nil),
// 							loader.New(ctx, loader.Meter, &loader.Options{Color: ctx.Styles.Colors.DangerDark, Text: "Style is reset here. Fix!"}),
// 						},
// 					}),
// 				}),
// 			},
// 			{Xs: 6,
// 				Item: box.New(ctx, &box.Options[CustomData]{Bg: ctx.Styles.Colors.Warning}),
// 			},
// 		}},
// 		app.AsRoot(),
// 	)

// 	return model{
// 		base: gridView,
// 	}
// }

// type model struct {
// 	base *app.Base[CustomData]
// }

// func (m model) Init() tea.Cmd {
// 	return m.base.Init()
// }

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "ctrl+c", "q":
// 			return m, tea.Quit
// 		}
// 	}
// 	cmd := m.base.Update(msg)

// 	return m, cmd

// }

// func (m model) View() string {
// 	return m.base.Render()
// }

// func main() {
// 	p := tea.NewProgram(NewRoot(), tea.WithAltScreen())
// 	if _, err := p.Run(); err != nil {
// 		os.Exit(1)
// 	}
// }
