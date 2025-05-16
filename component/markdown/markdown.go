package markdown

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/charmbracelet/glamour"
)

// TODO:
//   - Add support for styles and custom styles
//   - Fix wordwrapping

type Props struct {
	Text string
	app.Layout
}

func New(ctx *app.Ctx, text string) app.C {
	props := Props{
		Text: text,
		Layout: app.Layout{
			GrowX: true,
			GrowY: true,
		},
	}
	return ctx.Render(Markdown, props)
}

func Markdown(ctx *app.Ctx, props app.Props) string {
	markdownProps, ok := props.(Props)
	if !ok {
		panic("Markdown: props must be of type Props")
	}

	glamourRenderer, setGlamourRenderer := app.UseState[*glamour.TermRenderer](ctx, nil)

	app.UseEffect(ctx, func() {
		var r, _ = glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
		)
		setGlamourRenderer(r)
	}, app.RunOnceDeps)

	if glamourRenderer == nil {
		return ""
	}

	out, _ := glamourRenderer.Render(markdownProps.Text)
	return out
}
