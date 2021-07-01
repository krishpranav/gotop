package utils

import (
	"image"
	"sort"

	drawille "github.com/cjbassi/gotop/src/termui/drawille-go"
	ui "github.com/gizak/termui/v3"
)

type LineGraph struct {
	*ui.Block

	Data   map[string][]float64
	Labels map[string]string

	HorizontalScale int
	MaxVal          float64

	LineColors       map[string]ui.Color
	DefaultLineColor ui.Color
}

func NewLineGraph() *LineGraph {
	return &LineGraph{
		Block: ui.NewBlock(),

		Data:   make(map[string][]float64),
		Labels: make(map[string]string),

		HorizontalScale: 5,
		LineColors:      make(map[string]ui.Color),
	}
}
