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
