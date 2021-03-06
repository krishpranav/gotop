package general

import (
	"context"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/krishpranav/gotop/src/utils"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

func roundOff(num uint64) float64 {
	x := float64(num) / (1024 * 1024 * 1024)
	return math.Round(x*10) / 10
}

func GetCPURates() ([]float64, error) {
	cpuRates, err := cpu.Percent(time.Second, true)
	if err != nil {
		return nil, err
	}
	return cpuRates, nil
}

func ServeCPURates(ctx context.Context, cpuChannel chan utils.DataStats) error {
	cpuRates, err := cpu.Percent(time.Second, true)
	if err != nil {
		return err
	}
	data := utils.DataStats{
		CpuStats: cpuRates,
		FieldSet: "CPU",
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case cpuChannel <- data:
		return nil
	}
}

func ServeMemRates(ctx context.Context, dataChannel chan utils.DataStats) error {
	memory, err := mem.VirtualMemory()
	if err != nil {
		return err
	}

	memRates := []float64{roundOff(memory.Total), roundOff(memory.Available), roundOff(memory.Used), roundOff(memory.Free)}

	data := utils.DataStats{
		MemStats: memRates,
		FieldSet: "MEM",
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case dataChannel <- data:
		return nil
	}
}

func ServeDiskRates(ctx context.Context, dataChannel chan utils.DataStats) error {
	var partitions []disk.PartitionStat
	var err error
	partitions, err = disk.Partitions(false)
	if err != nil {
		return err
	}

	rows := [][]string{{"Mount", "Total", "Used %", "Used", "Free", "FS Type"}}
	for _, value := range partitions {
		usageVals, _ := disk.Usage(value.Mountpoint)

		if strings.HasPrefix(value.Device, "/dev/loop") {
			continue
		} else if strings.HasPrefix(value.Mountpoint, "/var/lib/docker") {
			continue
		} else {

			path := usageVals.Path
			total := fmt.Sprintf("%.2f G", float64(usageVals.Total)/(1024*1024*1024))
			used := fmt.Sprintf("%.2f G", float64(usageVals.Used)/(1024*1024*1024))
			usedPercent := fmt.Sprintf("%.2f %s", usageVals.UsedPercent, "%")
			free := fmt.Sprintf("%.2f G", float64(usageVals.Free)/(1024*1024*1024))
			fs := usageVals.Fstype
			row := []string{path, total, usedPercent, used, free, fs}
			rows = append(rows, row)

		}
	}

	data := utils.DataStats{
		DiskStats: rows,
		FieldSet:  "DISK",
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case dataChannel <- data:
		return nil
	}
}

func ServeNetRates(ctx context.Context, dataChannel chan utils.DataStats) error {
	netStats, err := net.IOCounters(false)
	if err != nil {
		return err
	}
	IO := make(map[string][]float64)
	for _, IOStat := range netStats {
		nic := []float64{float64(IOStat.BytesSent) / (1024), float64(IOStat.BytesRecv) / (1024)}
		IO[IOStat.Name] = nic
	}

	data := utils.DataStats{
		NetStats: IO,
		FieldSet: "NET",
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case dataChannel <- data:
		return nil
	}
}
