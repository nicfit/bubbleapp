package app

func collectIdsVisitor[T any](node Fc[T], index int, parent Fc[T], ctx *Context[T]) {
	if node == nil {
		return
	}

	ctx.IDMap[node.Base().ID] = node
	ctx.ids = append(ctx.ids, node.Base().ID)
}
