// Taken from the Bubble library.
// Needed to be modified to support mouse events

package table

import (
	"strconv"

	"github.com/alexanderbh/bubbleapp/app"

	"github.com/charmbracelet/bubbles/v2/help"
	"github.com/charmbracelet/bubbles/v2/key"
	"github.com/charmbracelet/bubbles/v2/viewport"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/mattn/go-runewidth"
)

type baseTable[T any] struct {
	base   *app.Base[T]
	ctx    *app.Context[T]
	KeyMap KeyMap
	Help   help.Model

	cols          []column
	rows          []Row
	data          func(*app.Context[T]) (clms []Column, rows []Row)
	cursor        int
	cursorChanged bool
	focus         bool
	rowHover      int
	styles        Styles

	viewport viewport.Model
	start    int
	end      int
}

type column struct {
	Title string
	Width int
}

type KeyMap struct {
	LineUp       key.Binding
	LineDown     key.Binding
	PageUp       key.Binding
	PageDown     key.Binding
	HalfPageUp   key.Binding
	HalfPageDown key.Binding
	GotoTop      key.Binding
	GotoBottom   key.Binding
}

func (km KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{km.LineUp, km.LineDown}
}

func (km KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{km.LineUp, km.LineDown, km.GotoTop, km.GotoBottom},
		{km.PageUp, km.PageDown, km.HalfPageUp, km.HalfPageDown},
	}
}

func defaultKeyMap() KeyMap {
	return KeyMap{
		LineUp: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "up"),
		),
		LineDown: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "down"),
		),
		PageUp: key.NewBinding(
			key.WithKeys("b", "pgup"),
			key.WithHelp("b/pgup", "page up"),
		),
		PageDown: key.NewBinding(
			key.WithKeys("f", "pgdown", "space"),
			key.WithHelp("f/pgdn", "page down"),
		),
		HalfPageUp: key.NewBinding(
			key.WithKeys("u", "ctrl+u"),
			key.WithHelp("u", "½ page up"),
		),
		HalfPageDown: key.NewBinding(
			key.WithKeys("d", "ctrl+d"),
			key.WithHelp("d", "½ page down"),
		),
		GotoTop: key.NewBinding(
			key.WithKeys("home", "g"),
			key.WithHelp("g/home", "go to start"),
		),
		GotoBottom: key.NewBinding(
			key.WithKeys("end", "G"),
			key.WithHelp("G/end", "go to end"),
		),
	}
}

type Styles struct {
	Header   lipgloss.Style
	Cell     lipgloss.Style
	Selected lipgloss.Style
	Hovered  lipgloss.Style
}

// DefaultStyles returns a set of default style definitions for this table.
func defaultStyles[T any](ctx *app.Context[T]) Styles {
	return Styles{
		Selected: lipgloss.NewStyle().Bold(true).Foreground(ctx.Styles.Colors.Tertiary),
		Hovered:  lipgloss.NewStyle().Bold(true).Foreground(ctx.Styles.Colors.Primary).Background(ctx.Styles.Colors.HighlightBackground),
		Header:   lipgloss.NewStyle().Bold(true).Padding(0, 1),
		Cell:     lipgloss.NewStyle().Padding(0, 1),
	}
}

func (m *baseTable[T]) SetStyles(s Styles) {
	m.styles = s
	m.updateViewport()
}

type Option[T any] func(*baseTable[T])

func newBaseTable[T any](ctx *app.Context[T], parentBase *app.Base[T], opts ...Option[T]) *baseTable[T] {
	m := baseTable[T]{
		base:     parentBase,
		ctx:      ctx,
		cursor:   -1,
		rowHover: -1,
		viewport: viewport.New(viewport.WithHeight(20)), //nolint:mnd

		KeyMap: defaultKeyMap(),
		Help:   help.New(),
		styles: defaultStyles(ctx),
	}

	for _, opt := range opts {
		opt(&m)
	}

	m.updateViewport()

	return &m
}

func withData[T any](data func(ctx *app.Context[T]) (clms []Column, rows []Row)) Option[T] {
	return func(m *baseTable[T]) {
		m.data = data
	}
}

// WithFocused sets the focus state of the table.
func withFocused[T any](f bool) Option[T] {
	return func(m *baseTable[T]) {
		m.focus = f
	}
}

// WithStyles sets the table styles.
func withStyles[T any](s Styles) Option[T] {
	return func(m *baseTable[T]) {
		m.styles = s
	}
}

// WithKeyMap sets the key map.
func withKeyMap[T any](km KeyMap) Option[T] {
	return func(m *baseTable[T]) {
		m.KeyMap = km
	}
}

func (m *baseTable[T]) setHover(n int) {
	if m.rowHover != n {
		m.rowHover = n
		m.updateViewport()
	}
}
func (m *baseTable[T]) clearHover() {
	if m.rowHover != -1 {
		m.rowHover = -1
		m.updateViewport()
	}
}

// update is the Bubble Tea update loop.
func (m *baseTable[T]) Update(ctx *app.Context[T], msg tea.Msg) {
	//cmds := []tea.Cmd{}

	switch msg := msg.(type) {
	// case tea.MouseMsg:
	// 	switch msg := msg.(type) {
	// 	case tea.MouseMotionMsg:
	// 		for i := range m.rows {
	// 			if ctx.Zone.Get(m.base.ID + "_" + strconv.Itoa(i)).InBounds(msg) {
	// 				m.setHover(i)
	// 				break
	// 			}
	// 			if i == len(m.rows)-1 {
	// 				m.clearHover()
	// 			}
	// 		}
	// 	case tea.MouseClickMsg:
	// 		if msg.Button == tea.MouseLeft {
	// 			for i := range m.rows {
	// 				if ctx.Zone.Get(m.base.ID + "_" + strconv.Itoa(i)).InBounds(msg) {
	// 					m.setFocus()
	// 					m.setCursor(i)
	// 					m.updateViewport()
	// 					cmds = append(cmds, func() tea.Msg {
	// 						return app.FocusComponentMsg{
	// 							TargetID: m.tableID,
	// 						}
	// 					})
	// 				}
	// 			}
	// 		}
	// 	}
	case tea.KeyPressMsg:
		if m.focus {
			switch {
			case key.Matches(msg, m.KeyMap.LineUp):
				m.MoveUp(1)
			case key.Matches(msg, m.KeyMap.LineDown):
				m.MoveDown(1)
			case key.Matches(msg, m.KeyMap.PageUp):
				m.MoveUp(m.viewport.Height())
			case key.Matches(msg, m.KeyMap.PageDown):
				m.MoveDown(m.viewport.Height())
			case key.Matches(msg, m.KeyMap.HalfPageUp):
				m.MoveUp(m.viewport.Height() / 2) //nolint:mnd
			case key.Matches(msg, m.KeyMap.HalfPageDown):
				m.MoveDown(m.viewport.Height() / 2) //nolint:mnd
			case key.Matches(msg, m.KeyMap.GotoTop):
				m.GotoTop()
			case key.Matches(msg, m.KeyMap.GotoBottom):
				m.GotoBottom()
			}
		}
	}

	// if m.cursorChanged && m.cursor >= 0 && m.cursor < len(m.rows) {
	// 	m.cursorChanged = false
	// 	cmds = append(cmds, func() tea.Msg {
	// 		return RowFocusMsg{
	// 			TableID: m.tableID,
	// 			Row:     m.cursor,
	// 			RowID:   m.rows[m.cursor][0], // Column 0 is the ID for now
	// 		}
	// 	})
	// }

}

// Focus focuses the table, allowing the user to move around the rows and
// interact.
func (m *baseTable[T]) setFocus() {
	m.focus = true
	if m.cursor < 0 {
		m.setCursor(0)
	}
	if m.cursor >= len(m.rows) {
		m.setCursor(len(m.rows) - 1)
	}
	m.updateViewport()
}

// blur blurs the table, preventing selection or movement.
func (m *baseTable[T]) blur() {
	m.focus = false
	m.updateViewport()
}

// view renders the component.
func (m *baseTable[T]) Render(ctx *app.Context[T]) string {
	m.updateViewport()
	headersView := m.headersView()
	m.viewport.SetHeight(m.base.Height - lipgloss.Height(headersView))
	m.viewport.SetWidth(m.base.Width)
	return headersView + "\n" + app.RegisterMouse(m.ctx, m.base.ID+"_body", m, m.viewport.View())
}

// helpView is a helper method for rendering the help menu from the keymap.
// Note that this view is not rendered by default and you must call it
// manually in your application, where applicable.
func (m baseTable[T]) helpView() string {
	return m.Help.View(m.KeyMap)
}

// updateViewport updates the list content based on the previously defined
// columns and rows.
func (m *baseTable[T]) updateViewport() {
	if m.data != nil {
		var rawCols []Column
		rawCols, m.rows = m.data(m.ctx)
		m.cols = columnMapping(m.base.Width, rawCols)
	}
	renderedRows := make([]string, 0, len(m.rows))

	// Render only rows from: m.cursor-m.viewport.Height to: m.cursor+m.viewport.Height
	// Constant runtime, independent of number of rows in a table.
	// Limits the number of renderedRows to a maximum of 2*m.viewport.Height
	if m.cursor >= 0 {
		m.start = clamp(m.cursor-m.viewport.Height(), 0, m.cursor)
	} else {
		m.start = 0
	}
	m.end = clamp(max(0, m.cursor)+m.viewport.Height(), m.cursor, len(m.rows))
	for i := m.start; i < m.end; i++ {
		renderedRows = append(renderedRows, m.renderRow(i))
	}

	m.viewport.SetContent(
		lipgloss.JoinVertical(lipgloss.Left, renderedRows...),
	)
}

// selectedRow returns the selected row.
// You can cast it to your own implementation.
func (m baseTable[T]) selectedRow() Row {
	if m.cursor < 0 || m.cursor >= len(m.rows) {
		return nil
	}

	return m.rows[m.cursor]
}

// SetWidth sets the width of the viewport of the table.
func (m *baseTable[T]) setWidth(w int) {
	m.viewport.SetWidth(w)
	m.updateViewport()
}

// SetHeight sets the height of the viewport of the table.
func (m *baseTable[T]) setHeight(h int) {
	m.viewport.SetHeight(h - lipgloss.Height(m.headersView()))
	m.updateViewport()
}

// Height returns the viewport height of the table.
func (m baseTable[T]) height() int {
	return m.viewport.Height()
}

// Width returns the viewport width of the table.
func (m baseTable[T]) width() int {
	return m.viewport.Width()
}

// Cursor returns the index of the selected row.
func (m baseTable[T]) getCursor() int {
	return m.cursor
}

// SetCursor sets the cursor position in the table.
func (m *baseTable[T]) setCursor(n int) {
	m.cursor = clamp(n, 0, len(m.rows)-1)
	m.cursorChanged = true
}

// MoveUp moves the selection up by any number of rows.
// It can not go above the first row.
func (m *baseTable[T]) MoveUp(n int) {
	m.setCursor(m.cursor - n)
	switch {
	case m.start == 0:
		m.viewport.SetYOffset(clamp(m.viewport.YOffset, 0, m.cursor))
	case m.start < m.viewport.Height():
		m.viewport.YOffset = (clamp(clamp(m.viewport.YOffset+n, 0, m.cursor), 0, m.viewport.Height()))
	case m.viewport.YOffset >= 1:
		m.viewport.YOffset = clamp(m.viewport.YOffset+n, 1, m.viewport.Height())
	}
	m.updateViewport()
}

// MoveDown moves the selection down by any number of rows.
// It can not go below the last row.
func (m *baseTable[T]) MoveDown(n int) {
	m.setCursor(m.cursor + n)
	m.updateViewport()

	switch {
	case m.end == len(m.rows) && m.viewport.YOffset > 0:
		m.viewport.SetYOffset(clamp(m.viewport.YOffset-n, 1, m.viewport.Height()))
	case m.cursor > (m.end-m.start)/2 && m.viewport.YOffset > 0:
		m.viewport.SetYOffset(clamp(m.viewport.YOffset-n, 1, m.cursor))
	case m.viewport.YOffset > 1:
	case m.cursor > m.viewport.YOffset+m.viewport.Height()-1:
		m.viewport.SetYOffset(clamp(m.viewport.YOffset+1, 0, 1))
	}
}

// GotoTop moves the selection to the first row.
func (m *baseTable[T]) GotoTop() {
	m.MoveUp(m.cursor)
}

// GotoBottom moves the selection to the last row.
func (m *baseTable[T]) GotoBottom() {
	m.MoveDown(len(m.rows))
}

func (m baseTable[T]) headersView() string {
	s := make([]string, 0, len(m.cols))
	for _, col := range m.cols {
		if col.Width <= 0 {
			continue
		}
		style := lipgloss.NewStyle().Width(col.Width).MaxWidth(col.Width).Inline(true)
		renderedCell := style.Render(runewidth.Truncate(col.Title, col.Width, "…"))
		s = append(s, m.styles.Header.Render(renderedCell))
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, s...)
}

func (m *baseTable[T]) renderRow(r int) string {
	s := make([]string, 0, len(m.cols))
	for i, value := range m.rows[r] {
		if m.cols[i].Width <= 0 {
			continue
		}
		style := lipgloss.NewStyle().Width(m.cols[i].Width).MaxWidth(m.cols[i].Width).Inline(true)
		renderedCell := m.styles.Cell.Render(style.Render(runewidth.Truncate(value, m.cols[i].Width, "…")))
		s = append(s, renderedCell)
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, s...)

	if r == m.cursor {
		return app.RegisterMouse(m.ctx, m.base.ID+"_"+strconv.Itoa(r), m, m.styles.Selected.Render(row))
	}
	if r == m.rowHover {
		return app.RegisterMouse(m.ctx, m.base.ID+"_"+strconv.Itoa(r), m, m.styles.Hovered.Render(row))
	}

	return app.RegisterMouse(m.ctx, m.base.ID+"_"+strconv.Itoa(r), m, row)
}

func clamp(v, low, high int) int {
	return min(max(v, low), high)
}

func (m *baseTable[T]) Children(ctx *app.Context[T]) []app.Fc[T] {
	return nil
}
func (m *baseTable[T]) Base() *app.Base[T] {
	return m.base
}
