package main

import (
	"context"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"time"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/table"
	"github.com/alexanderbh/bubbleapp/component/text"
	"github.com/alexanderbh/bubbleapp/component/tickfps"

	tea "github.com/charmbracelet/bubbletea/v2"
)

func NewRoot(ctx *app.Ctx, _ app.Props) string {
	processes, setProcesses := app.UseState(ctx, []table.Row{})

	app.UseEffectWithCleanup(ctx, func() func() {
		processCtx, cancel := context.WithCancel(context.Background())
		go monitorProcesses(processCtx, setProcesses)
		return cancel
	}, app.RunOnceDeps)

	return stack.New(ctx, func(ctx *app.Ctx) {
		text.New(ctx, "# Processes: "+strconv.Itoa(len(processes)))
		table.New(ctx, table.WithDataFunc(func(ctx *app.Ctx) ([]table.Column, []table.Row) {
			return clms, processes
		}))
		tickfps.NewAtInterval(ctx, 1*time.Second)
		button.New(ctx, "Quit", ctx.Quit, button.WithType(button.Compact), button.WithVariant(button.Danger))

	})
}

func main() {
	// pprof - used for debugging performance - just ignore
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	ctx := app.NewCtx()

	app := app.New(ctx, NewRoot)
	p := tea.NewProgram(app, tea.WithAltScreen(), tea.WithMouseAllMotion())
	app.SetTeaProgram(p)
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
