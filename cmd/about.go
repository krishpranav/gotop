package cmd

import (
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/spf13/cobra"
)

var gotopVersion string = "1.0.0"

var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "about is a command that gives information about the project.",
	Run: func(cmd *cobra.Command, args []string) {

		if err := ui.Init(); err != nil {
			log.Fatalf("failed to initialize termui: %v", err)
		}
		defer ui.Close()

		About := widgets.NewParagraph()
		About.Title = " Gotop "
		About.TitleStyle.Fg = ui.ColorCyan
		About.Border = true
		About.BorderStyle.Fg = ui.ColorBlue
		About.Text =
			"\nA system profiler written purely in golang!\n\n" +
				"version: " + gotopVersion + "\n\n" +
				"Made by krishpranav\n\n"

		uiEvents := ui.PollEvents()
		t := time.NewTicker(100 * time.Millisecond)
		tick := t.C

		for {
			select {
			case e := <-uiEvents:
				switch e.ID {
				case "q", "<C-c>":
					return
				}
			case <-tick:
				ui.Clear()
				w, h := ui.TerminalDimensions()
				About.SetRect((w-35)/2, (h-10)/2, (w+35)/2, (h+10)/2)
				ui.Render(About)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}
