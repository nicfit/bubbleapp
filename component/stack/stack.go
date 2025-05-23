package stack

import (
	"strings"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/charmbracelet/lipgloss/v2"
)

type StackProps struct {
	FCs app.FCs
	app.Layout
}

func Stack(c *app.Ctx, props app.Props) string {
	stackProps, _ := props.(StackProps)

	fcs := app.UseFCs(c, stackProps.FCs)

	var processedFCs []string
	if stackProps.Layout.Direction == app.Horizontal {
		gap := strings.Repeat(" ", stackProps.Layout.GapX)
		for i, fc := range fcs {
			if fc != "" {
				processedFCs = append(processedFCs, fc)
				// Add gap after each child except the last
				if stackProps.GapX > 0 && i < len(fcs)-1 {
					processedFCs = append(processedFCs, gap)
				}
			}
		}
	} else {
		for i, fc := range fcs {
			if fc != "" {
				processedFCs = append(processedFCs, fc)
				// Add vertical gap after each child except the last
				if stackProps.GapY > 0 && i < len(fcs)-1 {
					processedFCs = append(processedFCs, " ")
				}
			}
		}
	}

	var result string
	if stackProps.Layout.Direction == app.Horizontal {
		result = lipgloss.JoinHorizontal(lipgloss.Top, processedFCs...)
	} else {
		result = lipgloss.JoinVertical(lipgloss.Left, processedFCs...)
	}

	return result
}

func New(c *app.Ctx, fcs app.FCs, props ...StackProp) *app.C {
	appliedProps := StackProps{
		FCs: fcs,
		Layout: app.Layout{
			Direction: app.Vertical,
			GrowX:     true,
			GrowY:     true,
		},
	}
	for _, prop := range props {
		if prop != nil {
			prop(&appliedProps)
		}
	}

	return c.Render(Stack, appliedProps)
}

type StackProp func(*StackProps)

func WithDirection(direction app.LayoutDirection) StackProp {
	return func(props *StackProps) {
		props.Layout.Direction = direction
	}
}
func WithGap(gap int) StackProp {
	return func(props *StackProps) {
		props.GapX = gap
		props.GapY = gap
	}
}

func WithGrowX(grow bool) StackProp {
	return func(props *StackProps) {
		props.GrowX = grow
	}
}
func WithGrowY(grow bool) StackProp {
	return func(props *StackProps) {
		props.GrowY = grow
	}
}
func WithGrow(grow bool) StackProp {
	return func(props *StackProps) {
		props.GrowX = grow
		props.GrowY = grow
	}
}
