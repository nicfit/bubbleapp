package main

type ProcessInfo struct {
	PID    int32
	Name   string
	CPU    float64
	Memory uint64
}

type AppData struct {
	Processes []ProcessInfo
}
