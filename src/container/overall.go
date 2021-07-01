package container

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type ContainerMetrics struct {
	TotalCPU     float64
	TotalMem     float64
	TotalNet     netStat
	TotalBlk     blkStat
	PerContainer []PerContainerMetrics
}

func GetOverallMetrics(ctx context.Context, cli *client.Client, all bool) (ContainerMetrics, error) {
	metrics := ContainerMetrics{}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: all})
	if err != nil {
		return metrics, err
	}

	metrcisChan := make(chan PerContainerMetrics, len(containers))

	for _, container := range containers {
		go getMetrics(ctx, cli, container, metrcisChan)
	}

	var totalCPU, totalMem float64
	totalNet := netStat{}
	totalBlk := blkStat{}

	for range containers {
		metric := <-metrcisChan

		totalCPU += metric.Cpu

		totalMem += metric.Mem

		totalNet.Rx += metric.Net.Rx
		totalNet.Tx += metric.Net.Tx

		totalBlk.Read += metric.Blk.Read
		totalBlk.Write += metric.Blk.Write

		metrics.PerContainer = append(metrics.PerContainer, metric)
	}

	metrics.TotalCPU = totalCPU
	metrics.TotalMem = totalMem
	metrics.TotalNet = totalNet
	metrics.TotalBlk = totalBlk

	return metrics, nil
}

func getMetrics(ctx context.Context, cli *client.Client, c types.Container, ch chan PerContainerMetrics) {

	metrics := PerContainerMetrics{}
	defer func() {
		ch <- metrics
	}()

	stats, err := cli.ContainerStats(ctx, c.ID, false)
	if err != nil {
		return
	}

	data := types.StatsJSON{}
	err = json.NewDecoder(stats.Body).Decode(&data)
	if err != nil {
		return
	}
	stats.Body.Close()

	cpuPercent := getCPUPercent(&data)

	memPercent := 0.0
	if data.MemoryStats.Limit > 0 {
		memPercent = float64(data.MemoryStats.Usage) / float64(data.MemoryStats.Limit) * 100
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

	var rx, tx float64

	for _, v := range data.Networks {
		rx += float64(v.RxBytes)
		tx += float64(v.TxBytes)
	}

	metrics = PerContainerMetrics{
		ID:     c.ID[:10],
		Image:  c.Image,
		Name:   strings.TrimLeft(strings.Join(c.Names, ","), "/"),
		Status: c.Status,
		State:  c.State,
		Cpu:    cpuPercent,
		Mem:    memPercent,
		Net:    netStat{Rx: rx, Tx: tx},
		Blk:    blkStat{Read: blkRead, Write: blkWrite},
	}
}