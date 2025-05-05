package app

import (
	"github.com/charmbracelet/lipgloss/v2"
)

// Calculate sizes of the component tree
func (a *App[T]) Layout() {
	a.ctx.IDMap = make(map[string]Fc[T])
	a.ctx.ids = make([]string, 0)
	a.ctx.LayoutPhase = true
	defer func() {
		a.ctx.LayoutPhase = false
	}()

	// TODO: Can the Zone manager be reset here? If not why? Otherwise things will live in the zone forever.
	a.ctx.UIState.setWidth(a.ctx.root.Base().ID, a.ctx.Width)
	a.ctx.UIState.setHeight(a.ctx.root.Base().ID, a.ctx.Height)

	// --- Pass 0: Collect IDs (Top-Down) ---
	Visit(a.ctx.root, 0, nil, a.ctx, collectIdsVisitor, PreOrder)

	// --- Pass 0.5: Clean up state ---
	a.ctx.UIState.cleanup(a.ctx.ids)

	// --- Pass 1: Calculate Intrinsic Widths (Bottom-Up) ---
	Visit(a.ctx.root, 0, nil, a.ctx, calculateIntrinsicWidthVisitor, PostOrder)

	// --- Pass 2: Distribute Available Width (Top-Down) ---
	Visit(a.ctx.root, 0, nil, a.ctx, distributeAvailableWidthVisitor, PreOrder)

	// --- Pass 3: Perform text wrapping (Bottom-Up) ---
	// TODO: Implement text wrapping logic

	// --- Pass 4: Calculate Intrinsic Heights (Bottom-Up) ---
	Visit(a.ctx.root, 0, nil, a.ctx, calculateIntrinsicHeightVisitor, PostOrder)

	// --- Pass 5: Distribute Available Height (Top-Down) ---
	Visit(a.ctx.root, 0, nil, a.ctx, distributeAvailableHeightVisitor, PreOrder)
}

type VisitorFunc[T any] func(node Fc[T], index int, parent Fc[T], ctx *Context[T])

type Order int

const (
	PreOrder Order = iota
	PostOrder
)

func Visit[T any](node Fc[T], index int, parent Fc[T], ctx *Context[T], visitor VisitorFunc[T], order Order) {
	if node == nil {
		return
	}

	if order == PreOrder {
		visitor(node, index, parent, ctx)
	}

	for i, child := range node.Children(ctx) {
		Visit(child, i, node, ctx, visitor, order)
	}

	if order == PostOrder {
		visitor(node, index, parent, ctx)
	}
}

func calculateIntrinsicWidthVisitor[T any](node Fc[T], _ int, _ Fc[T], ctx *Context[T]) {
	if node == nil {
		return
	}
	if !node.Base().Opts.GrowX {
		renderResult := node.Render(ctx)
		width := lipgloss.Width(renderResult)
		ctx.UIState.setWidth(node.Base().ID, width)
	}
}

func calculateIntrinsicHeightVisitor[T any](node Fc[T], _ int, _ Fc[T], ctx *Context[T]) {
	if node == nil {
		return
	}
	if !node.Base().Opts.GrowY {
		renderResult := node.Render(ctx)
		height := lipgloss.Height(renderResult)
		ctx.UIState.setHeight(node.Base().ID, height)
	}
}

func distributeAvailableWidthVisitor[T any](node Fc[T], _ int, _ Fc[T], ctx *Context[T]) {
	if node == nil {
		return
	}
	children := node.Children(ctx)
	if len(children) == 0 {
		return
	}

	availableWidth := ctx.UIState.GetWidth(node.Base().ID)
	direction := node.Base().LayoutDirection

	if direction == Vertical {
		for _, child := range children {
			if child.Base().Opts.GrowX {
				ctx.UIState.setWidth(child.Base().ID, availableWidth)
			}
		}
	} else {
		nonGrowingChildrenWidth := 0
		growingChildrenCount := 0
		var growingChildren []Fc[T]

		for _, child := range children {
			if child.Base().Opts.GrowX {
				growingChildrenCount++
				growingChildren = append(growingChildren, child)
			} else {
				nonGrowingChildrenWidth += ctx.UIState.GetWidth(child.Base().ID)
			}
		}

		remainingWidth := availableWidth - nonGrowingChildrenWidth
		if remainingWidth < 0 {
			remainingWidth = 0
		}

		if growingChildrenCount > 0 {
			baseWidth := remainingWidth / growingChildrenCount
			remainder := remainingWidth % growingChildrenCount

			for i, child := range growingChildren {
				childWidth := baseWidth
				if i < remainder {
					childWidth++
				}
				ctx.UIState.setWidth(child.Base().ID, childWidth)
			}
		}
	}
}

func distributeAvailableHeightVisitor[T any](node Fc[T], _ int, _ Fc[T], ctx *Context[T]) {
	if node == nil {
		return
	}
	children := node.Children(ctx)
	if len(children) == 0 {
		return
	}

	availableHeight := ctx.UIState.GetHeight(node.Base().ID)
	direction := node.Base().LayoutDirection

	if direction == Horizontal {
		for _, child := range children {
			if child.Base().Opts.GrowY {
				ctx.UIState.setHeight(child.Base().ID, availableHeight)
			}
		}
	} else {
		nonGrowingChildrenHeight := 0
		growingChildrenCount := 0
		var growingChildren []Fc[T]

		for _, child := range children {
			if child.Base().Opts.GrowY {
				growingChildrenCount++
				growingChildren = append(growingChildren, child)
			} else {
				nonGrowingChildrenHeight += ctx.UIState.GetHeight(child.Base().ID)
			}
		}

		remainingHeight := availableHeight - nonGrowingChildrenHeight
		if remainingHeight < 0 {
			remainingHeight = 0
		}

		if growingChildrenCount > 0 {
			baseHeight := remainingHeight / growingChildrenCount
			remainder := remainingHeight % growingChildrenCount

			for i, child := range growingChildren {
				childHeight := baseHeight
				if i < remainder {
					childHeight++
				}
				ctx.UIState.setHeight(child.Base().ID, childHeight)
			}
		}
	}
}
