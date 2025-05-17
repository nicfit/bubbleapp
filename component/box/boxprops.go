package box

import (
	"image/color"

	"github.com/charmbracelet/lipgloss/v2"
)

func WithKey(key string) BoxProp {
	return func(props *BoxProps) {
		props.Key = key
	}
}

// WithBg sets the background color for the box.
func WithBg(bg color.Color) BoxProp {
	return func(props *BoxProps) {
		props.Bg = bg
	}
}

// WithDisableFollow disables the viewport's auto-scrolling to the bottom on content change.
func WithDisableFollow(disable bool) BoxProp {
	return func(props *BoxProps) {
		props.DisableFollow = disable
	}
}

func WithGrow(grow bool) BoxProp {
	return func(props *BoxProps) {
		props.Layout.GrowX = grow
		props.Layout.GrowY = grow
	}
}

func WithGrowX(grow bool) BoxProp {
	return func(props *BoxProps) {
		props.Layout.GrowX = grow
	}
}

func WithGrowY(grow bool) BoxProp {
	return func(props *BoxProps) {
		props.Layout.GrowY = grow
	}
}

func WithWidth(width int) BoxProp {
	return func(props *BoxProps) {
		props.Layout.Width = width
		props.Layout.GrowX = false
	}
}

func WithHeight(height int) BoxProp {
	return func(props *BoxProps) {
		props.Layout.Height = height
		props.Layout.GrowY = false
	}
}

func WithBorder(border lipgloss.Border) BoxProp {
	return func(props *BoxProps) {
		props.Border.Border = border
		props.Border.BorderBottom = true
		props.Border.BorderLeft = true
		props.Border.BorderRight = true
		props.Border.BorderTop = true
	}
}

func WithBorderTop(border lipgloss.Border) BoxProp {
	return func(props *BoxProps) {
		props.Border.Border = border
		props.Border.BorderTop = true
	}
}

func WithBorderBottom(border lipgloss.Border) BoxProp {
	return func(props *BoxProps) {
		props.Border.Border = border
		props.Border.BorderBottom = true
	}
}
func WithBorderLeft(border lipgloss.Border) BoxProp {
	return func(props *BoxProps) {
		props.Border.Border = border
		props.Border.BorderLeft = true
	}
}
func WithBorderRight(border lipgloss.Border) BoxProp {
	return func(props *BoxProps) {
		props.Border.Border = border
		props.Border.BorderRight = true
	}
}

func WithBorderColor(color color.Color) BoxProp {
	return func(props *BoxProps) {
		props.Border.Color = color
	}
}
