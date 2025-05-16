package tabs

import (
	"strconv"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/box"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/tabtitles"
)

type Tab struct {
	Title   string
	Content app.FC
}

type Props struct {
	Tabs []Tab
	app.Layout
}

type Prop func(*Props)

func New(c *app.Ctx, ts []Tab, prop ...Prop) app.C {
	p := Props{
		Tabs: ts,
		Layout: app.Layout{
			GrowX: true,
			GrowY: true,
		},
	}
	for _, t := range prop {
		t(&p)
	}
	return c.Render(Tabs, p)
}

func Tabs(c *app.Ctx, props app.Props) string {
	p, _ := props.(Props) // Use type assertion

	if p.Tabs == nil {
		return ""
	}

	activeTab, setActiveTab := app.UseState(c, 0)

	// TODO: UseMemo for this
	titles := make([]string, len(p.Tabs))
	for i, t := range p.Tabs {
		titles[i] = t.Title
	}

	return stack.New(c, func(c *app.Ctx) []app.C {
		return []app.C{
			tabtitles.New(c, titles, activeTab, func(activeTab int) {
				setActiveTab(activeTab)
			}),
			box.New(c, func(c *app.Ctx) app.C {
				return p.Tabs[activeTab].Content(c)
			}, box.WithKey(strconv.Itoa(activeTab)), box.WithDisableFollow(true)),
		}
	}).String()

}
