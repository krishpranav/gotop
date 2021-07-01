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