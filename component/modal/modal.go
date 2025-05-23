package modal

import (
	"github.com/alexanderbh/bubbleapp/app"
)

type Props struct {
	Child app.FC
}

type prop func(*Props)

func modal(c *app.Ctx, rawProps app.Props) string {
	_, ok := rawProps.(Props)
	if !ok {
		return ""
	}

	return ""
}

// New creates a new text element.
func New(c *app.Ctx, Child app.FC, opts ...prop) *app.C {
	p := Props{
		Child: Child,
	}

	for _, opt := range opts {
		if opt != nil {
			opt(&p)
		}
	}
	return c.Render(modal, p)
}
