package container

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/pesos/grofer/src/utils"
)


type OverallContainerPage struct {
	Grid         *ui.Grid
	CPUChart     *widgets.Gauge
	MemChart     *widgets.Gauge
	NetChart     *utils.BarChart
	BlkChart     *utils.BarChart
	DetailsTable *utils.Table
}