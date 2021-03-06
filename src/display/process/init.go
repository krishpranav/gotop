package process

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/krishpranav/gotop/src/utils"
)

// PerProcPage holds the ui elements rendered by the command gotop proc -p PID
type PerProcPage struct {
	Grid             *ui.Grid
	CPUChart         *widgets.Gauge
	MemChart         *widgets.Gauge
	PIDTable         *widgets.Table
	ChildProcsTable  *utils.Table
	CTXSwitchesChart *utils.BarChart
	PageFaultsChart  *utils.BarChart
	MemStatsChart    *utils.BarChart
}

// NewProcPage initializes a new page from the PerProcPage struct and returns it
func NewPerProcPage() *PerProcPage {
	page := &PerProcPage{
		Grid:             ui.NewGrid(),
		CPUChart:         widgets.NewGauge(),
		MemChart:         widgets.NewGauge(),
		PIDTable:         widgets.NewTable(),
		ChildProcsTable:  utils.NewTable(),
		CTXSwitchesChart: utils.NewBarChart(),
		PageFaultsChart:  utils.NewBarChart(),
		MemStatsChart:    utils.NewBarChart(),
	}
	page.InitPerProc()
	return page
}

// InitPerProc initializes and sets the ui and grid for gotop proc -p PID
func (page *PerProcPage) InitPerProc() {
	// Initialize Gauge for CPU Chart
	page.CPUChart.Title = " CPU % "
	page.CPUChart.LabelStyle.Fg = ui.ColorClear
	page.CPUChart.BarColor = ui.ColorGreen
	page.CPUChart.BorderStyle.Fg = ui.ColorCyan
	page.CPUChart.TitleStyle.Fg = ui.ColorClear

	// Initialize Gauge for Memory Chart
	page.MemChart.Title = " Mem % "
	page.MemChart.LabelStyle.Fg = ui.ColorClear
	page.MemChart.BarColor = ui.ColorGreen
	page.MemChart.BorderStyle.Fg = ui.ColorCyan
	page.MemChart.TitleStyle.Fg = ui.ColorClear

	// Initialize Table for PID Details Table
	page.PIDTable.TextStyle = ui.NewStyle(ui.ColorClear)
	page.PIDTable.TextAlignment = ui.AlignCenter
	page.PIDTable.RowSeparator = false
	page.PIDTable.Title = " PID "
	page.PIDTable.BorderStyle.Fg = ui.ColorCyan
	page.PIDTable.TitleStyle.Fg = ui.ColorClear

	// Initialize List for Child Processes list
	page.ChildProcsTable.Title = " Child Processes "
	page.ChildProcsTable.BorderStyle.Fg = ui.ColorCyan
	page.ChildProcsTable.TitleStyle.Fg = ui.ColorClear
	page.ChildProcsTable.ColWidths = []int{10, 10}
	page.ChildProcsTable.Header = []string{"PID", "Command"}
	page.ChildProcsTable.ShowCursor = true
	page.ChildProcsTable.CursorColor = ui.ColorCyan
	page.ChildProcsTable.ColResizer = func() {
		x := page.ChildProcsTable.Inner.Dx() - 10
		page.ChildProcsTable.ColWidths = []int{
			10,
			ui.MaxInt(10, x),
		}
	}

	// Initialize Bar Chart for CTX Swicthes Chart
	page.CTXSwitchesChart.Data = []float64{0, 0}
	page.CTXSwitchesChart.Labels = []string{"Volun", "Involun"}
	page.CTXSwitchesChart.Title = " Ctx switches "
	page.CTXSwitchesChart.BorderStyle.Fg = ui.ColorCyan
	page.CTXSwitchesChart.TitleStyle.Fg = ui.ColorClear
	page.CTXSwitchesChart.BarWidth = 9
	page.CTXSwitchesChart.BarColors = []ui.Color{ui.ColorGreen, ui.ColorCyan}
	page.CTXSwitchesChart.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorClear)}
	page.CTXSwitchesChart.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}

	// Initialize Bar Chart for Page Faults Chart
	page.PageFaultsChart.Data = []float64{0, 0}
	page.PageFaultsChart.Labels = []string{"minr", "mjr"}
	page.PageFaultsChart.Title = " Page Faults "
	page.PageFaultsChart.BorderStyle.Fg = ui.ColorCyan
	page.PageFaultsChart.TitleStyle.Fg = ui.ColorClear
	page.PageFaultsChart.BarWidth = 9
	page.PageFaultsChart.BarColors = []ui.Color{ui.ColorGreen, ui.ColorCyan}
	page.PageFaultsChart.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorClear)}
	page.PageFaultsChart.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}

	// Initialize Bar Chart for Memory Stats Chart
	page.MemStatsChart.Data = []float64{0, 0, 0, 0}
	page.MemStatsChart.Labels = []string{"RSS", "Data", "Stack", "Swap"}
	page.MemStatsChart.Title = " Mem Stats (mb) "
	page.MemStatsChart.BorderStyle.Fg = ui.ColorCyan
	page.MemStatsChart.TitleStyle.Fg = ui.ColorClear
	page.MemStatsChart.BarWidth = 9
	page.MemStatsChart.BarColors = []ui.Color{ui.ColorGreen, ui.ColorMagenta, ui.ColorYellow, ui.ColorCyan}
	page.MemStatsChart.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorClear)}
	page.MemStatsChart.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}

	// Initialize Grid layout
	page.Grid.Set(
		ui.NewCol(0.5,
			ui.NewRow(0.125, page.CPUChart),
			ui.NewRow(0.125, page.MemChart),
			ui.NewRow(0.35, page.PIDTable),
			ui.NewRow(0.4, page.ChildProcsTable),
		),
		ui.NewCol(0.5,
			ui.NewRow(0.6,
				ui.NewCol(0.5, page.CTXSwitchesChart),
				ui.NewCol(0.5, page.PageFaultsChart),
			),
			ui.NewRow(0.4, page.MemStatsChart),
		),
	)

	w, h := ui.TerminalDimensions()
	page.Grid.SetRect(0, 0, w, h)
}

// AllProcPage struct holds the ui elements rendered by the gotop proc command
type AllProcPage struct {
	Grid      *ui.Grid
	ProcTable *utils.Table
}

// NewAllProcsPage initializes a new page from the AllProcPage struct and returns it
func NewAllProcsPage() *AllProcPage {
	page := &AllProcPage{
		Grid:      ui.NewGrid(),
		ProcTable: utils.NewTable(),
	}
	page.InitAllProc()
	return page
}

// InitAllProc initializes and sets the ui and grid for gotop proc
func (page *AllProcPage) InitAllProc() {
	page.ProcTable.Header = []string{
		"PID",
		"Command",
		"CPU",
		"Memory",
		"Status",
		"Foreground",
		"Creation Time",
		"Thread Count",
	}
	page.ProcTable.ColWidths = []int{10, 40, 10, 10, 8, 12, 25, 15}
	page.ProcTable.ColResizer = func() {
		x := page.ProcTable.Inner.Dx() - (10 + 10 + 10 + 8 + 12 + 25 + 15)
		page.ProcTable.ColWidths = []int{
			10,
			ui.MaxInt(40, x),
			10,
			10,
			8,
			12,
			25,
			15,
		}
	}
	page.ProcTable.ShowCursor = true
	page.ProcTable.RowStyle = ui.NewStyle(ui.ColorClear)
	page.ProcTable.ColColor[1] = ui.ColorGreen
	page.ProcTable.BorderStyle.Fg = ui.ColorCyan
	page.Grid.Set(
		ui.NewRow(1.0, page.ProcTable),
	)

	w, h := ui.TerminalDimensions()
	page.Grid.SetRect(0, 0, w, h)
}
