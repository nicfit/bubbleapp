package app

import (
	"strconv"
	"strings"
)

// idContext stores the context for generating unique IDs for components in a tree structure.
// It maintains a list of IDs, the current path in the tree, and a count of how many times
// each ID has been used in the current path.
type idContext struct {
	// The current path of the tree. Used to calculate IDs
	idPath []string
	// The number of times an ID has been used in the current path
	idPathCount map[string]int
}

func newIdContext() *idContext {
	return &idContext{
		idPath:      []string{},
		idPathCount: make(map[string]int),
	}
}

func (ctx *idContext) initPath() {
	ctx.idPath = []string{}
	ctx.idPathCount = make(map[string]int)
}

// Used to get an ID when there are children further below.
// Remember to call PopID() when done.
func (ctx *idContext) push(name string) string {
	// Create a key for idPathCount to ensure uniqueness of counts
	// based on the current position in the hierarchy.
	parentPathString := strings.Join(ctx.idPath, "_")
	countKey := name // Default for root elements
	if parentPathString != "" {
		countKey = parentPathString + "_" + name
	}

	index := ctx.idPathCount[countKey]
	ctx.idPathCount[countKey]++

	// The segment added to idPath should be simple: name + [index]
	currentSegment := name + "[" + strconv.Itoa(index) + "]"
	ctx.idPath = append(ctx.idPath, currentSegment)

	// The ID for the component is the full path.
	return strings.Join(ctx.idPath, "_")
}

func (ctx *idContext) pop() {
	if len(ctx.idPath) == 0 {
		return
	}
	ctx.idPath = ctx.idPath[:len(ctx.idPath)-1]
}

func (ctx *idContext) getID() string {
	return strings.Join(ctx.idPath, "_")
}
