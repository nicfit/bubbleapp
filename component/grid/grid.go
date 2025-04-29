package grid

import (
	"math"

	"github.com/alexanderbh/bubbleapp/app"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

// THIS IS NOT DONE. Work in progress. Made a stack first since this became too
// complicated until I knew the approach I wanted to take.

const (
	breakpointSm = 60
	breakpointMd = 90
	breakpointLg = 120
)

type GridItemConfig struct {
	item app.UIModel
	Xs   int
	Sm   int
	Md   int
	Lg   int
}

func (c GridItemConfig) GetSpanForWidth(width int) int {
	var span int // Declare span without initial value
	switch {
	case width >= breakpointLg && c.Lg > 0:
		span = c.Lg
	case width >= breakpointMd && c.Md > 0:
		span = c.Md
	case width >= breakpointSm && c.Sm > 0:
		span = c.Sm
	case c.Xs > 0:
		span = c.Xs
	default:
		span = 12 // Assign default value here
	}

	if span < 1 {
		return 1
	}
	if span > 12 {
		return 12
	}
	return span
}

type GridItemConfigOption func(*GridItemConfig)

func WithXs(xs int) GridItemConfigOption {
	return func(c *GridItemConfig) {
		c.Xs = xs
	}
}
func WithSm(sm int) GridItemConfigOption {
	return func(c *GridItemConfig) {
		c.Sm = sm
	}
}
func WithMd(md int) GridItemConfigOption {
	return func(c *GridItemConfig) {
		c.Md = md
	}
}
func WithLg(lg int) GridItemConfigOption {
	return func(c *GridItemConfig) {
		c.Lg = lg
	}
}
func NewItem(item app.UIModel, opts ...GridItemConfigOption) GridItemConfig {
	config := GridItemConfig{
		item: item,
		Xs:   12,
		Sm:   12,
		Md:   12,
		Lg:   12,
	}
	for _, opt := range opts {
		opt(&config)
	}
	return config
}

type model struct {
	base        *app.Base
	itemConfigs map[string]GridItemConfig
}

func New(ctx *app.Context, items []app.UIModel) model {
	return model{
		base:        app.New(ctx, app.WithGrowX(true), app.WithGrowY(true)),
		itemConfigs: make(map[string]GridItemConfig),
	}
}

func (m model) Init() tea.Cmd {
	return m.base.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	cmd := m.base.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	childrenSlice := m.base.GetChildren()

	if len(childrenSlice) == 0 {
		return ""
	}

	containerWidth := m.base.Ctx.Width

	var rows [][]string
	var currentRowItems []string
	currentRowSpan := 0

	widthPerSpanUnit := float64(containerWidth) / 12.0

	// Iterate directly over children from base
	for _, child := range childrenSlice {
		childID := child.Base().ID

		config, ok := m.itemConfigs[childID]
		if !ok {
			config = GridItemConfig{Xs: 12}
		}

		childSpan := config.GetSpanForWidth(containerWidth)
		childView := child.View()

		targetChildWidth := int(math.Floor(widthPerSpanUnit * float64(childSpan)))
		if targetChildWidth < 1 {
			targetChildWidth = 1
		}

		styledChildView := lipgloss.NewStyle().Width(targetChildWidth).MaxWidth(targetChildWidth).Render(childView)

		if currentRowSpan+childSpan > 12 {
			if len(currentRowItems) > 0 {
				rows = append(rows, currentRowItems)
			}
			currentRowItems = []string{styledChildView}
			currentRowSpan = childSpan
		} else {
			currentRowItems = append(currentRowItems, styledChildView)
			currentRowSpan += childSpan
		}
	}

	if len(currentRowItems) > 0 {
		rows = append(rows, currentRowItems)
	}

	renderedRows := make([]string, len(rows))
	for i, rowItems := range rows {
		renderedRows[i] = lipgloss.JoinHorizontal(lipgloss.Top, rowItems...)
	}

	return lipgloss.JoinVertical(lipgloss.Left, renderedRows...)
}

func (m model) Base() *app.Base {
	return m.base
}
