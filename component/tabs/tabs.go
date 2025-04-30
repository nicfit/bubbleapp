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
	Content func(ctx *app.Context[T]) *app.Base[T]
}

type model[T any] struct {
	ID string

	base *app.Base[T]

	activeTabID string
	titlesID    string

	tabContent []*app.Base[T]
}

func New[T any](ctx *app.Context[T], tabs []TabElement[T]) model[T] {

	tabTitles := make([]string, len(tabs))
	tabContent := make([]*app.Base[T], len(tabs))

	for i, tab := range tabs {
		tabTitles[i] = tab.Title
		tabContent[i] = tab.Content(ctx)
	}

	idPrefix := ctx.Zone.NewPrefix()
	tabTitlesModel := tabtitles.New(ctx, tabTitles, idPrefix+"-titles")

	base := app.New(ctx, app.WithGrow(true))

	contentBox := box.New(ctx, box.Options[T]{
		Child: tabContent[0],
	})
	stackChild := stack.New(ctx, stack.Options[T]{
		Children: []*app.Base[T]{
			tabTitlesModel.Base(),
			contentBox.Base(),
		},
	})

	base.AddChild(stackChild.Base())

	return model[T]{
		ID: idPrefix,

		tabContent: tabContent,
		titlesID:   tabTitlesModel.Base().ID,

		base:        base,
		activeTabID: contentBox.Base().GetChildren()[0].ID,
	}
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
	case tabtitles.TabChangedMsg:
		currentTabID := m.activeTabID
		newTab := m.tabContent[msg.ActiveTab]
		m.activeTabID = newTab.ID
		m.base.GetChildren()[0].GetChildren()[1].ReplaceChild(currentTabID, newTab)

		cmds = append(cmds, newTab.Init(), func() tea.Msg {
			return tea.WindowSizeMsg{
				Width:  m.base.Width,
				Height: m.base.Height,
			}
		})
	}

	cmd = m.base.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model[T]) View() string {
	return m.base.Render()
}

func (m model[T]) Base() *app.Base[T] {
	m.base.Model = m
	return m.base
}
