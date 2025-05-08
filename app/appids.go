package app

import (
	"strconv"
	"strings"
)

type fcIDContext struct {
	// IDs is a list of all IDs in tree
	ids []string
	// The current path of the tree. Used to calculate IDs
	idPath []string
	// The number of times an ID has been used in the current path
	idPathCount map[string]int
}

func newFCIDContext() *fcIDContext {
	return &fcIDContext{
		ids:         []string{},
		idPath:      []string{},
		idPathCount: make(map[string]int),
	}
}

func (ctx *fcIDContext) initPath() {
	ctx.idPath = []string{}
	ctx.idPathCount = make(map[string]int)
}

func (ctx *fcIDContext) initIDCollections() {
	ctx.ids = []string{}
}

// Used to get an ID when there are children further below.
// Remember to call PopID() when done.
func (ctx *fcIDContext) push(name string) string {
	path := strings.Join(ctx.idPath, "_")
	key := path + "_" + name
	index := ctx.idPathCount[key]
	ctx.idPathCount[key]++
	nameWithCount := name + "[" + strconv.Itoa(index) + "]"
	ctx.idPath = append(ctx.idPath, nameWithCount)
	return path + "_" + nameWithCount
}

func (ctx *fcIDContext) pop() {
	if len(ctx.idPath) == 0 {
		return
	}
	ctx.idPath = ctx.idPath[:len(ctx.idPath)-1]
}

func (ctx *fcIDContext) getID() string {
	if len(ctx.idPath) == 0 {
		return "root"
	}
	return strings.Join(ctx.idPath, "_")
}
