package tabtitles

import (
	"strconv"
	"strings"

	"github.com/alexanderbh/bubbleapp/app"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type uiState struct {
	activeTab int
}

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

var (
	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder   = tabBorderWithBottom("┘", " ", "└")
	unusedTabBorder   = tabBorderWithBottom("┘", "─", " ")
)

type tabtitles[T any] struct {
	base       *app.Base
	titles     func(ctx *app.Context[T]) []string
	tabChanged func(activeTab int)

	inactiveTabStyle        lipgloss.Style
	inactiveTabStyleFocused lipgloss.Style
	activeTabStyle          lipgloss.Style
	unusedTabStyle          lipgloss.Style
	unusedTabStyleFocused   lipgloss.Style
}

func New[T any](ctx *app.Context[T], titles []string, tabChanged func(activeTab int), baseOptions ...app.BaseOption) *tabtitles[T] {
	return NewDynamic(ctx, func(ctx *app.Context[T]) []string {
		return titles
	}, tabChanged, baseOptions...)
}

func NewDynamic[T any](ctx *app.Context[T], titles func(ctx *app.Context[T]) []string, tabChanged func(activeTab int), baseOptions ...app.BaseOption) *tabtitles[T] {
	if baseOptions == nil {
		baseOptions = []app.BaseOption{}
	}

	base, cleanup := app.NewBase(ctx, "tabtitles", append([]app.BaseOption{app.WithFocusable(true), app.WithGrowX(true)}, baseOptions...)...)
	defer cleanup()

	return &tabtitles[T]{
		base:       base,
		titles:     titles,
		tabChanged: tabChanged,

		inactiveTabStyle:        lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(lipgloss.Color("#ACACAC")).Foreground(lipgloss.Color("#ACACAC")).Padding(0, 1),
		inactiveTabStyleFocused: lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(lipgloss.Color("#FFFFFF")).Foreground(lipgloss.Color("#FFFFFF")).Padding(0, 1),
		activeTabStyle:          lipgloss.NewStyle().Border(activeTabBorder, true).BorderForeground(lipgloss.Color("#0188a5")).Padding(0, 1),
		unusedTabStyle:          lipgloss.NewStyle().Border(unusedTabBorder, false, false, true, false).BorderForeground(lipgloss.Color("#ACACAC")).Foreground(lipgloss.Color("#ACACAC")),
		unusedTabStyleFocused:   lipgloss.NewStyle().Border(unusedTabBorder, false, false, true, false).BorderForeground(lipgloss.Color("#FFFFFF")).Foreground(lipgloss.Color("#FFFFFF")),
	}
}

func (m tabtitles[T]) Init() tea.Cmd {
	return nil
}

func (m tabtitles[T]) Update(ctx *app.Context[T], msg tea.Msg) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		uiState := m.getState(ctx)
		titles := m.titles(ctx)
		if ctx.UIState.Focused == m.base.ID {
			switch keypress := msg.String(); keypress {
			case "right":
				newTab := min(uiState.activeTab+1, len(titles)-1)
				if newTab != uiState.activeTab {
					uiState.activeTab = newTab
					m.tabChanged(uiState.activeTab)
					return
				}
			case "left":
				newTab := max(uiState.activeTab-1, 0)
				if newTab != uiState.activeTab {
					uiState.activeTab = newTab
					m.tabChanged(uiState.activeTab)
					return
				}
			}
		}
	case tea.MouseClickMsg:
		uiState := m.getState(ctx)
		titles := m.titles(ctx) // too expensive for hover state?
		if msg.Button == tea.MouseLeft {
			for i := range titles {
				if ctx.Zone.Get(m.base.ID + strconv.Itoa(i)).InBounds(msg) {
					if i != uiState.activeTab {
						uiState.activeTab = i
						m.tabChanged(uiState.activeTab)
						return
					}
					break
				}
			}
		}
	}

}

func (m *tabtitles[T]) Render(ctx *app.Context[T]) string {
	doc := strings.Builder{}
	var renderedTabs []string
	width := ctx.UIState.GetWidth(m.base.ID)

	for i, t := range m.titles(ctx) {
		var style lipgloss.Style
		isActive := i == m.getState(ctx).activeTab
		if isActive {
			style = m.activeTabStyle
		} else {
			if ctx.UIState.Focused == m.base.ID {
				style = m.inactiveTabStyleFocused
			} else {
				style = m.inactiveTabStyle
			}
		}
		currentRow := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)

		renderedTab := app.RegisterMouse(ctx, m.base.ID+strconv.Itoa(i), m, style.Render(t))
		if !ctx.LayoutPhase && lipgloss.Width(currentRow)+lipgloss.Width(renderedTab) > width {
			doc.WriteString(currentRow + "\n")
			renderedTabs = []string{}
		}
		renderedTabs = append(renderedTabs, renderedTab)
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	if lipgloss.Width(row) < width {
		var style lipgloss.Style
		if ctx.UIState.Focused == m.base.ID {
			style = m.unusedTabStyleFocused
		} else {
			style = m.unusedTabStyle
		}
		row = lipgloss.JoinHorizontal(lipgloss.Center, row, style.Render(strings.Repeat(" ", width-lipgloss.Width(row))))
	}

	doc.WriteString(row)
	return doc.String()
}

func (m *tabtitles[T]) Children(ctx *app.Context[T]) []app.Fc[T] {
	return nil
}
func (m *tabtitles[T]) Base() *app.Base {
	return m.base
}

func (m *tabtitles[T]) getState(ctx *app.Context[T]) *uiState {
	state := app.GetUIState[T, uiState](ctx, m.base.ID)
	if state == nil {
		state = &uiState{
			activeTab: 0,
		}
		app.SetUIState(ctx, m.base.ID, state)
	}
	return state
}
