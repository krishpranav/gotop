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

func getCPUPercent(data *types.StatsJSON) float64 {
	cpuPercent := 0.0
	numCPUs := len(data.CPUStats.CPUUsage.PercpuUsage)

	cpuDelta := float64(data.CPUStats.CPUUsage.TotalUsage) - float64(data.PreCPUStats.CPUUsage.TotalUsage)

	systemDelta := float64(data.CPUStats.SystemUsage) - float64(data.PreCPUStats.SystemUsage)

	if cpuDelta > 0.0 && systemDelta > 0.0 {
		cpuPercent = (cpuDelta / systemDelta) * float64(numCPUs) * 100.0
	}
	return cpuPercent
}

func getPerCPUPercents(data *types.StatsJSON) []string {
	preLen := len(data.PreCPUStats.CPUUsage.PercpuUsage)
	postLen := len(data.CPUStats.CPUUsage.PercpuUsage)
	numCPUs := ui.MaxInt(preLen, postLen)

	perCpuPercents := make([]string, numCPUs)
	systemDelta := float64(data.CPUStats.SystemUsage) - float64(data.PreCPUStats.SystemUsage)

	if preLen != postLen {
		for i := range perCpuPercents {
			perCpuPercents[i] = "NA"
		}
	} else {
		for i, usage := range data.CPUStats.CPUUsage.PercpuUsage {
			perCpuPercent := 0.0

			cpuDelta := float64(usage) - float64(data.PreCPUStats.CPUUsage.PercpuUsage[i])

			if cpuDelta > 0.0 && systemDelta > 0.0 {
				perCpuPercent = (cpuDelta / systemDelta) * float64(numCPUs) * 100.0
			}
			perCpuPercents[i] = fmt.Sprintf("%.2f%%", perCpuPercent)
		}
	}
	return perCpuPercents
}

func ContainerWait(ctx context.Context, cli *client.Client, cid, state string) error {

	t := time.NewTicker(100 * time.Millisecond)
	tick := t.C

	for range tick {
		data, err := cli.ContainerInspect(ctx, cid)
		if err != nil {
			return err
		}
		if data.State.Status == state {
			return nil
		}
	}

	return nil
}

func GetContainerMetrics(ctx context.Context, cli *client.Client, cid string) (PerContainerMetrics, error) {

	metrics := PerContainerMetrics{}

	args := filters.NewArgs(
		filters.KeyValuePair{
			Key:   "id",
			Value: cid,
		},
	)

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{Filters: args})
	if err != nil {
		return metrics, err
	}

	if len(containers) > 1 {
		return metrics, fmt.Errorf("multiple containers with same ID exist")
	} else if len(containers) < 1 {
		return metrics, general.ErrInvalidContainer
	}

	c := containers[0]

	inspectData, err := cli.ContainerInspect(ctx, cid)
	if err != nil {
		return metrics, nil
	}

	stats, err := cli.ContainerStats(ctx, cid, false)
	if err != nil {
		return metrics, err
	}

	data := types.StatsJSON{}
	err = json.NewDecoder(stats.Body).Decode(&data)
	if err != nil {
		return metrics, err
	}
	stats.Body.Close()

	cpuPercent := getCPUPercent(&data)

	perCpuPercents := getPerCPUPercents(&data)

	memPercent := float64(data.MemoryStats.Usage) / float64(data.MemoryStats.Limit) * 100

	var rx, tx float64
	for _, v := range data.Networks {
		rx += float64(v.RxBytes)
		tx += float64(v.TxBytes)
	}

	var blkRead, blkWrite uint64
	for _, bioEntry := range data.BlkioStats.IoServiceBytesRecursive {
		switch strings.ToLower(bioEntry.Op) {
		case "read":
			blkRead = blkRead + bioEntry.Value
		case "write":
			blkWrite = blkWrite + bioEntry.Value
		}
	}

	// Get Network Settings
	netData := []netInfo{}
	for _, network := range c.NetworkSettings.Networks {
		id := network.NetworkID

		net, err := cli.NetworkInspect(ctx, id, types.NetworkInspectOptions{})
		if err != nil {
			continue
		}

		n := netInfo{
			Name:    net.Name,
			Driver:  net.Driver,
			Ip:      network.IPAddress,
			Ingress: net.Ingress,
		}

		netData = append(netData, n)
	}

	portData := []portMap{}
	for _, port := range c.Ports {
		p := portMap{
			IP:        port.IP,
			Host:      int(port.PublicPort),
			Container: int(port.PrivatePort),
			Protocol:  port.Type,
		}

		portData = append(portData, p)
	}

	mountData := []mountInfo{}
	for _, mount := range c.Mounts {
		m := mountInfo{
			Src:  mount.Source,
			Dst:  mount.Destination,
			Mode: mount.Mode,
		}

		mountData = append(mountData, m)
	}

	procs, err := cli.ContainerTop(ctx, cid, []string{})
	if err != nil {
		return metrics, nil
	}

	procData := []procInfo{}
	for _, proc := range procs.Processes {
		p := procInfo{
			UID: proc[0],
			PID: proc[1],
			CMD: proc[7],
		}

		procData = append(procData, p)
	}

	metrics = PerContainerMetrics{
		ID:      c.ID[:10],
		Image:   c.Image,
		Name:    strings.TrimLeft(strings.Join(c.Names, ","), "/"),
		Status:  c.Status,
		State:   c.State,
		Cpu:     cpuPercent,
		Mem:     memPercent,
		Net:     netStat{Rx: rx, Tx: tx},
		Blk:     blkStat{Read: blkRead, Write: blkWrite},
		Pid:     fmt.Sprintf("%d", inspectData.State.Pid),
		NetInfo: netData,
		PerCPU:  perCpuPercents,
		PortMap: portData,
		Mounts:  mountData,
		Procs:   procData,
	}

	return metrics, nil
}