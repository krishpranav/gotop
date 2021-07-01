package misc

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var keybindings []string

type HelpMenu struct {
	*widgets.List
}

// NewHelpMenu is a constructor for the HelpMenu type
func NewHelpMenu() *HelpMenu {
	return &HelpMenu{
		List: widgets.NewList(),
	}
}

func (help *HelpMenu) Resize(termWidth, termHeight int) {
	textWidth := 50
	for _, line := range keybindings {
		if textWidth < len(line) {
			textWidth = len(line) + 2
		}
	}
	textHeight := len(keybindings) + 3
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

	help.List.SetRect(x, y, textWidth+x, textHeight+y)
}

// Draw puts the required text into the widget
func (help *HelpMenu) Draw(buf *ui.Buffer) {
	help.List.Title = " Keybindings "

	help.List.Rows = keybindings
	help.List.TextStyle = ui.NewStyle(ui.ColorYellow)
	help.List.WrapText = false
	help.List.Draw(buf)
}

func SelectHelpMenu(page string) {
	switch page {
	case "proc":
		keybindings = procKeybindings
	case "proc_pid":
		keybindings = perProcKeyBindings
	case "main":
		keybindings = mainKeybindings
	case "cont":
		keybindings = containerKeybindings
	case "cont_cid":
		keybindings = perContainerKeyBindings
	}
}
