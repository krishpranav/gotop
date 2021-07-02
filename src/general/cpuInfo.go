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
