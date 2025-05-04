package app

import (
	"image/color"

	"github.com/alexanderbh/bubbleapp/style"
	zone "github.com/alexanderbh/bubblezone/v2"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type Context[T any] struct {
	Zone            *zone.Manager
	ZoneMap         map[string]Fc[T]
	Cmds            *[]tea.Cmd
	Styles          *style.Styles
	Focused         Fc[T]
	BackgroundColor color.Color
	Width           int
	Height          int
	Data            *T
}

func NewContext[T any](data *T) *Context[T] {
	return &Context[T]{
		Zone:    zone.New(),
		ZoneMap: make(map[string]Fc[T]),
		Cmds:    &[]tea.Cmd{},
		Styles:  style.DefaultStyles(),
		Data:    data,
	}
}

func (ctx *Context[T]) Quit() {
	*ctx.Cmds = append(*ctx.Cmds, tea.Quit)
}
