package app

import (
	"strconv"
	"strings"
)

type idContext[T any] struct {
	// IDs is a list of all IDs in tree
	ids []string
	// Maps ID to component
	idMap map[string]Fc[T]
	// The current path of the tree. Used to calculate IDs
	idPath []string
	// The number of times an ID has been used in the current path
	idPathCount map[string]int
}

func newIDContext[T any]() *idContext[T] {
	return &idContext[T]{
		ids:         []string{},
		idMap:       make(map[string]Fc[T]),
		idPath:      []string{},
		idPathCount: make(map[string]int),
	}
}

func (ctx *idContext[T]) getNode(id string) Fc[T] {
	if node, ok := ctx.idMap[id]; ok {
		return node
	}
	return nil
}

func (ctx *idContext[T]) initPath() {
	ctx.idPath = []string{"root"}
	ctx.idPathCount = make(map[string]int)
}

func (ctx *idContext[T]) initIDCollections() {
	ctx.idMap = make(map[string]Fc[T])
	ctx.ids = []string{}
}

// Used to get an ID when there are children further below.
// Remember to call PopID() when done.
func (ctx *idContext[T]) push(name string) string {
	path := strings.Join(ctx.idPath, "_")
	key := path + "_" + name
	index := ctx.idPathCount[key]
	ctx.idPathCount[key]++
	nameWithCount := name + "[" + strconv.Itoa(index) + "]"
	ctx.idPath = append(ctx.idPath, nameWithCount)
	return path + "_" + nameWithCount
}

func (ctx *idContext[T]) pop() {
	if len(ctx.idPath) == 0 {
		return
	}
	ctx.idPath = ctx.idPath[:len(ctx.idPath)-1]
}

func collectIdsVisitor[T any](node Fc[T], index int, parent Fc[T], ctx *Context[T]) {
	if node == nil {
		return
	}

	ctx.id.idMap[node.Base().ID] = node
	ctx.id.ids = append(ctx.id.ids, node.Base().ID)
}
