package main

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
	"github.com/alexanderbh/bubbleapp/component/button"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

// Define a root component that demonstrates conditional rendering of children
func Root(c *app.Ctx, props app.Props) string {
	// Use state to track whether details should be shown
	showDetails, setShowDetails := app.UseState(c, false)

	// Define a children function with conditional logic
	children := func(c *app.Ctx) []app.C {
		// Always show the title and toggle button
		result := []app.C{
			c.Render(text.Text, text.Props{
				Text: "Conditional Children Example",
				Bold: true,
			}),
			c.Render(button.Button, button.Props{
				Text: showDetails ? "Hide Details" : "Show Details",
				OnAction: func() {
					setShowDetails(!showDetails)
				},
			}),
		}
		
		// Conditionally add details when showDetails is true
		if showDetails {
			result = append(result, 
				c.Render(text.Text, text.Props{
					Text: "These are the details that can be toggled on and off.",
					Foreground: lipgloss.Color("5"),
				}),
			)
		}
		
		return result
	}

	// Render the stack with conditional children
	return stack.New(c, children, 
		stack.WithDirection(app.Vertical),
		stack.WithGap(1),
	).String()
}

func main() {
	// Create new app context
	ctx := app.NewCtx()

	// Create app with Root component
	bubbletea := app.New(ctx, Root)

	// Create tea program
	teaProgram := tea.NewProgram(bubbletea)

	bubbletea.SetTeaProgram(teaProgram)

	// Run the program
	_, err := teaProgram.Run()
	if err != nil {
		panic(err)
	}
}
