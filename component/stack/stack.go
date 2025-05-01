package stack

import (
	"github.com/alexanderbh/bubbleapp/app"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type Options[T any] struct {
	Horizontal bool
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

		// Determine layout direction and corresponding properties
		totalSize := msg.Height
		perpendicularSize := msg.Width
		growFlag := func(o app.BaseOptions) bool { return o.GrowY }
		measureFunc := lipgloss.Height
		if m.opts.Horizontal {
			totalSize = msg.Width
			perpendicularSize = msg.Height
			growFlag = func(o app.BaseOptions) bool { return o.GrowX }
			measureFunc = lipgloss.Width
		}

		growerCount := 0
		for _, child := range m.base.Children {
			// Set the dimension perpendicular to the layout direction
			if m.opts.Horizontal {
				child.Height = perpendicularSize
			} else {
				child.Width = perpendicularSize
			}
			if growFlag(child.Opts) {
				growerCount++
			}
		}

		nonGrowerSize := 0
		for _, child := range m.base.Children {
			if !growFlag(child.Opts) {
				childSize := measureFunc(child.Model.View())
				nonGrowerSize += childSize

				childMsg := tea.WindowSizeMsg{}
				if m.opts.Horizontal {
					childMsg.Width = childSize
					childMsg.Height = perpendicularSize
				} else {
					childMsg.Width = perpendicularSize
					childMsg.Height = childSize
				}

				newChild, cmd := child.Model.Update(childMsg)
				newChildTyped := newChild.(app.UIModel[T])
				newChildTyped.Base().Model = newChildTyped
				m.base.ReplaceChild(child.ID, newChildTyped.Base())
				cmds = append(cmds, cmd)
			}
		}

		if growerCount > 0 {
			availableSize := totalSize - nonGrowerSize
			if availableSize < 0 {
				availableSize = 0
			}
			growerSize := 0
			remainder := 0
			if growerCount > 0 {
				growerSize = availableSize / growerCount
				remainder = availableSize % growerCount
			}

			for _, child := range m.base.Children {
				if growFlag(child.Opts) {
					currentGrowerSize := growerSize
					if remainder > 0 {
						currentGrowerSize++
						remainder--
					}

					childMsg := tea.WindowSizeMsg{}
					if m.opts.Horizontal {
						childMsg.Width = currentGrowerSize
						childMsg.Height = perpendicularSize
					} else {
						childMsg.Width = perpendicularSize
						childMsg.Height = currentGrowerSize
					}

					newChild, cmd := child.Model.Update(childMsg)
					newChildTyped := newChild.(app.UIModel[T])
					newChildTyped.Base().Model = newChildTyped
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
	childrenViews := make([]string, 0, len(m.base.Children))
	for _, child := range m.base.Children {
		childrenViews = append(childrenViews, child.Model.View())
	}

	if m.opts.Horizontal {
		return m.style.Render(lipgloss.JoinHorizontal(lipgloss.Top, childrenViews...))
	}
	return m.style.Render(lipgloss.JoinVertical(lipgloss.Left, childrenViews...))
}

func (m model[T]) Base() *app.Base[T] {
	m.base.Model = m
	return m.base
}
