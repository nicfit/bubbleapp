package main

import (
	"os"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/box"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/grid"
	"github.com/alexanderbh/bubbleapp/component/loader"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
	"github.com/alexanderbh/bubbleapp/style"

	zone "github.com/alexanderbh/bubblezone/v2"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type CustomData struct{}

func NewRoot() model[CustomData] {
	ctx := &app.Context[CustomData]{
		Styles: style.DefaultStyles(),
		Zone:   zone.New(),
	}

	gridView := grid.New(ctx,
		grid.Item[CustomData]{
			Xs: 12,
			Item: box.New(ctx, &box.Options[CustomData]{
				Bg:    ctx.Styles.Colors.PrimaryDark,
				Child: text.New(ctx, "I wish I could center text! Some day...", nil).Base(),
			}).Base(),
		},
		grid.Item[CustomData]{
			Xs:   6,
			Item: box.New(ctx, &box.Options[CustomData]{Bg: ctx.Styles.Colors.InfoLight}).Base(),
		},
		grid.Item[CustomData]{
			Xs: 6,
			Item: stack.New(ctx, &stack.Options[CustomData]{
				Children: []*app.Base[CustomData]{
					text.New(ctx, "Background mess up if this text has foreground style.", nil).Base(),
					text.New(ctx, "Fix the margin to the left here. Not intentional.", nil).Base(),
					button.New(ctx, "BUTTON 1", nil).Base(),
				},
			}).Base(),
		},
		grid.Item[CustomData]{
			Xs:   3,
			Item: button.New(ctx, "BUTTON 2", &button.Options{Variant: button.Danger}).Base(),
		},
		grid.Item[CustomData]{
			Xs: 6,
			Item: box.New(ctx, &box.Options[CustomData]{
				Bg: ctx.Styles.Colors.InfoDark,
				Child: stack.New(ctx, &stack.Options[CustomData]{
					Children: []*app.Base[CustomData]{
						text.New(ctx, "I am in a stack!", nil).Base(),
						loader.New(ctx, loader.Meter, &loader.Options{
							Text:  "Text style messes up bg. Fix!",
							Color: ctx.Styles.Colors.Black,
						}).Base(),
					},
				}).Base(),
			}).Base(),
		},
		grid.Item[CustomData]{
			Xs:   3,
			Item: box.New(ctx, &box.Options[CustomData]{Bg: ctx.Styles.Colors.Success}).Base(),
		},
	)

	base := app.New(ctx, app.AsRoot())
	base.AddChild(gridView.Base())

	return model[CustomData]{
		base: base,
	}
}

type model[T CustomData] struct {
	base *app.Base[T]
}

func (m model[T]) Init() tea.Cmd {
	return m.base.Init()
}

func (m model[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	cmd := m.base.Update(msg)

	return m, cmd

}

func (m model[T]) View() string {
	return m.base.Render()
}

func main() {
	p := tea.NewProgram(NewRoot(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
