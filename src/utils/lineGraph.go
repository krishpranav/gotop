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

// NewLineGraph creates and returns a lineGraph instance
func NewLineGraph() *LineGraph {
	return &LineGraph{
		Block: ui.NewBlock(),

		Data:   make(map[string][]float64),
		Labels: make(map[string]string),

		HorizontalScale: 5,

		LineColors: make(map[string]ui.Color),
	}
}

func (l *LineGraph) Draw(buf *ui.Buffer) {
	l.Block.Draw(buf)
	c := drawille.NewCanvas()
	colors := make([][]ui.Color, l.Inner.Dx()+2)
	for i := range colors {
		colors[i] = make([]ui.Color, l.Inner.Dy()+2)
	}

	seriesList := make([]string, len(l.Data))
	i := 0
	l.MaxVal = 1
	for seriesName := range l.Data {
		for _, val := range l.Data[seriesName] {
			if val > l.MaxVal {
				l.MaxVal = val
			}
		}
		seriesList[i] = seriesName
		i++
	}
	sort.Strings(seriesList)

	for i := len(seriesList) - 1; i >= 0; i-- {
		seriesName := seriesList[i]
		seriesData := l.Data[seriesName]
		seriesLineColor, ok := l.LineColors[seriesName]
		if !ok {
			seriesLineColor = l.DefaultLineColor
		}

		lastY, lastX := -1, -1
		for i := len(seriesData) - 1; i >= 0; i-- {
			x := ((l.Inner.Dx() + 1) * 2) - 1 - (((len(seriesData) - 1) - i) * l.HorizontalScale)
			y := ((l.Inner.Dy() + 1) * 4) - 1 - int((float64((l.Inner.Dy())*4)-1)*(seriesData[i]/float64(l.MaxVal)))
			if x < 0 {

				if x > 0-l.HorizontalScale {
					for _, p := range drawille.Line(lastX, lastY, x, y) {
						if p.X > 0 {
							c.Set(p.X, p.Y)
							colors[p.X/2][p.Y/4] = seriesLineColor
						}
					}
				}
				break
			}
			if lastY == -1 {
				c.Set(x, y)
				colors[x/2][y/4] = seriesLineColor
			} else {
				c.DrawLine(lastX, lastY, x, y)
				for _, p := range drawille.Line(lastX, lastY, x, y) {
					colors[p.X/2][p.Y/4] = seriesLineColor
				}
			}
			lastX, lastY = x, y
		}

		for y, line := range c.Rows(c.MinX(), c.MinY(), c.MaxX(), c.MaxY()) {
			for x, char := range line {
				x /= 3
				if x == 0 {
					continue
				}
				if char != 10240 {
					buf.SetCell(
						ui.NewCell(char, ui.NewStyle(colors[x][y])),
						image.Pt(l.Inner.Min.X+x-1, l.Inner.Min.Y+y-1),
					)
				}
			}
		}
	}

	for i, seriesName := range seriesList {
		if i+2 > l.Inner.Dy() {
			continue
		}
		seriesLineColor, ok := l.LineColors[seriesName]
		if !ok {
			seriesLineColor = l.DefaultLineColor
		}

		str := seriesName + " " + l.Labels[seriesName]
		for k, char := range str {
			if char != ' ' {
				buf.SetCell(
					ui.NewCell(char, ui.NewStyle(seriesLineColor)),
					image.Pt(l.Inner.Min.X+2+k, l.Inner.Min.Y+i+1),
				)
			}
		}

	}
}
