package stack

import (
	"strings"

	"github.com/alexanderbh/bubbleapp/app"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type Options struct {
	Horizontal bool
	RowGap     int
}

type stack[T any] struct {
	base     *app.Base
	opts     Options
	style    lipgloss.Style
	children func(ctx *app.Context[T]) []app.Fc[T]
}

func New[T any](ctx *app.Context[T], children []app.Fc[T], options *Options, baseOptions ...app.BaseOption) *stack[T] {
	return NewDynamic(ctx, func(ctx *app.Context[T]) []app.Fc[T] {
		return children
	}, options, baseOptions...)
}

func NewDynamic[T any](ctx *app.Context[T], children func(ctx *app.Context[T]) []app.Fc[T], options *Options, baseOptions ...app.BaseOption) *stack[T] {
	if options == nil {
		options = &Options{}
	}
	if baseOptions == nil {
		baseOptions = []app.BaseOption{}
	}
	dir := app.Vertical
	if options.Horizontal {
		dir = app.Horizontal
	}
	base := app.NewBase[T](append([]app.BaseOption{app.WithGrow(true), app.WithLayoutDirection(dir)}, baseOptions...)...)

	return &stack[T]{
		base:     base,
		opts:     *options,
		style:    lipgloss.NewStyle(),
		children: children,
	}
}

func (m *stack[T]) Render(ctx *app.Context[T]) string {
	children := m.children(ctx)
	childrenViews := make([]string, 0, len(children))
	for _, child := range children {
		childRender := child.Render(ctx)
		if m.opts.Horizontal {

		} else {
			childRender = childRender + strings.Repeat("\n", m.opts.RowGap)
		}
		childrenViews = append(childrenViews, childRender)

	}

	if m.opts.Horizontal {
		return m.style.Render(lipgloss.JoinHorizontal(lipgloss.Top, childrenViews...))
	}
	return m.style.Render(lipgloss.JoinVertical(lipgloss.Left, childrenViews...))
}

func (m *stack[T]) Update(ctx *app.Context[T], msg tea.Msg) {
}

func (m *stack[T]) Children(ctx *app.Context[T]) []app.Fc[T] {
	children := m.children(ctx)
	return children
}

func (m *stack[T]) Base() *app.Base {
	return m.base
}

// func (m model[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	var (
// 		cmd  tea.Cmd
// 		cmds []tea.Cmd
// 	)

// 	switch msg := msg.(type) {
// 	case tea.WindowSizeMsg:
// 		m.base.Height = msg.Height
// 		m.base.Width = msg.Width
// 		m.style = m.style.Width(msg.Width).Height(msg.Height)

// 		// Determine layout direction and corresponding properties
// 		totalSize := msg.Height
// 		perpendicularSize := msg.Width
// 		growFlag := func(o app.BaseOptions) bool { return o.GrowY }
// 		measureFunc := lipgloss.Height
// 		if m.opts.Horizontal {
// 			totalSize = msg.Width
// 			perpendicularSize = msg.Height
// 			growFlag = func(o app.BaseOptions) bool { return o.GrowX }
// 			measureFunc = lipgloss.Width
// 		}

// 		growerCount := 0
// 		for _, child := range m.base.Children {
// 			// Set the dimension perpendicular to the layout direction
// 			if m.opts.Horizontal {
// 				child.Height = perpendicularSize
// 			} else {
// 				child.Width = perpendicularSize
// 			}
// 			if growFlag(child.Opts) {
// 				growerCount++
// 			}
// 		}

// 		nonGrowerSize := 0
// 		for _, child := range m.base.Children {
// 			if !growFlag(child.Opts) {
// 				childSize := measureFunc(child.Model.View())
// 				nonGrowerSize += childSize

// 				childMsg := tea.WindowSizeMsg{}
// 				if m.opts.Horizontal {
// 					childMsg.Width = childSize
// 					childMsg.Height = perpendicularSize
// 				} else {
// 					childMsg.Width = perpendicularSize
// 					childMsg.Height = childSize
// 				}

// 				newChild, cmd := child.Model.Update(childMsg)
// 				newChildTyped := newChild.(app.UIModel[T])
// 				newChildTyped.Base().Model = newChildTyped
// 				m.base.ReplaceChild(child.ID, newChildTyped.Base())
// 				cmds = append(cmds, cmd)
// 			}
// 		}

// 		if growerCount > 0 {
// 			availableSize := totalSize - nonGrowerSize
// 			if availableSize < 0 {
// 				availableSize = 0
// 			}
// 			growerSize := 0
// 			remainder := 0
// 			if growerCount > 0 {
// 				growerSize = availableSize / growerCount
// 				remainder = availableSize % growerCount
// 			}

// 			for _, child := range m.base.Children {
// 				if growFlag(child.Opts) {
// 					currentGrowerSize := growerSize
// 					if remainder > 0 {
// 						currentGrowerSize++
// 						remainder--
// 					}

// 					childMsg := tea.WindowSizeMsg{}
// 					if m.opts.Horizontal {
// 						childMsg.Width = currentGrowerSize
// 						childMsg.Height = perpendicularSize
// 					} else {
// 						childMsg.Width = perpendicularSize
// 						childMsg.Height = currentGrowerSize
// 					}

// 					newChild, cmd := child.Model.Update(childMsg)
// 					newChildTyped := newChild.(app.UIModel[T])
// 					newChildTyped.Base().Model = newChildTyped
// 					m.base.ReplaceChild(child.ID, newChildTyped.Base())
// 					cmds = append(cmds, cmd)
// 				}
// 			}
// 		}

// 		return m, tea.Batch(cmds...)
// 	}

// 	cmd = app.UpdateChildren(m.base, msg)
// 	cmds = append(cmds, cmd)

// 	return m, tea.Batch(cmds...)
// }
