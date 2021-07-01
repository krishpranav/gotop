package utils

import (
	"fmt"
	"image"
	"log"
	"strings"

	ui "github.com/gizak/termui/v3"
)

type Table struct {
	*ui.Block

	Header []string
	Rows   [][]string

	HeaderStyle ui.Style
	RowStyle    ui.Style

	ColWidths []int
	ColGap    int
	PadLeft   int

	ShowCursor  bool
	CursorColor ui.Color

	ShowLocation bool

	UniqueCol    int
	SelectedItem string
	SelectedRow  int
	TopRow       int
	ColColor     map[int]ui.Color
	ColResizer   func()
}

func NewTable() *Table {
	return &Table{
		Block:       ui.NewBlock(),
		HeaderStyle: ui.NewStyle(ui.ColorClear, ui.ColorClear, ui.ModifierBold),
		RowStyle:    ui.NewStyle(ui.Theme.Default.Fg),
		SelectedRow: 0,
		TopRow:      0,
		UniqueCol:   0,
		ColResizer:  func() {},
		ColColor:    make(map[int]ui.Color),
		CursorColor: ui.ColorCyan,
	}
}
