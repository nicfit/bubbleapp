package stack

import (
	"github.com/alexanderbh/bubbleapp/app"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type Options[T any] struct {
	Horizontal bool // Implemented in the future
	Children   []*app.Base[T]
}

type model[T any] struct {
	base  *app.Base[T]
	opts  *Options[T]
	style lipgloss.Style
}

func New[T any](ctx *app.Context[T], options *Options[T]) *app.Base[T] {
	if options == nil {
		options = &Options[T]{}
	}
	base := app.New(ctx, app.WithGrow(true))

	if options.Children != nil {
		base.AddChildren(options.Children...)
	}

	return model[T]{
		base:  base,
		opts:  options,
		style: lipgloss.NewStyle(),
	}.Base()
}

func (m model[T]) Init() tea.Cmd {
	return m.base.Init()
}

func (m model[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.base.Height = msg.Height
		m.base.Width = msg.Width
		m.style = m.style.Width(msg.Width).Height(msg.Height)
		var (
			cmds []tea.Cmd
		)
		growerCount := 0
		for _, child := range m.base.GetChildren() {
			if child.Opts.GrowY {
				growerCount++
			}
		}

		nonGrowerHeight := 0
		for _, child := range m.base.GetChildren() {
			if !child.Opts.GrowY {
				childHeight := lipgloss.Height(child.Model.View())
				nonGrowerHeight += childHeight

				newChild, cmd := child.Model.Update(tea.WindowSizeMsg{
					Width:  msg.Width,
					Height: childHeight,
				})
				newChildTyped := newChild.(app.UIModel[T])
				newChildTyped.Base().Model = newChildTyped // TODO: This seems fragile. Is it needed? Can it be moved to the ReplaceChild?
				m.base.ReplaceChild(child.ID, newChildTyped.Base())
				cmds = append(cmds, cmd)
			}
		}

		if growerCount > 0 {
			availableHeight := msg.Height - nonGrowerHeight
			growerHeight := availableHeight / growerCount
			remainder := availableHeight % growerCount
			for _, child := range m.base.GetChildren() {
				if child.Opts.GrowY {
					currentGrowerHeight := growerHeight
					if remainder > 0 {
						currentGrowerHeight++
						remainder--
					}
					newChild, cmd := child.Model.Update(tea.WindowSizeMsg{
						Width:  msg.Width,
						Height: currentGrowerHeight,
					})
					newChildTyped := newChild.(app.UIModel[T])
					newChildTyped.Base().Model = newChildTyped // TODO: This seems fragile. Is it needed? Can it be moved to the ReplaceChild?
					m.base.ReplaceChild(child.ID, newChildTyped.Base())
					cmds = append(cmds, cmd)
				}
			}
		}

		return m, tea.Batch(cmds...)
	}

	cmd = m.base.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model[T]) View() string {
	childrenViews := make([]string, 0, len(m.base.GetChildren()))
	for _, child := range m.base.GetChildren() {
		childrenViews = append(childrenViews, child.Model.View())
	}
	return m.style.Render(lipgloss.JoinVertical(lipgloss.Left, childrenViews...))
}

func (m model[T]) Base() *app.Base[T] {
	m.base.Model = m
	return m.base
}
