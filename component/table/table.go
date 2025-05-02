package table

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/style"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

// This is a bit of a mess right now, but it works okay.
// I should probably just have made it 1 component instead of trying
// to wrap the Bubble table.

type RowFocusMsg struct {
	TableID string
	Row     int
	RowID   string
}

type RowActionMsg struct {
	Key     tea.KeyMsg
	TableID string
	Row     int
	RowID   string
}

// The first element in the row is considered the ID
// Maybe add a toggle to show/hide the first column (the ID)
type Row []string

type ColumnWidth struct {
	Int  int
	Grow bool
}

// TODO: Change to options struct like other components
func WidthGrow() ColumnWidth {
	return ColumnWidth{Grow: true}
}
func WidthInt(i int) ColumnWidth {
	return ColumnWidth{Int: i}
}

type Column struct {
	Title string
	Width ColumnWidth
}

// TODO: Add a "remote" data source option
// Make interface representing a data source, getRows, getColumns, sorting, update
type Options struct {
	style.Margin
}

type model[T any] struct {
	base       *app.Base[T]
	table      baseTable[T]
	opts       *Options
	style      lipgloss.Style
	focusStyle lipgloss.Style
	columns    []Column
}

func New[T any](ctx *app.Context[T], clms []Column, rows []Row, options *Options) *app.Base[T] {
	if options == nil {
		options = &Options{}
	}

	base := app.New(ctx, app.WithFocusable(true), app.WithGrow(true))

	baseStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(ctx.Styles.Colors.Ghost)
	focusStyle := baseStyle.BorderForeground(ctx.Styles.Colors.White)

	baseStyle = style.ApplyMargin(baseStyle, options.Margin)

	t := newBaseTable(ctx, base.ID,
		withColumns[T](columnMapping(20, clms)),
		withRows[T](rows),
		withFocused[T](true),
		withHeight[T](7),
	)

	s := defaultStyles(ctx)

	// TODO: Make options for table model with different variants
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(ctx.Styles.Colors.GhostLight).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(ctx.Styles.Colors.PrimaryLight).
		Background(ctx.Styles.Colors.UIPanelBackground).
		Bold(false)
	s.Hovered = s.Hovered.
		Foreground(ctx.Styles.Colors.PrimaryLight).
		Background(ctx.Styles.Colors.HighlightBackground).
		Bold(false)
	t.SetStyles(s)

	return model[T]{
		base:       base,
		table:      t,
		opts:       options,
		style:      baseStyle,
		focusStyle: focusStyle,
		columns:    clms,
	}.Base()
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

func (m model[T]) Init() tea.Cmd {
	return nil
}

func (m model[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.MouseMsg:
		if !m.base.Ctx.Zone.Get(m.base.ID).InBounds(msg) {
			m.table.clearHover()
			return m, nil
		}
	case tea.WindowSizeMsg:
		m.base.Height = msg.Height
		m.base.Width = msg.Width
	case app.FocusComponentMsg:
		if msg.TargetID == m.base.ID {
			m.table.setFocus()
		} else {
			m.table.blur()
		}
	case tea.KeyMsg:
		if m.base.Focused && m.table.getCursor() >= 0 {
			cmds = append(cmds, func() tea.Msg {
				return RowActionMsg{
					Key:     msg,
					TableID: m.base.ID,
					Row:     m.table.getCursor(),
					RowID:   m.table.rows[m.table.getCursor()][0], // Column 0 is the ID for now
				}
			})
		}
	}

	m.table.setHeight(max(0, m.base.Height-m.style.GetVerticalFrameSize()-2))
	m.table.setWidth(max(0, m.base.Width-m.style.GetHorizontalFrameSize()-2))
	m.table.setColumns(columnMapping(max(0, m.base.Width-len(m.columns)*2)-2, m.columns))
	newTable, cmd := m.table.Update(msg)
	newTableTyped := newTable.(baseTable[T])
	m.table = newTableTyped

	cmds = append(cmds, cmd)

	cmd = m.base.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model[T]) View() string {
	s := m.style
	if m.base.Focused {
		s = m.focusStyle
	}
	return m.base.Ctx.Zone.Mark(m.base.ID, s.Render(m.table.View()))
}

func (m model[T]) Base() *app.Base[T] {
	m.base.Model = m
	return m.base
}
