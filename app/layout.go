package app

import (
	"reflect"
)

type Layout struct {
	Direction LayoutDirection
	GrowX     bool
	GrowY     bool
	GapX      int
	GapY      int
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
	LayoutPhaseFinalRender
)

type componentTree struct {
	nodes map[string]*ComponentNode
	root  *ComponentNode
}

func newComponentTree() *componentTree {
	return &componentTree{
		nodes: make(map[string]*ComponentNode),
		root:  nil,
	}
}

type layoutManager struct {
	componentTree *componentTree
	currentParent []*ComponentNode
	width         int
	height        int
}

func newLayoutManager() *layoutManager {
	return &layoutManager{
		componentTree: newComponentTree(),
		currentParent: nil,
	}
}

type ComponentNode struct {
	ID         string
	Parent     *ComponentNode
	LastRender string
	Props      Props            // Props passed to this instance
	Children   []*ComponentNode // Children in rendered order
	Layout     Layout
}

func (lm *layoutManager) addComponent(id string, props Props) *ComponentNode {

	node := &ComponentNode{
		ID:     id,
		Props:  props,
		Layout: extractLayoutFromProps(props),
	}

	if len(lm.currentParent) > 0 {
		parent := lm.currentParent[len(lm.currentParent)-1]
		node.Parent = parent

		parent.Children = append(parent.Children, node)
	}
	lm.currentParent = append(lm.currentParent, node)

	if lm.componentTree.root == nil {
		lm.componentTree.root = node
	}

	lm.componentTree.nodes[id] = node

	return node
}

func (lm *layoutManager) pop() {
	if len(lm.currentParent) > 0 {
		lm.currentParent = lm.currentParent[:len(lm.currentParent)-1]
	}
}

func (lm *layoutManager) getComponent(id string) *ComponentNode {
	node, ok := lm.componentTree.nodes[id]
	if !ok {
		return nil
	}
	return node
}

func (lm *layoutManager) distributeWidth(c *Ctx) {
	Visit(lm.componentTree.root, 0, c, distributeAvailableWidthVisitor, PreOrder)
}
func (lm *layoutManager) distributeHeight(c *Ctx) {
	Visit(lm.componentTree.root, 0, c, distributeAvailableHeightVisitor, PreOrder)
}

type VisitorFunc func(node *ComponentNode, index int, ctx *Ctx)

type Order int

const (
	PreOrder Order = iota
	PostOrder
)

func Visit(node *ComponentNode, index int, ctx *Ctx, visitor VisitorFunc, order Order) {
	if node == nil {
		return
	}

	if order == PreOrder {
		visitor(node, index, ctx)
	}

	for i, child := range node.Children {
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

func distributeAvailableWidthVisitor(node *ComponentNode, _ int, c *Ctx) {
	if node == nil {
		return
	}
	children := node.Children
	if len(children) == 0 {
		return
	}

	availableWidth := c.UIState.GetWidth(node.ID)
	direction := node.Layout.Direction

	if direction == Vertical {
		for _, child := range children {
			if child.Layout.GrowX {
				c.UIState.setWidth(child.ID, availableWidth)
			}
		}
	} else { // Horizontal
		nonGrowingChildrenWidth := 0
		growingChildrenCount := 0
		var growingChildren []*ComponentNode

		for _, child := range children {
			if child.Layout.GrowX {
				growingChildrenCount++
				growingChildren = append(growingChildren, child)
			} else {
				nonGrowingChildrenWidth += c.UIState.GetWidth(child.ID)
			}
		}

		totalGapWidth := 0
		// Calculate total gap width if there's more than one child and a positive gap is specified by the parent
		if len(children) > 1 && node.Layout.GapX > 0 {
			totalGapWidth = (len(children) - 1) * node.Layout.GapX
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
				c.UIState.setWidth(child.ID, childWidth)
			}
		}
	}
}

func distributeAvailableHeightVisitor(node *ComponentNode, _ int, ctx *Ctx) {
	if node == nil {
		return
	}
	children := node.Children
	if len(children) == 0 {
		return
	}

	availableHeight := ctx.UIState.GetHeight(node.ID)
	direction := node.Layout.Direction

	if direction == Horizontal {
		for _, child := range children {
			if child.Layout.GrowY {
				ctx.UIState.setHeight(child.ID, availableHeight)
			}
		}
	} else { // Vertical
		nonGrowingChildrenHeight := 0
		growingChildrenCount := 0
		var growingChildren []*ComponentNode

		for _, child := range children {
			if child.Layout.GrowY {
				growingChildrenCount++
				growingChildren = append(growingChildren, child)
			} else {
				nonGrowingChildrenHeight += ctx.UIState.GetHeight(child.ID)
			}
		}

		totalGapHeight := 0
		if len(children) > 1 && node.Layout.GapY > 0 {
			totalGapHeight = (len(children) - 1) * node.Layout.GapY
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
				ctx.UIState.setHeight(child.ID, childHeight)
			}
		}
	}
}
