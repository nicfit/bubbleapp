package app

import (
	"image/color"
	"strconv"
	"strings"

	"github.com/alexanderbh/bubbleapp/style"
	zone "github.com/alexanderbh/bubblezone/v2"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type Context[T any] struct {
	root            Fc[T]
	IDMap           map[string]Fc[T]
	ids             []string
	UIState         *StateStore
	Zone            *zone.Manager
	ZoneMap         map[string]Fc[T]
	Cmds            *[]tea.Cmd
	Styles          *style.Styles
	BackgroundColor color.Color
	Width           int
	Height          int
	LayoutPhase     bool
	Data            *T

	idPath      []string
	idPathCount map[string]int
}

func NewContext[T any](data *T) *Context[T] {
	return &Context[T]{
		Zone:    zone.New(),
		ZoneMap: make(map[string]Fc[T]),
		IDMap:   make(map[string]Fc[T]),
		ids:     []string{},
		UIState: NewStateStore(),
		Cmds:    &[]tea.Cmd{},
		Styles:  style.DefaultStyles(),
		Data:    data,
	}
}

// Used to get an ID when there are children further below.
// Remember to call PopID() when done.
func (ctx *Context[T]) PushID(name string) string {
	path := strings.Join(ctx.idPath, "_")
	key := path + "_" + name
	index := ctx.idPathCount[key]
	ctx.idPathCount[key]++
	nameWithCount := name + "[" + strconv.Itoa(index) + "]"
	ctx.idPath = append(ctx.idPath, nameWithCount)
	return path + "_" + nameWithCount
}

// Used to get a leaf node ID
func (ctx *Context[T]) GetID(name string) string {
	path := strings.Join(ctx.idPath, "_")
	path = path + "_" + name
	id := path + "[" + strconv.Itoa(ctx.idPathCount[path]) + "]"
	ctx.idPathCount[path]++
	return id
}

func (ctx *Context[T]) PopID() {
	if len(ctx.idPath) == 0 {
		return
	}
	ctx.idPath = ctx.idPath[:len(ctx.idPath)-1]
}

func (ctx *Context[T]) Quit() {
	*ctx.Cmds = append(*ctx.Cmds, tea.Quit)
}
