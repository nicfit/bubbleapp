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
}

func Stack(c *app.FCContext, props app.Props) string {
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
func New(c *app.FCContext, children app.Children, props ...StackProp) string {
	appliedProps := StackProps{
		Children: children,
	}
	for _, prop := range props {
		prop(&appliedProps)
	}

	return c.Render(Stack, appliedProps)
}
