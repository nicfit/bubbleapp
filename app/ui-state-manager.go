package app

type uiStateContext struct {
	Focused      string
	Hovered      string
	HoveredChild string
}

func NewUIStateContext() *uiStateContext {
	return &uiStateContext{}
}
