package stack

import (
	"strings"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/charmbracelet/lipgloss"
)

type StackProps struct {
	Direction app.LayoutDirection
	Gap       int
	Children  app.Children
	// app.Layout
}

func Stack(c *app.Ctx, props app.Props) string {
	stackProps, _ := props.(StackProps)

	children := app.UseChildren(c, stackProps.Children)

	var processedChildren []string
	if stackProps.Direction == app.Horizontal {
		gap := strings.Repeat(" ", stackProps.Gap)
		for i, child := range children {
			if child != "" {
				processedChildren = append(processedChildren, child)
				// Add gap after each child except the last
				if stackProps.Gap > 0 && i < len(children)-1 {
					processedChildren = append(processedChildren, gap)
				}
			}
		}
	} else {
		for i, child := range children {
			if child != "" {
				processedChildren = append(processedChildren, child)
				// Add vertical gap after each child except the last
				if stackProps.Gap > 0 && i < len(children)-1 {
					processedChildren = append(processedChildren, strings.Repeat(" ", stackProps.Gap))
				}
			}
		}
	}

	if stackProps.Direction == app.Horizontal {
		return lipgloss.JoinHorizontal(lipgloss.Top, processedChildren...)
	}
	return lipgloss.JoinVertical(lipgloss.Left, processedChildren...)

}

func New(c *app.Ctx, children app.Children, props ...StackProp) string {
	appliedProps := StackProps{
		Children: children,
		// Layout: app.Layout{
		// 	GrowX: true,
		// 	GrowY: true,
		// },
	}
	for _, prop := range props {
		prop(&appliedProps)
	}

	return c.Render(Stack, appliedProps)
}

type StackProp func(*StackProps)

func WithDirection(direction app.LayoutDirection) StackProp {
	return func(props *StackProps) {
		props.Direction = direction
	}
}
func WithGap(gap int) StackProp {
	return func(props *StackProps) {
		props.Gap = gap
	}
}

// func WithGrowX(grow bool) StackProp {
// 	return func(props *StackProps) {
// 		props.GrowX = grow
// 	}
// }
// func WithGrowY(grow bool) StackProp {
// 	return func(props *StackProps) {
// 		props.GrowY = grow
// 	}
// }
// func WithGrow(grow bool) StackProp {
// 	return func(props *StackProps) {
// 		props.GrowX = grow
// 		props.GrowY = grow
// 	}
// }
