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

func NewRoot(c *app.Ctx) app.C {
	processes, setProcesses := app.UseState(c, []table.Row{})

	app.UseEffectWithCleanup(c, func() func() {
		processCtx, cancel := context.WithCancel(context.Background())
		go monitorProcesses(processCtx, func(r []table.Row) { setProcesses(r) })
		return cancel
	}, app.RunOnceDeps)

	return stack.New(c, func(c *app.Ctx) []app.C {
		return []app.C{
			text.New(c, "# Processes: "+strconv.Itoa(len(processes))),
			table.New(c, table.WithDataFunc(func(c *app.Ctx) ([]table.Column, []table.Row) {
				return clms, processes
			})),
			tickfps.NewAtInterval(c, 1*time.Second),
			button.New(c, "Quit", c.Quit, button.WithVariant(button.Danger)),
		}
	})
}

func main() {
	// pprof - used for debugging performance - just ignore
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	c := app.NewCtx()

	bubbleApp := app.New(c, NewRoot)
	p := tea.NewProgram(bubbleApp, tea.WithAltScreen(), tea.WithMouseAllMotion())
	bubbleApp.SetTeaProgram(p)
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
