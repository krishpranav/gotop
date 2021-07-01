package misc

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var ErrorString string

type ErrorBox struct {
	*widgets.List
}

// NewErrorBox is a constructor for the ErrorBox type
func NewErrorBox() *ErrorBox {
	return &ErrorBox{
		List: widgets.NewList(),
	}
}

// Resize resizes the widget based on specified width
// and height
func (errBox *ErrorBox) Resize(termWidth, termHeight int) {
	textWidth := 50
	for _, line := range errorKeybindings {
		if textWidth < len(line) {
			textWidth = len(line) + 2
		}
	}
	textHeight := len(errorKeybindings) + 5
	x := (termWidth - textWidth) / 2
	y := (termHeight - textHeight) / 2
	if x < 0 {
		x = 0
		textWidth = termWidth
	}
	if y < 0 {
		y = 0
		textHeight = termHeight
	}

	errBox.List.SetRect(x, y, textWidth+x, textHeight+y)
}

// Draw puts the required text into the widget
func (errBox *ErrorBox) Draw(buf *ui.Buffer) {
	errBox.List.Title = " Error "

	errBox.List.Rows = []string{ErrorString, ""}
	errBox.List.Rows = append(errBox.List.Rows, errorKeybindings...)
	errBox.List.TextStyle = ui.NewStyle(ui.ColorYellow)
	errBox.List.WrapText = false
	errBox.List.Draw(buf)
}

// SetErrorString sets the error string to be displayed
func SetErrorString(errStr string) {
	ErrorString = errStr
}
