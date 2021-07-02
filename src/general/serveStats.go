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
