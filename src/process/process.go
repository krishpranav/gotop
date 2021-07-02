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

func InitAllProcs() (map[int32]*Process, error) {
	var processes map[int32]*Process = make(map[int32]*Process)
	pids, err := proc.Processes()

	if err != nil {
		return processes, err
	}

	for _, proc := range pids {
		tempProc := &Process{Proc: proc}
		processes[proc.Pid] = tempProc
	}
	return processes, nil
}

func NewProcess(pid int32) (*Process, error) {
	process, err := proc.NewProcess(pid)
	if err != nil {
		return nil, err
	}
	newProcess := &Process{Proc: process}
	return newProcess, nil
}
