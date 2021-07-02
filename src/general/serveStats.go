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
