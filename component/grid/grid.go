package grid

import (
	"math"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/box"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

const (
	breakpointSm = 60
	breakpointMd = 90
	breakpointLg = 120
)

type Item[T any] struct {
	Item *app.Base[T]
	Xs   int
	Sm   int
	Md   int
	Lg   int
}

type model[T any] struct {
	base        *app.Base[T]
	itemConfigs map[string]Item[T]
}

func New[T any](ctx *app.Context[T], items ...Item[T]) *app.Base[T] {
	m := model[T]{
		base:        app.New(ctx, app.WithGrow(true)),
		itemConfigs: make(map[string]Item[T]),
	}

	m.addItems(items...)

	return m.Base()
}

func (m model[T]) addItems(items ...Item[T]) {
	for _, item := range items {
		if item.Xs == 0 {
			item.Xs = 12
		}
		// We need a box here for now. To ensure it has grow set to fill out the grid cell
		itemBox := box.New(m.base.Ctx, &box.Options[T]{
			Child: item.Item,
		})
		m.base.AddChild(itemBox)
		m.itemConfigs[itemBox.ID] = item
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
	case tea.WindowSizeMsg:
		m.base.Height = msg.Height
		m.base.Width = msg.Width
		containerWidth := msg.Width

		// OKAY THIS CODE BELOW HERE IS PURE AI VIBING.
		// I just wanted some pointers on how to do this, and it just blurted out this mess.

		// But so far it seems to work.. I will have to read and understand it at somepoint.
		// Still not sure it is the right approach.

		// AI VIBING BELOW
		type rowInfo struct {
			children       []*app.Base[T]
			maxChildHeight int // Max height of non-growing children in this row
			hasGrower      bool
			totalSpan      int
		}
		var rows []rowInfo
		var currentRow rowInfo
		currentRowSpan := 0
		widthPerSpanUnit := float64(containerWidth) / 12.0

		children := m.base.Children
		childWidths := make(map[string]int) // Store calculated widths

		for i, child := range children {
			childID := child.ID
			config, ok := m.itemConfigs[childID]
			if !ok {
				config = Item[T]{Xs: 12} // Default fallback
			}
			childSpan := config.GetSpanForWidth(containerWidth)

			// Calculate target width (similar to View logic, simplified remainder handling)
			targetChildWidth := int(math.Round(widthPerSpanUnit * float64(childSpan)))
			if targetChildWidth < 1 {
				targetChildWidth = 1
			}
			// Ensure total width doesn't exceed container width due to rounding
			if i == len(children)-1 && currentRowSpan+childSpan <= 12 {
				currentTotalWidth := 0
				for _, c := range currentRow.children {
					currentTotalWidth += childWidths[c.ID]
				}
				targetChildWidth = containerWidth - currentTotalWidth
			} else if currentRowSpan+childSpan > 12 {
				// If adding this child exceeds 12, calculate width based on remaining space in the *next* row (which is just this child for now)
				targetChildWidth = int(math.Round(widthPerSpanUnit * float64(childSpan)))
				if targetChildWidth < 1 {
					targetChildWidth = 1
				}
				if targetChildWidth > containerWidth {
					targetChildWidth = containerWidth
				}
			}

			childWidths[childID] = targetChildWidth // Store width

			if currentRowSpan+childSpan > 12 && len(currentRow.children) > 0 {
				// Finish previous row
				currentRow.totalSpan = currentRowSpan
				rows = append(rows, currentRow)
				// Start new row
				currentRow = rowInfo{}
				currentRowSpan = 0
			}

			// Add child to current row
			currentRow.children = append(currentRow.children, child)
			currentRowSpan += childSpan
			if child.Opts.GrowY {
				currentRow.hasGrower = true
			} else {
				// Calculate natural height only for non-growers
				childHeight := lipgloss.Height(child.Model.View()) // Get natural height
				if childHeight > currentRow.maxChildHeight {
					currentRow.maxChildHeight = childHeight
				}
			}
		}
		// Add the last row
		if len(currentRow.children) > 0 {
			currentRow.totalSpan = currentRowSpan
			rows = append(rows, currentRow)
		}

		// --- 2. Calculate Height Distribution ---
		nonGrowerTotalHeight := 0
		growerRowCount := 0
		for _, row := range rows {
			if !row.hasGrower {
				nonGrowerTotalHeight += row.maxChildHeight
			} else {
				growerRowCount++
			}
		}

		availableHeightForGrowers := msg.Height - nonGrowerTotalHeight
		if availableHeightForGrowers < 0 {
			availableHeightForGrowers = 0 // Avoid negative height
		}

		heightPerGrowerRow := 0
		remainder := 0
		if growerRowCount > 0 {
			heightPerGrowerRow = availableHeightForGrowers / growerRowCount
			remainder = availableHeightForGrowers % growerRowCount
		}

		// --- 3. Update Children Row by Row ---
		for _, row := range rows {
			rowHeight := 0
			if row.hasGrower {
				rowHeight = heightPerGrowerRow
				if remainder > 0 {
					rowHeight++
					remainder--
				}
			} else {
				rowHeight = row.maxChildHeight
			}
			if rowHeight < 0 {
				rowHeight = 0
			} // Ensure non-negative height

			// Distribute height within the row
			// Simple approach: Give all children in the row the calculated rowHeight.
			// More complex: Distribute remaining height in growing rows among growing children.
			// Let's use the simple approach for now.
			for _, child := range row.children {
				childID := child.ID
				targetWidth := childWidths[childID] // Use pre-calculated width

				// Determine target height for this child
				targetHeight := rowHeight // Default to row height
				if !row.hasGrower && !child.Opts.GrowY {
					// If it's a non-growing row and non-growing child, use its natural height,
					// but capped by the row height (in case other items forced row higher).
					naturalHeight := lipgloss.Height(child.Model.View())
					if naturalHeight < targetHeight {
						targetHeight = naturalHeight
					}
				}
				if targetHeight < 0 {
					targetHeight = 0
				}

				newChild, updateCmd := child.Model.Update(tea.WindowSizeMsg{
					Width:  targetWidth,
					Height: targetHeight,
				})
				if newChildTyped, ok := newChild.(app.UIModel[T]); ok {
					newChildTyped.Base().Model = newChildTyped
					m.base.ReplaceChild(childID, newChildTyped.Base())
				} // else: handle error or unexpected type?
				if updateCmd != nil {
					cmds = append(cmds, updateCmd)
				}
			}
		}

		return m, tea.Batch(cmds...)
		// AI VIBING END
	}

	cmd = m.base.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model[T]) View() string {
	if len(m.base.Children) == 0 {
		return ""
	}

	containerWidth := m.base.Ctx.Width

	var rows [][]string
	var currentRowItems []string
	currentRowSpan := 0

	widthPerSpanUnit := float64(containerWidth) / 12.0
	remainder := containerWidth % 12 // is this right??

	for _, child := range m.base.Children {
		childID := child.ID

		config, ok := m.itemConfigs[childID]
		if !ok {
			config = Item[T]{Xs: 12}
		}

		childSpan := config.GetSpanForWidth(containerWidth)

		targetChildWidth := int(math.Floor(widthPerSpanUnit * float64(childSpan)))
		if remainder > 0 {
			targetChildWidth++
			remainder--
		}
		if targetChildWidth < 1 {
			targetChildWidth = 1
		}

		child.Width = targetChildWidth // this seems like a hack. something is off on this approach
		childView := child.Model.View()

		styledChildView := lipgloss.NewStyle().Render(childView)

		if currentRowSpan+childSpan > 12 {
			if len(currentRowItems) > 0 {
				rows = append(rows, currentRowItems)
			}
			currentRowItems = []string{styledChildView}
			currentRowSpan = childSpan
		} else {
			currentRowItems = append(currentRowItems, styledChildView)
			currentRowSpan += childSpan
		}
	}

	if len(currentRowItems) > 0 {
		rows = append(rows, currentRowItems)
	}

	renderedRows := make([]string, len(rows))
	for i, rowItems := range rows {
		renderedRows[i] = lipgloss.JoinHorizontal(lipgloss.Top, rowItems...)
	}

	return lipgloss.JoinVertical(lipgloss.Left, renderedRows...)
}

func (m model[T]) Base() *app.Base[T] {
	m.base.Model = m
	return m.base
}

func (c Item[T]) GetSpanForWidth(width int) int {
	span := c.Xs
	if span <= 0 {
		span = 12
	}

	if c.Sm > 0 && width >= breakpointSm {
		span = c.Sm
	}
	if c.Md > 0 && width >= breakpointMd {
		span = c.Md
	} else if width >= breakpointMd && c.Sm > 0 {
		span = c.Sm
	}

	if c.Lg > 0 && width >= breakpointLg {
		span = c.Lg
	} else if width >= breakpointLg {
		if c.Md > 0 {
			span = c.Md
		} else if c.Sm > 0 {
			span = c.Sm
		}
	}

	if span < 1 {
		return 1
	}
	if span > 12 {
		return 12
	}
	return span
}
