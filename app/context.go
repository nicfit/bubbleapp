package app

import (
	"github.com/alexanderbh/bubbleapp/style"
	zone "github.com/alexanderbh/bubblezone/v2"
)

type Context struct {
	Zone         *zone.Manager
	Styles       *style.Styles
	FocusManager *FocusManager
	Width        int
	Height       int
}
