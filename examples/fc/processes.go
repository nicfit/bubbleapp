package main

import (
	"log"
	"time"

	"github.com/shirou/gopsutil/v4/process"
)

func monitorProcesses(state *AppState) {
	for {
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

		state.Processes = currentProcesses

		time.Sleep(time.Second / 1)
	}
}
