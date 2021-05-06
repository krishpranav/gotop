package main

// imports
import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/gizak/termui"
)

const statFilePath = "/proc/stat"
const meminfoFilePath = "/proc/meminfo"
const netinfoFilePath = "/proc/net/dev"

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type ProcessList struct {

}

type CpuStat struct {
	user 	float32
	nice 	float32
	system	float32
	idle 	float32
}

type CpusStats struct {
	stat map[string]CpuStat
	proc map[string]CpuStat
}

