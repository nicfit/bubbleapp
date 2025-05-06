package grid

// import (
// 	"math"

// 	"github.com/alexanderbh/bubbleapp/app"
// 	"github.com/alexanderbh/bubbleapp/component/box"
// 	tea "github.com/charmbracelet/bubbletea/v2"
// 	"github.com/charmbracelet/lipgloss/v2"
// )

// // Disclaimer: The size calculations are mostly AI vibed.
// // Should look at it at some point but it just spit it out and it seems to work.
// const (
// 	breakpointSm = 60
// 	breakpointMd = 90
// 	breakpointLg = 120
// )

// type Options[T any] struct {
// 	Items []Item[T]
// }

// type Item[T any] struct {
// 	Item *app.Base
// 	Xs   int
// 	Sm   int
// 	Md   int
// 	Lg   int
// }

// type model[T any] struct {
// 	base        *app.Base
// 	itemConfigs map[string]Item[T]
// }

// func New[T any](ctx *app.Context[T], options *Options[T], baseOptions ...app.BaseOption) *app.Base {
// 	if options == nil {
// 		options = &Options[T]{}
// 	}
// 	if baseOptions == nil {
// 		baseOptions = []app.BaseOption{}
// 	}
// 	base, cleanup := app.NewBase(ctx, "grid", append([]app.BaseOption{app.WithGrow(true)}, baseOptions...)...)
// 	defer cleanup()

// 	m := model[T]{
// 		base:        base,
// 		itemConfigs: make(map[string]Item[T]),
// 	}

// 	if options.Items != nil {
// 		m.addItems(options.Items...)
// 	}

// 	return m.Base()
// }

// func (m model[T]) addItems(items ...Item[T]) {
// 	for _, item := range items {
// 		if item.Xs == 0 {
// 			item.Xs = 12
// 		}
// 		// We need a box here for now. To ensure it has grow set to fill out the grid cell
// 		itemBox := box.New(m.base.Ctx, &box.Options[T]{
// 			Child: item.Item,
// 		})
// 		m.base.AddChild(itemBox)
// 		m.itemConfigs[itemBox.ID] = item
// 	}
// }

// func (m model[T]) Init() tea.Cmd {
// 	return m.base.Init()
// }

// func (m model[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	var (
// 		cmd  tea.Cmd
// 		cmds []tea.Cmd
// 	)

// 	switch msg := msg.(type) {
// 	case tea.WindowSizeMsg:
// 		m.base.Height = msg.Height
// 		m.base.Width = msg.Width
// 		containerWidth := msg.Width
// 		if containerWidth <= 0 { // Handle zero/negative width
// 			// Set all children to width 0 and return
// 			for _, child := range m.base.Children {
// 				newChild, updateCmd := child.Model.Update(tea.WindowSizeMsg{Width: 0, Height: 0})
// 				if newChildTyped, ok := newChild.(app.UIModel[T]); ok {
// 					newChildTyped.Base().Model = newChildTyped
// 					newChildTyped.Base().Width = 0
// 					newChildTyped.Base().Height = 0
// 					m.base.ReplaceChild(child.ID, newChildTyped.Base())
// 				}
// 				if updateCmd != nil {
// 					cmds = append(cmds, updateCmd)
// 				}
// 			}
// 			return m, tea.Batch(cmds...)
// 		}

// 		// --- 1. Group children into rows and calculate initial layout ---
// 		type rowInfo struct {
// 			children       []*app.Base
// 			spans          []int // Store spans for width calculation
// 			maxChildHeight int   // Max height of non-growing children in this row (estimate)
// 			hasGrower      bool
// 		}
// 		var rows []rowInfo
// 		var currentRow rowInfo
// 		currentRowSpan := 0

// 		children := m.base.Children

// 		for _, child := range children {
// 			childID := child.ID
// 			config, ok := m.itemConfigs[childID]
// 			if !ok {
// 				config = Item[T]{Xs: 12} // Default fallback
// 			}
// 			childSpan := config.GetSpanForWidth(containerWidth)

// 			if currentRowSpan+childSpan > 12 && len(currentRow.children) > 0 {
// 				rows = append(rows, currentRow)
// 				currentRow = rowInfo{}
// 				currentRowSpan = 0
// 			}

// 			currentRow.children = append(currentRow.children, child)
// 			currentRow.spans = append(currentRow.spans, childSpan) // Store span
// 			currentRowSpan += childSpan
// 			if child.Opts.GrowY {
// 				currentRow.hasGrower = true
// 			}
// 			// Note: Calculating natural height here uses the *previous* view state.
// 			// It's an estimate used for height planning before the width update.
// 			if !child.Opts.GrowY {
// 				childHeight := lipgloss.Height(child.Model.View())
// 				if childHeight > currentRow.maxChildHeight {
// 					currentRow.maxChildHeight = childHeight
// 				}
// 			}
// 		}
// 		if len(currentRow.children) > 0 {
// 			rows = append(rows, currentRow)
// 		}

// 		// --- Intermediate: Prepare map for final dimensions ---
// 		finalChildDims := make(map[string]struct{ Width, Height int })

// 		// --- 2. Calculate Width Distribution Per Row ---
// 		widthPerSpanUnit := float64(containerWidth) / 12.0
// 		for _, row := range rows {
// 			totalBaseWidth := 0
// 			childBaseWidths := make([]int, len(row.children))

// 			// Calculate base widths using floor
// 			for j, childSpan := range row.spans {
// 				baseWidth := int(math.Floor(widthPerSpanUnit * float64(childSpan)))
// 				if baseWidth < 0 {
// 					baseWidth = 0
// 				}
// 				childBaseWidths[j] = baseWidth
// 				totalBaseWidth += baseWidth
// 			}

// 			// Calculate remaining pixels for this row
// 			remainingPixels := containerWidth - totalBaseWidth
// 			if remainingPixels < 0 {
// 				remainingPixels = 0
// 			}

// 			// Distribute remaining pixels
// 			for j, child := range row.children {
// 				finalWidth := childBaseWidths[j]
// 				if remainingPixels > 0 {
// 					finalWidth++
// 					remainingPixels--
// 				}
// 				// Store final width for now, height comes next
// 				finalChildDims[child.ID] = struct{ Width, Height int }{Width: finalWidth}
// 			}
// 			// Add any leftover remainder (due to rounding edge cases or very small widths) to the last child.
// 			if remainingPixels > 0 && len(row.children) > 0 {
// 				lastChildID := row.children[len(row.children)-1].ID
// 				dims := finalChildDims[lastChildID]
// 				dims.Width += remainingPixels
// 				finalChildDims[lastChildID] = dims
// 			}
// 		}

// 		// --- 3. Calculate Height Distribution (largely unchanged) ---
// 		nonGrowerTotalHeight := 0
// 		growerRowCount := 0
// 		for i, row := range rows {
// 			if !row.hasGrower {
// 				// Re-estimate max natural height for non-growers in this row based on current view state
// 				maxNaturalHeightInRow := 0
// 				for _, child := range row.children {
// 					if !child.Opts.GrowY {
// 						naturalHeight := lipgloss.Height(child.Model.View())
// 						if naturalHeight > maxNaturalHeightInRow {
// 							maxNaturalHeightInRow = naturalHeight
// 						}
// 					}
// 				}
// 				rows[i].maxChildHeight = maxNaturalHeightInRow // Update row info used below
// 				nonGrowerTotalHeight += rows[i].maxChildHeight
// 			} else {
// 				growerRowCount++
// 			}
// 		}

// 		availableHeightForGrowers := msg.Height - nonGrowerTotalHeight
// 		if availableHeightForGrowers < 0 {
// 			availableHeightForGrowers = 0
// 		}

// 		heightPerGrowerRow := 0
// 		remainderHeight := 0
// 		if growerRowCount > 0 {
// 			heightPerGrowerRow = availableHeightForGrowers / growerRowCount
// 			remainderHeight = availableHeightForGrowers % growerRowCount
// 		}

// 		// --- 4. Update Children Row by Row with Final Dimensions ---
// 		for _, row := range rows {
// 			rowHeight := 0
// 			if row.hasGrower {
// 				rowHeight = heightPerGrowerRow
// 				if remainderHeight > 0 {
// 					rowHeight++
// 					remainderHeight--
// 				}
// 			} else {
// 				rowHeight = row.maxChildHeight // Use the potentially re-estimated max natural height
// 			}
// 			if rowHeight < 0 {
// 				rowHeight = 0
// 			}

// 			for _, child := range row.children {
// 				childID := child.ID
// 				dims := finalChildDims[childID] // Get calculated width
// 				targetWidth := dims.Width

// 				// Determine target height for this child
// 				targetHeight := rowHeight // Default to row height
// 				if !child.Opts.GrowY {    // If child itself doesn't grow vertically
// 					// Use its estimated natural height, capped by the row height.
// 					naturalHeight := lipgloss.Height(child.Model.View()) // Still using previous view height estimate
// 					if naturalHeight < targetHeight {
// 						targetHeight = naturalHeight
// 					}
// 				}
// 				if targetHeight < 0 {
// 					targetHeight = 0
// 				}

// 				// Store final height (optional, mainly for clarity if needed elsewhere)
// 				// dims.Height = targetHeight
// 				// finalChildDims[childID] = dims

// 				// Send update message
// 				newChild, updateCmd := child.Model.Update(tea.WindowSizeMsg{
// 					Width:  targetWidth,
// 					Height: targetHeight,
// 				})
// 				if newChildTyped, ok := newChild.(app.UIModel[T]); ok {
// 					newChildTyped.Base().Model = newChildTyped
// 					// Update the base Width/Height fields as well, as View might rely on them
// 					newChildTyped.Base().Width = targetWidth
// 					newChildTyped.Base().Height = targetHeight
// 					m.base.ReplaceChild(childID, newChildTyped.Base())
// 				} // else: handle error or unexpected type?
// 				if updateCmd != nil {
// 					cmds = append(cmds, updateCmd)
// 				}
// 			}
// 		}

// 		return m, tea.Batch(cmds...)
// 	}

// 	cmd = m.base.Update(msg)
// 	cmds = append(cmds, cmd)

// 	return m, tea.Batch(cmds...)
// }

// func (m model[T]) View() string {
// 	if len(m.base.Children) == 0 {
// 		return ""
// 	}

// 	containerWidth := m.base.Ctx.Width // Use context width reflecting the latest size
// 	if containerWidth <= 0 {
// 		return ""
// 	}

// 	var rows [][]string
// 	var currentRowItems []string
// 	currentRowSpan := 0

// 	// Group children into rows based on spans, needed for horizontal joining
// 	for _, child := range m.base.Children {
// 		childID := child.ID
// 		config, ok := m.itemConfigs[childID]
// 		if !ok {
// 			config = Item[T]{Xs: 12} // Default fallback
// 		}
// 		childSpan := config.GetSpanForWidth(containerWidth)

// 		// Get the child's view. It should have been rendered using the
// 		// dimensions calculated and passed during the Update phase.
// 		childView := child.Model.View()

// 		// We don't need to apply width styles here, as the child's view
// 		// should already be sized correctly based on the Width set in its Base
// 		// during the Update phase. lipgloss.JoinHorizontal respects this.
// 		styledChildView := childView // Use the view directly

// 		if currentRowSpan+childSpan > 12 && len(currentRowItems) > 0 {
// 			rows = append(rows, currentRowItems)
// 			currentRowItems = []string{styledChildView}
// 			currentRowSpan = childSpan
// 		} else {
// 			currentRowItems = append(currentRowItems, styledChildView)
// 			currentRowSpan += childSpan
// 		}
// 	}

// 	if len(currentRowItems) > 0 {
// 		rows = append(rows, currentRowItems)
// 	}

// 	// Join rows and items
// 	renderedRows := make([]string, len(rows))
// 	for i, rowItems := range rows {
// 		// lipgloss.JoinHorizontal arranges the views side-by-side.
// 		// It respects the width of each `childView` string.
// 		renderedRows[i] = lipgloss.JoinHorizontal(lipgloss.Top, rowItems...)
// 	}

// 	// Join the rows vertically
// 	finalView := lipgloss.JoinVertical(lipgloss.Left, renderedRows...)

// 	// Optional: If the grid container itself needs specific dimensions or styling:
// 	// finalView = lipgloss.NewStyle().Width(m.base.Width).Height(m.base.Height).Render(finalView)

// 	return finalView
// }

// func (m model[T]) Base() *app.Base {
// 	m.base.Model = m
// 	return m.base
// }

// func (c Item[T]) GetSpanForWidth(width int) int {
// 	span := c.Xs
// 	if span <= 0 {
// 		span = 12
// 	}

// 	if c.Sm > 0 && width >= breakpointSm {
// 		span = c.Sm
// 	}
// 	if c.Md > 0 && width >= breakpointMd {
// 		span = c.Md
// 	} else if width >= breakpointMd && c.Sm > 0 {
// 		span = c.Sm
// 	}

// 	if c.Lg > 0 && width >= breakpointLg {
// 		span = c.Lg
// 	} else if width >= breakpointLg {
// 		if c.Md > 0 {
// 			span = c.Md
// 		} else if c.Sm > 0 {
// 			span = c.Sm
// 		}
// 	}

// 	if span < 1 {
// 		return 1
// 	}
// 	if span > 12 {
// 		return 12
// 	}
// 	return span
// }
