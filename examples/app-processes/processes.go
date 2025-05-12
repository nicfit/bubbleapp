package main

import (
	"context"
	"log"
	"sort"
	"strconv"
	"time"

	"github.com/alexanderbh/bubbleapp/component/table"
	"github.com/shirou/gopsutil/v4/process"
)

type ProcessInfo struct {
	PID    int32
	Name   string
	CPU    float64
	Memory uint64
}

func monitorProcesses(context context.Context, setProcesses func([]table.Row)) {
	for {
		select {
		case <-context.Done():
			return
		default:
			processes, err := process.Processes()
			if err != nil {
				log.Printf("Error getting processes: %v", err)
				time.Sleep(5 * time.Second) // Wait before retrying
				continue
			}

			currentProcesses := make([]ProcessInfo, 0, len(processes))

			for _, p := range processes {
				pid := p.Pid
				name, err := p.Name()
				if err != nil {
					name = "<error>"
				}

				cpuPercent, err := p.CPUPercent()
				if err != nil {
					cpuPercent = 0.0
				}

				memInfo, err := p.MemoryInfo()
				var rss uint64
				if err != nil {
					rss = 0
				} else if memInfo != nil {
					rss = memInfo.RSS
				}

				procInfo := ProcessInfo{
					PID:    pid,
					Name:   name,
					CPU:    cpuPercent,
					Memory: rss,
				}
				currentProcesses = append(currentProcesses, procInfo)
			}

			setProcesses(generateRowsOfProcesses(currentProcesses))
			time.Sleep(time.Second)
		}
	}
}

func generateRowsOfProcesses(ps []ProcessInfo) []table.Row {
	rows := make([]table.Row, len(ps))

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].CPU > ps[j].CPU
	})

	for i, proc := range ps {
		rows[i] = table.Row{
			strconv.Itoa(int(proc.PID)),
			proc.Name,
			strconv.FormatFloat(proc.CPU, 'f', 2, 64) + "%",
			strconv.FormatUint(proc.Memory/1024/1024, 10) + "MB",
		}
	}
	return rows
}
