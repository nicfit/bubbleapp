package app

import (
	"github.com/alexanderbh/bubbleapp/style"
	zone "github.com/alexanderbh/bubblezone/v2"
)

type Context[T any] struct {
	Zone      *zone.Manager
	Styles    *style.Styles
	FocusedID string
	Width     int
	Height    int
	Data      *T
}
