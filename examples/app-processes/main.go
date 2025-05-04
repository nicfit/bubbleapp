package main

import (
	"os"
	"sort"
	"strconv"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/table"
	"github.com/alexanderbh/bubbleapp/component/text"

	tea "github.com/charmbracelet/bubbletea/v2"
)

type ProcessInfo struct {
	PID    int32
	Name   string
	CPU    float64
	Memory uint64
}

type AppState struct {
	Processes []ProcessInfo
}

var clms = []table.Column{
	{Title: "PID", Width: table.WidthInt(10)},
	{Title: "Name", Width: table.WidthGrow()},
	{Title: "CPU", Width: table.WidthInt(10)},
	{Title: "Memory", Width: table.WidthInt(10)},
}

func NewRoot(ctx *app.Context[AppState]) app.Fc[AppState] {

	stackChildren := []app.Fc[AppState]{
		text.NewDynamic(ctx, func(ctx *app.Context[AppState]) string {
			return "# Processes: " + strconv.Itoa(len(ctx.Data.Processes))
		}, nil),
		table.New(ctx, func(ctx *app.Context[AppState]) ([]table.Column, []table.Row) {
			rows := make([]table.Row, len(ctx.Data.Processes))

			sort.Slice(ctx.Data.Processes, func(i, j int) bool {
				return ctx.Data.Processes[i].CPU > ctx.Data.Processes[j].CPU
			})

			for i, proc := range ctx.Data.Processes {
				rows[i] = table.Row{
					strconv.Itoa(int(proc.PID)),
					proc.Name,
					strconv.FormatFloat(proc.CPU, 'f', 2, 64) + "%",
					strconv.FormatUint(proc.Memory/1024/1024, 10) + "MB",
				}
			}
			return clms, rows
		}, nil),
		button.New(ctx, "Quit",
			func(ctx *app.Context[AppState]) {
				ctx.Quit()
			}, nil,
		),
	}

	return stack.New(ctx, stackChildren, nil)
}

func main() {
	ctx := app.NewContext(&AppState{})

	go monitorProcesses(ctx.Data)

	p := tea.NewProgram(app.NewApp(ctx, NewRoot(ctx)), tea.WithAltScreen(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
