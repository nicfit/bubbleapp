package app

func RegisterMouse[T any](ctx *Context[T], ID string, node Fc[T], content string) string {
	ctx.ZoneMap[ID] = node
	return ctx.Zone.Mark(ID, content)
}

// Helper function that can be used as an handler function
func Quit[T any](ctx *Context[T]) {
	ctx.Quit()
}
