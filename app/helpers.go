package app

func RegisterMouse[T any](ctx *Context[T], ID string, node Fc[T], content string) string {
	ctx.ZoneMap[ID] = node
	return ctx.Zone.Mark(ID, content)
}
func Quit[T any](ctx *Context[T]) {
	ctx.Quit()
}
