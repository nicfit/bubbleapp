package loader

import (
	"image/color"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/charmbracelet/lipgloss/v2"
)

type Props struct {
	Text                string
	Spinner             Spinner
	TextColor           color.Color
	TextBackgroundColor color.Color
	Color               color.Color
}

type prop func(*Props)

func Loader(c *app.Ctx, props app.Props) string {
	loaderProps, ok := props.(Props)
	if !ok {
		panic("Loader: props must be of type loader.Props")
	}

	frame, setFrame := app.UseState(c, 0)

	app.UseTick(c, loaderProps.Spinner.Interval, func() {
		setFrame(func(prev int) int { return (prev + 1) % len(loaderProps.Spinner.Frames) })
	})

	styleText := lipgloss.NewStyle().Padding(0, 0, 0, 1)
	styleSpinner := lipgloss.NewStyle()

	if loaderProps.TextColor != nil {
		styleText = styleText.Foreground(loaderProps.TextColor)
	}
	if loaderProps.TextBackgroundColor != nil {
		styleText = styleText.Background(loaderProps.TextBackgroundColor)
	}
	if loaderProps.Color != nil {
		styleSpinner = styleSpinner.Foreground(loaderProps.Color)
	}

	return styleSpinner.Render(loaderProps.Spinner.Frames[frame]) + styleText.Render(loaderProps.Text)
}

func NewWithoutText(c *app.Ctx, variant Spinner, prop ...prop) app.C {
	return New(c, variant, "", prop...)
}

func New(c *app.Ctx, variant Spinner, text string, prop ...prop) app.C {
	p := Props{
		Color:   c.Styles.Colors.Primary,
		Spinner: variant,
		Text:    text,
	}
	for _, opt := range prop {
		if opt != nil {
			opt(&p)
		}
	}
	return c.Render(Loader, p)
}

func WithColor(color color.Color) prop {
	return func(p *Props) {
		p.Color = color
	}
}
func WithTextColor(color color.Color) prop {
	return func(p *Props) {
		p.TextColor = color
	}
}
func WithTextBackgroundColor(color color.Color) prop {
	return func(p *Props) {
		p.TextBackgroundColor = color
	}
}
