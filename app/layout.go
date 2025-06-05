package app

import (
	"reflect"

	"github.com/charmbracelet/lipgloss/v2"
)

type Layout struct {
	Direction LayoutDirection
	GrowX     bool
	GrowY     bool
	GapX      int
	GapY      int
	Width     int
	Height    int
}

type LayoutDirection int

const (
	Vertical LayoutDirection = iota
	Horizontal
)

type layoutPhase int

const (
	LayoutPhaseIntrincintWidth layoutPhase = iota
	LayoutPhaseIntrincintHeight
	LayoutPhaseAbsolutePositions
	LayoutPhaseFinalRender
)

type layoutManager struct {
	currentParent []*C
	width         int
	height        int
}

func newLayoutManager() *layoutManager {
	return &layoutManager{
		currentParent: nil,
	}
}

func (lm *layoutManager) addComponent(comp *C) {
	if len(lm.currentParent) > 0 {
		parent := lm.currentParent[len(lm.currentParent)-1]
		comp.parent = parent

		parent.children = append(parent.children, comp)
	}
	lm.currentParent = append(lm.currentParent, comp)
}

func (lm *layoutManager) pop(c *Ctx, comp *C) {

	if comp != nil {

		if c.LayoutPhase == LayoutPhaseIntrincintWidth {
			if comp.parent == nil {
				comp.width = lm.width
			} else if !comp.layout.GrowX {
				width := lipgloss.Width(comp.String())
				comp.width = width
			}
		}
		if c.LayoutPhase == LayoutPhaseIntrincintHeight {
			if comp.parent == nil {
				comp.height = lm.height
			} else if !comp.layout.GrowY {
				if comp.String() == "" {
					comp.height = 0
				} else {
					comp.height = lipgloss.Height(comp.String())
				}
			}
		}
	}

	if len(lm.currentParent) > 0 {
		lm.currentParent = lm.currentParent[:len(lm.currentParent)-1]
	}
}

func (lm *layoutManager) distributeWidth(c *Ctx) {
	Visit(c.root, 0, c, distributeAvailableWidthVisitor, PreOrder)
}
func (lm *layoutManager) distributeHeight(c *Ctx) {
	Visit(c.root, 0, c, distributeAvailableHeightVisitor, PreOrder)
}

func (lm *layoutManager) calculatePositions(c *Ctx) {
	Visit(c.root, 0, c, calculateAbsolutePositionVisitor, PreOrder)
}

type VisitorFunc func(node *C, index int, ctx *Ctx)

type Order int

const (
	PreOrder Order = iota
	PostOrder
)

func Visit(node *C, index int, ctx *Ctx, visitor VisitorFunc, order Order) {
	if node == nil {
		return
	}

	if order == PreOrder {
		visitor(node, index, ctx)
	}

	for i, child := range node.children {
		Visit(child, i, ctx, visitor, order)
	}

	if order == PostOrder {
		visitor(node, index, ctx)
	}
}

// extractLayoutFromProps attempts to extract layout information from a props interface.
// It prioritizes direct type assertion to Layout, then checks for an embedded Layout field,
// and finally falls back to individual field checks via reflection.
func extractLayoutFromProps(props interface{}) Layout {
	defaultLayout := Layout{
		Direction: Vertical,
		GrowX:     false,
		GrowY:     false,
		GapX:      0,
		GapY:      0,
	}

	if props == nil {
		return defaultLayout
	}

	val := reflect.ValueOf(props)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return defaultLayout
		}
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return defaultLayout
	}

	for i := range val.NumField() {
		field := val.Field(i)
		if field.Type() == reflect.TypeOf(Layout{}) {
			if layout, ok := field.Interface().(Layout); ok {
				mergedLayout := defaultLayout
				if layout.Direction != defaultLayout.Direction {
					mergedLayout.Direction = layout.Direction
				}
				mergedLayout.GrowX = layout.GrowX
				mergedLayout.GrowY = layout.GrowY
				mergedLayout.GapX = layout.GapX
				mergedLayout.GapY = layout.GapY
				return mergedLayout
			}
		}
	}

	return defaultLayout
}

func calculateAbsolutePositionVisitor(node *C, _ int, c *Ctx) {
	if node == nil {
		return
	}

	if node.parent == nil {
		node.x = 0
		node.y = 0
	} else {
		parent := node.parent
		prevSibling := (*C)(nil)
		childIndex := 0
		for i, child := range parent.children {
			if child == node {
				childIndex = i
				break
			}
		}

		if childIndex > 0 {
			prevSibling = parent.children[childIndex-1]
		}

		if parent.layout.Direction == Horizontal {
			node.y = parent.y
			if prevSibling == nil { // First child
				node.x = parent.x
			} else {
				node.x = prevSibling.x + prevSibling.width + parent.layout.GapX
			}
		} else { // Vertical
			node.x = parent.x
			if prevSibling == nil { // First child
				node.y = parent.y
			} else {
				node.y = prevSibling.y + prevSibling.height + parent.layout.GapY
			}
		}
	}
}

func distributeAvailableWidthVisitor(node *C, _ int, c *Ctx) {
	if node == nil {
		return
	}
	children := node.children
	if len(children) == 0 {
		return
	}

	availableWidth := node.width
	direction := node.layout.Direction

	if direction == Vertical {
		for _, child := range children {
			if child.layout.GrowX {
				child.width = availableWidth
			}
		}
	} else { // Horizontal
		nonGrowingChildrenWidth := 0
		growingChildrenCount := 0
		var growingChildren []*C

		for _, child := range children {
			if child.layout.GrowX {
				growingChildrenCount++
				growingChildren = append(growingChildren, child)
			} else {
				nonGrowingChildrenWidth += child.width
			}
		}

		totalGapWidth := 0
		// Calculate total gap width if there's more than one child and a positive gap is specified by the parent
		if len(children) > 1 && node.layout.GapX > 0 {
			totalGapWidth = (len(children) - 1) * node.layout.GapX
		}

		remainingWidth := availableWidth - nonGrowingChildrenWidth - totalGapWidth
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
				child.width = childWidth
			}
		}
	}
}

func distributeAvailableHeightVisitor(node *C, _ int, ctx *Ctx) {
	if node == nil {
		return
	}
	children := node.children
	if len(children) == 0 {
		return
	}

	availableHeight := node.height
	direction := node.layout.Direction

	if direction == Horizontal {
		for _, child := range children {
			if child.layout.GrowY {
				child.height = availableHeight
			}
		}
	} else { // Vertical
		nonGrowingChildrenHeight := 0
		growingChildrenCount := 0
		var growingChildren []*C

		for _, child := range children {
			if child.layout.GrowY {
				growingChildrenCount++
				growingChildren = append(growingChildren, child)
			} else {
				nonGrowingChildrenHeight += child.height
			}
		}

		totalGapHeight := 0
		if len(children) > 1 && node.layout.GapY > 0 {
			totalGapHeight = (len(children) - 1) * node.layout.GapY
		}

		remainingHeight := availableHeight - nonGrowingChildrenHeight - totalGapHeight
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
				child.height = childHeight
			}
		}
	}
}
