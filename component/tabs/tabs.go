package tabs

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/box"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/tabtitles"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type TabElement struct {
	Title   string
	Content func(ctx *app.Context) app.UIModel
}

type model struct {
	ID string

	base *app.Base

	activeTabID string
	titlesID    string

	tabContent []app.UIModel
}

func New(ctx *app.Context, tabs []TabElement) model {

	tabTitles := make([]string, len(tabs))
	tabContent := make([]app.UIModel, len(tabs))

	for i, tab := range tabs {
		tabTitles[i] = tab.Title
		tabContent[i] = tab.Content(ctx)
	}

	idPrefix := ctx.Zone.NewPrefix()
	tabTitlesModel := tabtitles.New(ctx, tabTitles, idPrefix+"-titles")

	base := app.New(ctx, app.WithGrow(true))

	contentBox := box.New(ctx)
	contentBox.AddChild(tabContent[0])
	stackChild := stack.New(ctx)
	stackChild.AddChildren(tabTitlesModel, contentBox)

	base.AddChild(stackChild)

	return model{
		ID: idPrefix,

		tabContent: tabContent,
		titlesID:   tabTitlesModel.Base().ID,

		base:        base,
		activeTabID: contentBox.Base().GetChildren()[0].Base().ID,
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
	case tabtitles.TabChangedMsg:
		currentTabID := m.activeTabID
		newTab := m.tabContent[msg.ActiveTab]
		m.activeTabID = newTab.Base().ID
		m.base.GetChildren()[0].Base().GetChildren()[1].Base().ReplaceChild(currentTabID, newTab)

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

func (m model) View() string {
	return m.base.View()
}

func (m model) Base() *app.Base {
	return m.base
}
