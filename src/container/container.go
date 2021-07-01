package container

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	ui "github.com/gizak/termui/v3"
	"github.com/pesos/grofer/src/general"
)

type PerContainerMetrics struct {
	ID     string
	Image  string
	Name   string
	Status string
	State  string
	Cpu    float64
	Mem    float64
	Net    netStat
	Blk    blkStat
	Pid     string
	NetInfo []netInfo
	PerCPU  []string
	PortMap []portMap
	Mounts  []mountInfo
	Procs   []procInfo
}

type netStat struct {
	Rx float64
	Tx float64
}

type blkStat struct {
	Read  uint64
	Write uint64
}

type netInfo struct {
	Name    string
	Driver  string
	Ip      string
	Ingress bool
}

type mountInfo struct {
	Src  string
	Dst  string
	Mode string
}

type portMap struct {
	IP        string
	Host      int
	Container int
	Protocol  string
}

type procInfo struct {
	UID string
	PID string
	CMD string
}