// Taken from the Bubble library.

package table

import (
	"strconv"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/style"

	"github.com/charmbracelet/bubbles/v2/help"
	"github.com/charmbracelet/bubbles/v2/key"
	"github.com/charmbracelet/bubbles/v2/viewport"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/mattn/go-runewidth"
)

type Options struct {
	style.Margin
}

type Row []string

type ColumnWidth struct {
	Int  int
	Grow bool
}

type Column struct {
	Title string
	Width ColumnWidth
}

func WidthGrow() ColumnWidth {
	return ColumnWidth{Grow: true}
}
func WidthInt(i int) ColumnWidth {
	return ColumnWidth{Int: i}
}

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
	Base      lipgloss.Style
	BaseFocus lipgloss.Style
	Header    lipgloss.Style
	Cell      lipgloss.Style
	Selected  lipgloss.Style
	Hovered   lipgloss.Style
}

// DefaultStyles returns a set of default style definitions for this table.
func defaultStyles[T any](ctx *app.Context[T]) Styles {
	base := lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true, true, true, true).BorderForeground(ctx.Styles.Colors.Ghost)
	return Styles{
		Base:      base,
		BaseFocus: base.BorderForeground(ctx.Styles.Colors.White),
		Selected:  lipgloss.NewStyle().Bold(true).Foreground(ctx.Styles.Colors.PrimaryLight).Background(ctx.Styles.Colors.UIPanelBackground),
		Hovered:   lipgloss.NewStyle().Bold(true).Foreground(ctx.Styles.Colors.PrimaryLight).Background(ctx.Styles.Colors.HighlightBackground),
		Header:    lipgloss.NewStyle().Bold(true).Padding(0, 1).BorderStyle(lipgloss.NormalBorder()).BorderForeground(ctx.Styles.Colors.GhostLight).BorderBottom(true),
		Cell:      lipgloss.NewStyle().Padding(0, 1),
	}
}

type Option[T any] func(*baseTable[T])

func New[T any](ctx *app.Context[T], data func(ctx *app.Context[T]) (clms []Column, rows []Row), options *Options, baseOptions ...app.BaseOption) *baseTable[T] {
	if options == nil {
		options = &Options{}
	}
	if baseOptions == nil {
		baseOptions = []app.BaseOption{}
	}
	base := app.NewBase[T](append([]app.BaseOption{app.WithFocusable(true), app.WithGrow(true)}, baseOptions...)...)

	s := defaultStyles(ctx)

	s.Base = style.ApplyMargin(s.Base, options.Margin)
	s.BaseFocus = style.ApplyMargin(s.BaseFocus, options.Margin)

	m := baseTable[T]{
		base:     base,
		ctx:      ctx,
		data:     data,
		cursor:   -1,
		rowHover: -1,
		viewport: viewport.New(),

		KeyMap: defaultKeyMap(),
		Help:   help.New(),
		styles: s,
	}

	m.updateViewport()

	return &m
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
		if m.ctx.Focused == m {
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

// view renders the component.
func (m *baseTable[T]) Render(ctx *app.Context[T]) string {
	s := m.getBaseStyle(ctx)
	if m.data != nil {
		var rawCols []Column
		rawCols, m.rows = m.data(ctx)
		m.cols = columnMapping(m.base.Width-s.GetHorizontalFrameSize()-(m.styles.Header.GetHorizontalFrameSize()*len(rawCols)), rawCols)
	}
	if ctx.Focused == m {
		if m.cursor < 0 {
			m.setCursor(0)
		}
	}
	if m.cursor >= len(m.rows) {
		m.setCursor(len(m.rows) - 1)
	}
	m.updateViewport()
	headersView := m.headersView()
	m.viewport.SetHeight(m.base.Height - lipgloss.Height(headersView) - s.GetVerticalFrameSize())
	m.viewport.SetWidth(m.base.Width - s.GetHorizontalFrameSize())

	return s.Render(headersView + "\n" + app.RegisterMouse(ctx, m.base.ID+"_body", m, m.viewport.View()))
}

func (m *baseTable[T]) getBaseStyle(ctx *app.Context[T]) lipgloss.Style {
	if ctx.Focused == m {
		return m.styles.BaseFocus
	}
	return m.styles.Base
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

func columnMapping(width int, clms []Column) []column {
	numberOfGrowers := 0
	sizeOfStatic := 0
	for _, clm := range clms {
		if clm.Width.Grow {
			numberOfGrowers++
		} else {
			sizeOfStatic += clm.Width.Int
		}
	}

	growWidth := width - sizeOfStatic
	baseSizePerGrower := 0
	remainder := 0

	if numberOfGrowers > 0 {
		if growWidth < 0 {
			growWidth = 0
		}
		baseSizePerGrower = growWidth / numberOfGrowers
		remainder = growWidth % numberOfGrowers
	}

	columns := []column{}
	for _, clm := range clms {
		colWidth := 0
		if clm.Width.Grow {
			colWidth = baseSizePerGrower
			if remainder > 0 {
				colWidth++
				remainder--
			}
		} else {
			colWidth = clm.Width.Int
		}

		if colWidth < 0 {
			colWidth = 0
		}

		columns = append(columns, column{
			Title: clm.Title,
			Width: colWidth,
		})
	}

	return columns
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
