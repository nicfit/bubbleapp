package main

import (
	"os"
	"sort"
	"strconv"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/divider"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/table"
	"github.com/alexanderbh/bubbleapp/component/text"
	"github.com/alexanderbh/bubbleapp/style"

	zone "github.com/alexanderbh/bubblezone/v2"
	tea "github.com/charmbracelet/bubbletea/v2"
)

var clms = []table.Column{
	{Title: "PID", Width: table.WidthInt(10)},
	{Title: "Name", Width: table.WidthGrow()},
	{Title: "CPU", Width: table.WidthInt(10)},
	{Title: "Memory", Width: table.WidthInt(10)},
}

func NewRoot(appData *AppData) model {
	ctx := &app.Context[AppData]{
		Styles: style.DefaultStyles(),
		Zone:   zone.New(),
		Data:   appData,
	}

	processNumberText := text.New(ctx, "Processes:", nil)
	pTable := table.New(ctx, clms, nil, nil)
	stack := stack.New(ctx, &stack.Options[AppData]{
		Children: []*app.Base[AppData]{
			processNumberText,
			divider.New(ctx),
			pTable,
			text.New(ctx, "Press [q] to quit.", nil),
		}},
		app.AsRoot(),
	)

	return model{
		base:                stack,
		frame:               0,
		processNumberTextID: processNumberText.ID,
		pTableID:            pTable.ID,
	}
}

type model struct {
	base                *app.Base[AppData]
	frame               int
	processNumberTextID string
	pTableID            string
}

func (m model) Init() tea.Cmd {
	return m.base.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case app.TickMsg:
		if m.frame%12 == 0 {
			m.frame = 0
			// This is not pretty. Find a way to make thie ergonomic
			newText := text.New(m.base.Ctx, "Processes: "+strconv.Itoa(len(m.base.Ctx.Data.Processes)), nil)
			m.base.ReplaceChild(m.processNumberTextID, newText)
			m.processNumberTextID = newText.ID

			rows := make([]table.Row, len(m.base.Ctx.Data.Processes))

			sort.Slice(m.base.Ctx.Data.Processes, func(i, j int) bool {
				return m.base.Ctx.Data.Processes[i].CPU > m.base.Ctx.Data.Processes[j].CPU
			})

			// Probably better to do this in the goroutine
			for i, proc := range m.base.Ctx.Data.Processes {
				rows[i] = table.Row{
					strconv.Itoa(int(proc.PID)),
					proc.Name,
					strconv.FormatFloat(proc.CPU, 'f', 2, 64) + "%",
					strconv.FormatUint(proc.Memory/1024/1024, 10) + "MB",
				}
			}
			// This is not pretty. Find a way to make thie ergonomic
			newPTable := m.base.GetChild(m.pTableID).Model.(table.Model[AppData])
			newPTable.SetRows(rows)
			m.base.ReplaceChild(m.pTableID, newPTable.Base())
		}
		m.frame++
	}
	nM, cmd := m.base.Model.Update(msg)
	typedM := nM.(app.UIModel[AppData])
	m.base.Model = typedM
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)

}

func (m model) View() string {
	return m.base.Render()
}

func main() {
	appData := &AppData{
		Processes: []ProcessInfo{},
	}
	go monitorProcesses(appData)
	p := tea.NewProgram(NewRoot(appData), tea.WithAltScreen(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
