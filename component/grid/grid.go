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

type GridItemConfig struct {
	item app.UIModel
	Xs   int
	Sm   int
	Md   int
	Lg   int
}

func (c GridItemConfig) GetSpanForWidth(width int) int {
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

type GridItemConfigOption func(*GridItemConfig)

func WithXs(xs int) GridItemConfigOption {
	return func(c *GridItemConfig) {
		c.Xs = xs
	}
}
func WithSm(sm int) GridItemConfigOption {
	return func(c *GridItemConfig) {
		c.Sm = sm
	}
}
func WithMd(md int) GridItemConfigOption {
	return func(c *GridItemConfig) {
		c.Md = md
	}
}
func WithLg(lg int) GridItemConfigOption {
	return func(c *GridItemConfig) {
		c.Lg = lg
	}
}
func NewItem(item app.UIModel, opts ...GridItemConfigOption) GridItemConfig {
	config := GridItemConfig{
		item: item,
		Xs:   0,
		Sm:   0,
		Md:   0,
		Lg:   0,
	}
	for _, opt := range opts {
		opt(&config)
	}
	if config.Xs == 0 {
		config.Xs = 12
	}
	return config
}

type model struct {
	base        *app.Base
	itemConfigs map[string]GridItemConfig
}

func New(ctx *app.Context) model {
	return model{
		base:        app.New(ctx, app.WithGrow(true)),
		itemConfigs: make(map[string]GridItemConfig),
	}
}

func (m model) AddItems(items ...GridItemConfig) {
	for _, item := range items {
		itemBox := box.New(m.base.Ctx)
		itemBox.AddChild(item.item)
		m.base.AddChild(itemBox)
		m.itemConfigs[itemBox.Base().ID] = item
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
			children       []app.UIModel
			maxChildHeight int // Max height of non-growing children in this row
			hasGrower      bool
			totalSpan      int
		}
		var rows []rowInfo
		var currentRow rowInfo
		currentRowSpan := 0
		widthPerSpanUnit := float64(containerWidth) / 12.0

		children := m.base.GetChildren()
		childWidths := make(map[string]int) // Store calculated widths

		for i, child := range children {
			childID := child.Base().ID
			config, ok := m.itemConfigs[childID]
			if !ok {
				config = GridItemConfig{Xs: 12} // Default fallback
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
					currentTotalWidth += childWidths[c.Base().ID]
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
			if child.Base().Opts.GrowY {
				currentRow.hasGrower = true
			} else {
				// Calculate natural height only for non-growers
				childHeight := lipgloss.Height(child.View()) // Get natural height
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
				childID := child.Base().ID
				targetWidth := childWidths[childID] // Use pre-calculated width

				// Determine target height for this child
				targetHeight := rowHeight // Default to row height
				if !row.hasGrower && !child.Base().Opts.GrowY {
					// If it's a non-growing row and non-growing child, use its natural height,
					// but capped by the row height (in case other items forced row higher).
					naturalHeight := lipgloss.Height(child.View())
					if naturalHeight < targetHeight {
						targetHeight = naturalHeight
					}
				}
				if targetHeight < 0 {
					targetHeight = 0
				}

				newChild, updateCmd := child.Update(tea.WindowSizeMsg{
					Width:  targetWidth,
					Height: targetHeight,
				})
				if newChildTyped, ok := newChild.(app.UIModel); ok {
					m.base.ReplaceChild(childID, newChildTyped)
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

func (m model) View() string {
	if len(m.base.GetChildren()) == 0 {
		return ""
	}

	containerWidth := m.base.Ctx.Width

	var rows [][]string
	var currentRowItems []string
	currentRowSpan := 0

	widthPerSpanUnit := float64(containerWidth) / 12.0
	remainder := containerWidth % 12 // is this right??

	for _, child := range m.base.GetChildren() {
		childID := child.Base().ID

		config, ok := m.itemConfigs[childID]
		if !ok {
			config = GridItemConfig{Xs: 12}
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

		child.Base().Width = targetChildWidth // this seems like a hack. something is off on this approach
		childView := child.View()

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

func (m model) Base() *app.Base {
	return m.base
}
