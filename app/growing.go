package app

import (
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

func CalculateHeights(base *Base, msg tea.WindowSizeMsg) tea.Cmd {
	var (
		cmds []tea.Cmd
	)
	growerCount := 0
	for _, child := range base.GetChildren() {
		if child.Base().Opts.GrowY {
			growerCount++
		}
	}

	nonGrowerHeight := 0
	for _, child := range base.GetChildren() {
		if !child.Base().Opts.GrowY {
			childHeight := lipgloss.Height(child.View())
			nonGrowerHeight += childHeight

			newChild, cmd := child.Update(tea.WindowSizeMsg{
				Width:  msg.Width,
				Height: childHeight,
			})
			newChildTyped := newChild.(UIModel)
			base.ReplaceChild(child.Base().ID, newChildTyped)
			cmds = append(cmds, cmd)
		}
	}

	if growerCount > 0 {
		availableHeight := msg.Height - nonGrowerHeight
		growerHeight := availableHeight / growerCount
		remainder := availableHeight % growerCount
		for _, child := range base.GetChildren() {
			if child.Base().Opts.GrowY {
				currentGrowerHeight := growerHeight
				if remainder > 0 {
					currentGrowerHeight++
					remainder--
				}
				newChild, cmd := child.Update(tea.WindowSizeMsg{
					Width:  msg.Width,
					Height: currentGrowerHeight,
				})
				newChildTyped := newChild.(UIModel)
				base.ReplaceChild(child.Base().ID, newChildTyped)
				cmds = append(cmds, cmd)
			}
		}
	}

	return tea.Batch(cmds...)
}
