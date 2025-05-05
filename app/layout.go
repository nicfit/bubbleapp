package app

import (
	"github.com/charmbracelet/lipgloss/v2"
)

// Calculate sizes of the component tree
func (a *App[T]) Layout() {
	a.ctx.LayoutPhase = true
	defer func() {
		a.ctx.LayoutPhase = false
	}()
	// TODO: Can the Zone manager be reset here? If not why? Otherwise things will live in the zone forever.
	a.root.Base().Width = a.ctx.Width
	a.root.Base().Height = a.ctx.Height

	// --- Pass 1: Calculate Intrinsic Widths (Bottom-Up) ---
	Visit(a.root, nil, a.ctx, calculateIntrinsicWidthVisitor, PostOrder)

	// --- Pass 2: Distribute Available Width (Top-Down) ---
	Visit(a.root, nil, a.ctx, distributeAvailableWidthVisitor, PreOrder)

	// --- Pass 3: Perform text wrapping (Bottom-Up) ---
	// TODO: Implement text wrapping logic

	// --- Pass 4: Calculate Intrinsic Heights (Bottom-Up) ---
	Visit(a.root, nil, a.ctx, calculateIntrinsicHeightVisitor, PostOrder)

	// --- Pass 5: Distribute Available Height (Top-Down) ---
	Visit(a.root, nil, a.ctx, distributeAvailableHeightVisitor, PreOrder)
}

type VisitorFunc[T any] func(node Fc[T], parent Fc[T], ctx *Context[T])

type Order int

const (
	PreOrder Order = iota
	PostOrder
)

func Visit[T any](node Fc[T], parent Fc[T], ctx *Context[T], visitor VisitorFunc[T], order Order) {
	if node == nil {
		return
	}

	if order == PreOrder {
		visitor(node, parent, ctx)
	}

	for _, child := range node.Children(ctx) {
		Visit(child, node, ctx, visitor, order)
	}

	if order == PostOrder {
		visitor(node, parent, ctx)
	}
}

func calculateIntrinsicWidthVisitor[T any](node Fc[T], parent Fc[T], ctx *Context[T]) {
	if node == nil {
		return
	}
	if !node.Base().Opts.GrowX {
		renderResult := node.Render(ctx)
		node.Base().Width = lipgloss.Width(renderResult)
	}
}

func calculateIntrinsicHeightVisitor[T any](node Fc[T], parent Fc[T], ctx *Context[T]) {
	if node == nil {
		return
	}
	if !node.Base().Opts.GrowY {
		renderResult := node.Render(ctx)
		node.Base().Height = lipgloss.Height(renderResult)
	}
}

func distributeAvailableWidthVisitor[T any](node Fc[T], parent Fc[T], ctx *Context[T]) {
	if node == nil {
		return
	}
	children := node.Children(ctx)
	if len(children) == 0 {
		return
	}

	availableWidth := node.Base().Width
	direction := node.Base().LayoutDirection

	if direction == Vertical {
		for _, child := range children {
			if child.Base().Opts.GrowX {
				child.Base().Width = availableWidth
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
				nonGrowingChildrenWidth += child.Base().Width
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
				child.Base().Width = childWidth
			}
		}
	}
}

func distributeAvailableHeightVisitor[T any](node Fc[T], parent Fc[T], ctx *Context[T]) {
	if node == nil {
		return
	}
	children := node.Children(ctx)
	if len(children) == 0 {
		return
	}

	availableHeight := node.Base().Height
	direction := node.Base().LayoutDirection

	if direction == Horizontal {
		for _, child := range children {
			if child.Base().Opts.GrowY {
				child.Base().Height = availableHeight
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
				nonGrowingChildrenHeight += child.Base().Height
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
				child.Base().Height = childHeight
			}
		}
	}
}
