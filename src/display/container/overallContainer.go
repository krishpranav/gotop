package container

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	ui "github.com/gizak/termui/v3"
	"github.com/pesos/grofer/src/display/misc"
	info "github.com/pesos/grofer/src/general"

	"github.com/pesos/grofer/src/container"
	"github.com/pesos/grofer/src/utils"
)

var runProc = true
var helpVisible = false
var errorVisible = false

var sortIdx = -1
var sortAsc = false

const (
	UP_ARROW   = "▲"
	DOWN_ARROW = "▼"
)

var header = []string{
	"ID",
	"Image",
	"Name",
	"Status",
	"State",
	"CPU",
	"Memory",
	"Net I/O",
	"Block I/O",
}
