package tabtitles

import (
	"strconv"
	"strings"

	"github.com/alexanderbh/bubbleapp/app"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type TabChangedMsg struct {
	ActiveTab int
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

type model[T any] struct {
	ID        string
	base      *app.Base[T]
	titles    []string
	activeTab int

	inactiveTabStyle        lipgloss.Style
	inactiveTabStyleFocused lipgloss.Style
	activeTabStyle          lipgloss.Style
	unusedTabStyle          lipgloss.Style
	unusedTabStyleFocused   lipgloss.Style
}

func New[T any](ctx *app.Context[T], titles []string, idPrefix string) *app.Base[T] {
	return model[T]{
		ID:        idPrefix,
		base:      app.New(ctx, app.WithFocusable(true)),
		titles:    titles,
		activeTab: 0,

		inactiveTabStyle:        lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(lipgloss.Color("#ACACAC")).Foreground(lipgloss.Color("#ACACAC")).Padding(0, 1),
		inactiveTabStyleFocused: lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(lipgloss.Color("#FFFFFF")).Foreground(lipgloss.Color("#FFFFFF")).Padding(0, 1),
		activeTabStyle:          lipgloss.NewStyle().Border(activeTabBorder, true).BorderForeground(lipgloss.Color("#0188a5")).Padding(0, 1),
		unusedTabStyle:          lipgloss.NewStyle().Border(unusedTabBorder, false, false, true, false).BorderForeground(lipgloss.Color("#ACACAC")).Foreground(lipgloss.Color("#ACACAC")),
		unusedTabStyleFocused:   lipgloss.NewStyle().Border(unusedTabBorder, false, false, true, false).BorderForeground(lipgloss.Color("#FFFFFF")).Foreground(lipgloss.Color("#FFFFFF")),
	}.Base()
}

func (m model[T]) Init() tea.Cmd {
	return nil
}

func (m model[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {

	case tea.KeyMsg:
		if m.base.Focused {
			switch keypress := msg.String(); keypress {
			case "right":
				newTab := min(m.activeTab+1, len(m.titles)-1)
				if newTab != m.activeTab {
					m.activeTab = newTab
					return m, func() tea.Msg { return TabChangedMsg{ActiveTab: m.activeTab} }
				}
			case "left":
				newTab := max(m.activeTab-1, 0)
				if newTab != m.activeTab {
					m.activeTab = newTab
					return m, func() tea.Msg { return TabChangedMsg{ActiveTab: m.activeTab} }
				}
			}
		}
	case tea.MouseClickMsg:
		if msg.Button == tea.MouseLeft {
			for i := range m.titles {
				if m.base.Ctx.Zone.Get(m.ID + strconv.Itoa(i)).InBounds(msg) {
					if i != m.activeTab {
						m.activeTab = i
						return m, func() tea.Msg { return TabChangedMsg{ActiveTab: m.activeTab} }
					}
					break
				}
			}
		}
	}

	cmd = m.base.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model[T]) View() string {
	doc := strings.Builder{}
	var renderedTabs []string

	for i, t := range m.titles {
		var style lipgloss.Style
		isActive := i == m.activeTab
		if isActive {
			style = m.activeTabStyle
		} else {
			if m.base.Focused {
				style = m.inactiveTabStyleFocused
			} else {
				style = m.inactiveTabStyle
			}
		}
		currentRow := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)

		renderedTab := m.base.Ctx.Zone.Mark(m.ID+strconv.Itoa(i), style.Render(t))
		if lipgloss.Width(currentRow)+lipgloss.Width(renderedTab) > m.base.Width {
			doc.WriteString(currentRow + "\n")
			renderedTabs = []string{}
		}
		renderedTabs = append(renderedTabs, renderedTab)
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	if lipgloss.Width(row) < m.base.Width {
		var style lipgloss.Style
		if m.base.Focused {
			style = m.unusedTabStyleFocused
		} else {
			style = m.unusedTabStyle
		}
		row = lipgloss.JoinHorizontal(lipgloss.Center, row, style.Render(strings.Repeat(" ", m.base.Width-lipgloss.Width(row))))
	}

	doc.WriteString(row)
	return doc.String()
}

func (m model[T]) Base() *app.Base[T] {
	m.base.Model = m
	return m.base
}
