package tabs

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/box"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/tabtitles"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type TabElement[T any] struct {
	Title   string
	Content func(ctx *app.Context[T]) app.Fc[T]
}

type uiState struct {
	activeTab int
}

type tabs[T any] struct {
	base *app.Base

	tabs []TabElement[T]

	root app.Fc[T]
}

func New[T any](ctx *app.Context[T], ts []TabElement[T], baseOptions ...app.BaseOption) *tabs[T] {
	if baseOptions == nil {
		baseOptions = []app.BaseOption{}
	}

	base, cleanup := app.NewBase(ctx, "tabs", append([]app.BaseOption{app.WithGrow(true)}, baseOptions...)...)
	defer cleanup()

	tabTitles := make([]string, len(ts))

	for i, tab := range ts {
		tabTitles[i] = tab.Title
	}

	state := app.GetUIState[T, uiState](ctx, base.ID)
	if state == nil {
		state = &uiState{
			activeTab: 0,
		}
		app.SetUIState(ctx, base.ID, state)
	}

	stackChild := stack.New(ctx, func(ctx *app.Context[T]) []app.Fc[T] {
		return []app.Fc[T]{
			tabtitles.New(ctx, tabTitles, func(activeTab int) {
				state.activeTab = activeTab
			}),
			box.New(ctx, func(ctx *app.Context[T]) app.Fc[T] {
				return ts[state.activeTab].Content(ctx)
			}, nil),
		}
	}, nil)

	return &tabs[T]{
		base: base,
		tabs: ts,
		root: stackChild,
	}
}
func (m *tabs[T]) Render(ctx *app.Context[T]) string {

	return m.root.Render(ctx)
}

func (m *tabs[T]) Update(ctx *app.Context[T], msg tea.Msg) {
	// 	var (
	// 		cmd  tea.Cmd
	// 		cmds []tea.Cmd
	// 	)

	// 	switch msg := msg.(type) {
	// 	case tabtitles.TabChangedMsg:
	// 		currentContentBoxID := m.contentBoxID
	// 		currentContentBox := m.Base().GetChild(currentContentBoxID)
	// 		newTabContent := m.tabContent[msg.ActiveTab]
	// 		newBox := box.New(m.base.Ctx, &box.Options[T]{
	// 			Child: newTabContent,
	// 		})
	// 		updatedModel, _ := newBox.Model.Update(
	// 			tea.WindowSizeMsg{
	// 				Width:  currentContentBox.Width,
	// 				Height: currentContentBox.Height,
	// 			},
	// 		)
	// 		typedUpdatedModel := updatedModel.(app.UIModel[T])
	// 		m.contentBoxID = typedUpdatedModel.Base().ID
	// 		m.base.Children[0].ReplaceChild(currentContentBoxID, typedUpdatedModel.Base())

	// 		cmds = append(cmds, newBox.Model.Init())
	// 		return m, tea.Batch(cmds...)
	// 	}

	// 	cmd = m.base.Update(msg)
	// 	cmds = append(cmds, cmd)

	// 	return m, tea.Batch(cmds...)
}

func (m *tabs[T]) Children(ctx *app.Context[T]) []app.Fc[T] {
	return []app.Fc[T]{m.root}
}
func (m *tabs[T]) Base() *app.Base {
	return m.base
}
