package process

import (
	proc "github.com/shirou/gopsutil/process"
)

type Process struct {
	Proc           *proc.Process
	MemoryInfo     *proc.MemoryInfoStat
	PageFault      *proc.PageFaultsStat
	NumCtxSwitches *proc.NumCtxSwitchesStat
	Exe            string
	Name           string
	Status         string
	Children       []*proc.Process
	Gids           []int32
	CPUAffinity    []int32
	CreateTime     int64
	CPUPercent     float64
	Nice           int32
	NumThreads     int32
	MemoryPercent  float32
	IsRunning      bool
	Foreground     bool
	Background     bool
}
