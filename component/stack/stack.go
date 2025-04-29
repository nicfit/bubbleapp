package stack

import (
	"github.com/alexanderbh/bubbleapp/app"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type StackOptions struct {
	Vertical bool
}
type model struct {
	base  *app.Base
	opts  StackOptions
	style lipgloss.Style
}

type StackOption func(o *StackOptions)

func New(ctx *app.Context, opts ...StackOption) model {
	options := StackOptions{
		Vertical: true,
	}
	for _, opt := range opts {
		opt(&options)
	}
	base := app.New(ctx, app.WithGrow(true))

	return model{
		base:  base,
		opts:  options,
		style: lipgloss.NewStyle(),
	}
}

func (m *model) AddChild(item app.UIModel) {
	m.base.AddChild(item)
}
func (m *model) AddChildren(items ...app.UIModel) {
	for _, item := range items {
		m.base.AddChild(item)
	}
}

func (m model) Init() tea.Cmd {
	return m.base.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.base.Height = msg.Height
		m.base.Width = msg.Width
		m.style = m.style.Width(msg.Width).Height(msg.Height)

		growerCount := 0
		for _, child := range m.base.GetChildren() {
			if child.Base().Opts.GrowY {
				growerCount++
			}
		}

		nonGrowerHeight := 0
		for _, child := range m.base.GetChildren() {
			if !child.Base().Opts.GrowY {
				childHeight := lipgloss.Height(child.View())
				nonGrowerHeight += childHeight

				newChild, cmd := child.Update(tea.WindowSizeMsg{
					Width:  msg.Width,
					Height: childHeight,
				})
				newChildTyped := newChild.(app.UIModel)
				m.base.ReplaceChild(child.Base().ID, newChildTyped)
				cmds = append(cmds, cmd)
			}
		}

		if growerCount > 0 {
			availableHeight := msg.Height - nonGrowerHeight
			growerHeight := availableHeight / growerCount
			remainder := availableHeight % growerCount
			for _, child := range m.base.GetChildren() {
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
					newChildTyped := newChild.(app.UIModel)
					m.base.ReplaceChild(child.Base().ID, newChildTyped)
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

func (m model) View() string {

	childrenViews := make([]string, 0, len(m.base.GetChildren()))
	for _, child := range m.base.GetChildren() {
		childrenViews = append(childrenViews, child.View())
	}
	return m.style.Render(lipgloss.JoinVertical(lipgloss.Left, childrenViews...))
}

func (m model) Base() *app.Base {
	return m.base
}
