package main

import (
	"net/http"
	_ "net/http/pprof"
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

func NewRoot(ctx *app.Context[AppState]) app.Fc[AppState] {
	return stack.New(ctx, []app.Fc[AppState]{
		text.NewDynamic(ctx, func(ctx *app.Context[AppState]) string {
			return "# Processes: " + strconv.Itoa(len(ctx.Data.Processes))
		}, nil),
		table.NewDynamic(ctx, func(ctx *app.Context[AppState]) ([]table.Column, []table.Row) {
			rows := generateRowsOfProcesses(ctx.Data)
			return clms, rows
		}, nil),
		button.New(ctx, "Quit", app.Quit, &button.Options{Variant: button.Danger, Type: button.Compact}),
	}, nil)
}

func main() {
	// pprof - used for debugging performance - just ignore
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	ctx := app.NewContext(&AppState{})

	go monitorProcesses(ctx.Data)

	p := tea.NewProgram(app.NewApp(ctx, NewRoot), tea.WithAltScreen(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}

var clms = []table.Column{
	{Title: "PID", Width: table.WidthInt(10)},
	{Title: "Name", Width: table.WidthGrow()},
	{Title: "CPU", Width: table.WidthInt(10)},
	{Title: "Memory", Width: table.WidthInt(10)},
}

func generateRowsOfProcesses(state *AppState) []table.Row {
	rows := make([]table.Row, len(state.Processes))

	sort.Slice(state.Processes, func(i, j int) bool {
		return state.Processes[i].CPU > state.Processes[j].CPU
	})

	for i, proc := range state.Processes {
		rows[i] = table.Row{
			strconv.Itoa(int(proc.PID)),
			proc.Name,
			strconv.FormatFloat(proc.CPU, 'f', 2, 64) + "%",
			strconv.FormatUint(proc.Memory/1024/1024, 10) + "MB",
		}
	}
	return rows
}
