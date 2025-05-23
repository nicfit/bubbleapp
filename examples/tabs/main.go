package main

import (
	"fmt"
	"os"

	"net/http"
	_ "net/http/pprof"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/tabs"

	tea "github.com/charmbracelet/bubbletea/v2"
)

var tabsData = []tabs.Tab{
	{Title: "Overview", Content: overview},
	{Title: "Loaders", Content: loaders},
	{Title: "Yet another", Content: tabtab},
}

func NewRoot(c *app.Ctx) *app.C {
	return tabs.New(c, tabsData)
}

func main() {
	// pprof - used for debugging performance - just ignore
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	c := app.NewCtx()

	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()

	bubbleApp := app.New(c, NewRoot)
	p := tea.NewProgram(bubbleApp, tea.WithAltScreen(), tea.WithMouseAllMotion())
	bubbleApp.SetTeaProgram(p)
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
