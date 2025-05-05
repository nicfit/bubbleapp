package app

import (
	"strconv"
)

func generateIdsVisitor[T any](node Fc[T], index int, parent Fc[T], ctx *Context[T]) {
	if node == nil {
		return
	}

	parentID := "root"
	if parent != nil {
		parentID = parent.Base().ID
	}

	node.Base().ID = parentID + "_" + node.Base().TypeID + "[" + strconv.Itoa(index) + "]"
	ctx.IDMap[node.Base().ID] = node
	ctx.ids = append(ctx.ids, node.Base().ID)
}
