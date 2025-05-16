package tabtitles

import (
	"strconv"
	"strings"

	"github.com/alexanderbh/bubbleapp/app" // Assuming app.Ctx, app.Fc, app.View (string), etc. are defined here
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

// tabBorderWithBottom is a helper to create a lipgloss.Border with custom bottom characters.
func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

var (
	defaultInactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	defaultActiveTabBorder   = tabBorderWithBottom("┘", " ", "└")
	defaultUnusedTabBorder   = tabBorderWithBottom("┘", "─", " ")
)

// Props defines the properties for the TabTitles component.
type Props struct {
	Titles      []string
	ActiveTab   int
	OnTabChange func(activeID int)
	app.Layout
}

// Prop is a functional option for configuring TabTitles.
// No specific options are defined for TabTitles yet, but the type is kept for consistency.
type Prop func(*Props)

// New creates a new TabTitles component.
func New(c *app.Ctx, titles []string, activeTab int, onTabChange func(activeID int), opts ...Prop) app.C {
	p := Props{
		Titles:      titles,
		ActiveTab:   activeTab,
		OnTabChange: onTabChange,
		Layout: app.Layout{
			GrowX: true,
			GrowY: false,
		},
	}

	for _, opt := range opts {
		opt(&p)
	}

	if p.Titles == nil {
		p.Titles = []string{}
	}

	return c.Render(TabTitles, p)
}

// TabTitles is a functional component that displays a set of interactive tabs.
func TabTitles(ctx *app.Ctx, componentProps app.Props) string {
	p, _ := componentProps.(Props)

	// TODO: UseMemo for this
	activeStyle := lipgloss.NewStyle().Border(defaultActiveTabBorder, true).BorderForeground(ctx.Styles.Colors.Secondary).Padding(0, 1)
	inactiveStyle := lipgloss.NewStyle().Border(defaultInactiveTabBorder, true).BorderForeground(lipgloss.Color("#ACACAC")).Foreground(lipgloss.Color("#ACACAC")).Padding(0, 1)
	inactiveStyleFocused := lipgloss.NewStyle().Border(defaultInactiveTabBorder, true).BorderForeground(lipgloss.Color("#FFFFFF")).Foreground(lipgloss.Color("#FFFFFF")).Padding(0, 1)
	hoveredStyle := lipgloss.NewStyle().Border(defaultInactiveTabBorder, true).BorderForeground(ctx.Styles.Colors.Primary).Foreground(ctx.Styles.Colors.Primary).Padding(0, 1)
	unusedStyle := lipgloss.NewStyle().Border(defaultUnusedTabBorder, false, false, true, false).BorderForeground(lipgloss.Color("#ACACAC")).Foreground(lipgloss.Color("#ACACAC"))
	unusedStyleFocused := lipgloss.NewStyle().Border(defaultUnusedTabBorder, false, false, true, false).BorderForeground(lipgloss.Color("#FFFFFF")).Foreground(lipgloss.Color("#FFFFFF"))

	titles := p.Titles
	focused := app.UseIsFocused(ctx)
	_, hoveredChildID := app.UseIsHovered(ctx)

	app.UseKeyHandler(ctx, func(msg tea.KeyMsg) bool {
		numTitles := len(titles)
		if numTitles == 0 {
			return false
		}
		currentIndex := p.ActiveTab

		if currentIndex == -1 && numTitles > 0 {
			currentIndex = 0
		} else if currentIndex == -1 {
			return false
		}

		newIndex := currentIndex
		switch keypress := msg.String(); keypress {
		// Tab is hard to use as "change tab" so arrows are used.
		// The problem is if tab is used to change tab then it cannot
		// be used to change focus as well. So it will not be possible
		// to change focus to number 2 of 3 tabs for example. Since "Tab"
		// will just change the active tab
		case "right":
			newIndex = (currentIndex + 1) % numTitles
		case "left":
			newIndex = (currentIndex - 1 + numTitles) % numTitles
		default:
			return false
		}

		if newIndex != currentIndex {
			if p.OnTabChange != nil {
				p.OnTabChange(newIndex)
			}
			return true
		}
		return false
	})

	app.UseMouseHandler(ctx, func(msg tea.MouseMsg, childID string) bool {
		if p.OnTabChange == nil {
			return false
		}
		if _, ok := msg.(tea.MouseReleaseMsg); ok && msg.Mouse().Button == tea.MouseLeft {
			childSplit := strings.Split(childID, ":")
			i, _ := strconv.Atoi(childSplit[1])
			p.OnTabChange(i)
			return true
		}
		return false
	})

	// TODO add a UseMemo hook
	rowsBuilder := strings.Builder{}
	var currentLineTabs []string
	currentLineWidth := 0
	availableWidth, _ := app.UseSize(ctx)

	for i, titleInfo := range titles {
		tabChildGid := "tab:" + strconv.Itoa(i)
		isTabActive := i == p.ActiveTab

		isTabHovered := hoveredChildID == tabChildGid

		var currentStyle lipgloss.Style
		if isTabHovered {
			currentStyle = hoveredStyle
		} else if isTabActive {
			currentStyle = activeStyle
		} else {
			if focused {
				currentStyle = inactiveStyleFocused
			} else {
				currentStyle = inactiveStyle
			}
		}

		renderedTabString := currentStyle.Render(titleInfo)
		tabWidth := lipgloss.Width(renderedTabString)

		if currentLineWidth > 0 && currentLineWidth+tabWidth > availableWidth {
			rowsBuilder.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, currentLineTabs...))
			rowsBuilder.WriteString("\n")
			currentLineTabs = []string{}
			currentLineWidth = 0
		}

		currentLineTabs = append(currentLineTabs, ctx.MouseZoneChild(tabChildGid, renderedTabString))

		currentLineWidth += tabWidth
	}

	if len(currentLineTabs) > 0 {
		lastLineString := lipgloss.JoinHorizontal(lipgloss.Top, currentLineTabs...)
		// Fill remaining width on the last line
		if ctx.LayoutPhase == app.LayoutPhaseFinalRender && currentLineWidth < availableWidth {
			fillWidth := availableWidth - currentLineWidth
			if fillWidth > 0 {
				var fillStyle lipgloss.Style
				if focused {
					fillStyle = unusedStyleFocused
				} else {
					fillStyle = unusedStyle
				}
				lastLineString = lipgloss.JoinHorizontal(lipgloss.Center, lastLineString, fillStyle.Render(strings.Repeat(" ", fillWidth)))
			}
		}
		rowsBuilder.WriteString(lastLineString)
	}

	return rowsBuilder.String()
}
