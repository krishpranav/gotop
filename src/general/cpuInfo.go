package general

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pesos/grofer/src/utils"
)

// CPULoad type contains info about load on CPU from various sources
// as well as general stats about the CPU.
type CPULoad struct {
	CPURates [][]string `json:"-"`
	Usr      int        `json:"usr"`
	Nice     int        `json:"nice"`
	Sys      int        `json:"sys"`
	Iowait   int        `json:"iowait"`
	Soft     int        `json:"soft"`
	Steal    int        `json:"steal"`
	Guest    int        `json:"guest"`
	Gnice    int        `json:"gnice"`
	Idle     int        `json:"idle"`
	Irq      int        `json:"irq"`
}

func NewCPULoad() *CPULoad {
	return &CPULoad{}
}

func (c *CPULoad) readCPULoad() error {
	file, err := os.Open("/proc/stat")
	if err != nil {
		return err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	data, err := reader.ReadBytes(byte('\n'))
	if err != nil {
		return err
	}
	vals := strings.Field(string(data))[1:]
	var avg [10]float64
	sum := 0
	for i, x := range vals {
		curr, err := strconv.Atoi(x)
		if err != nil {
			return err
		} else {
			avg[i] = float64(curr)
			sum += curr
		}
	}

	for i, x := range avg {
		avg[i] = 100 * x / float64(sum)
	}

}
