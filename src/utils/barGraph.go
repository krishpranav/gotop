package utils

import (
	"fmt"
	"image"

	rw "github.com/mattn/go-runewidth"

	ui "github.com/gizak/termui/v3"
)

type BarChart struct {
	NumFormatter func(float64) string
	Labels       []string
	BarColors    []ui.Color
	LabelStyles  []ui.Style
	NumStyles    []ui.Style
	Data         []float64
	ui.Block
	BarWidth int
	BarGap   int
	MaxVal   float64
}

func NewBarChart() *BarChart {
	return &BarChart{
		Block:        *ui.NewBlock(),
		BarColors:    ui.Theme.BarChart.Bars,
		NumStyles:    ui.Theme.BarChart.Nums,
		LabelStyles:  ui.Theme.BarChart.Labels,
		NumFormatter: func(n float64) string { return fmt.Sprint(n) },
		BarGap:       1,
		BarWidth:     3,
	}
}