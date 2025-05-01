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

	contentBoxID string
	titlesID     string

	tabContent []*app.Base[T]
}

func New[T any](ctx *app.Context[T], tabs []TabElement[T]) *app.Base[T] {

	tabTitles := make([]string, len(tabs))
	tabContent := make([]*app.Base[T], len(tabs))

	for i, tab := range tabs {
		tabTitles[i] = tab.Title
		tabContent[i] = tab.Content(ctx)
	}

	idPrefix := ctx.Zone.NewPrefix()
	tabTitlesModel := tabtitles.New(ctx, tabTitles, idPrefix+"-titles")

	base := app.New(ctx, app.WithGrow(true))

	contentBox := box.New(ctx, &box.Options[T]{
		Child: tabContent[0],
	})
	stackChild := stack.New(ctx, &stack.Options[T]{
		Children: []*app.Base[T]{
			tabTitlesModel,
			contentBox,
		},
	})

	base.AddChild(stackChild)

	return model[T]{
		ID: idPrefix,

		tabContent: tabContent,
		titlesID:   tabTitlesModel.ID,

		base:         base,
		contentBoxID: contentBox.ID,
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
	case tabtitles.TabChangedMsg:
		currentContentBoxID := m.contentBoxID
		currentContentBox := m.Base().GetChild(currentContentBoxID)
		newTabContent := m.tabContent[msg.ActiveTab]
		newBox := box.New(m.base.Ctx, &box.Options[T]{
			Child: newTabContent,
		})
		updatedModel, _ := newBox.Model.Update(
			tea.WindowSizeMsg{
				Width:  currentContentBox.Width,
				Height: currentContentBox.Height,
			},
		)
		typedUpdatedModel := updatedModel.(app.UIModel[T])
		m.contentBoxID = typedUpdatedModel.Base().ID
		m.base.Children[0].ReplaceChild(currentContentBoxID, typedUpdatedModel.Base())

		cmds = append(cmds, newBox.Model.Init())
		return m, tea.Batch(cmds...)
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
