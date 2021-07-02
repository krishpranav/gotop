package cmd

import (
	"context"
	"fmt"

	proc "github.com/shirou/gopsutil/process"
	"github.com/spf13/cobra"

	procGraph "github.com/krishpranav/gotop/src/display/process"
	"github.com/krishpranav/gotop/src/general"
	"github.com/krishpranav/gotop/src/process"
	"github.com/krishpranav/gotop/src/utils"
	"golang.org/x/sync/errgroup"
)

const (
	defaultProcRefreshRate = 3000
	defaultProcPid         = 0
)

// procCmd represents the proc command
var procCmd = &cobra.Command{
	Use:   "proc",
	Short: "proc command is used to get per-process information",
	Long: `proc command is used to get information about each running process in the system.
Syntax:
  gotop proc
To get information about a particular process whose PID is known the -p or --pid flag can be used.
Syntax:
  gotop proc -p [PID]`,
	Aliases: []string{"process", "processess"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return fmt.Errorf("the proc command should have no arguments, see gotop proc --help for further info")
		}

		pid, _ := cmd.Flags().GetInt32("pid")
		procRefreshRate, _ := cmd.Flags().GetUint64("refresh")

		if procRefreshRate < 1000 {
			return fmt.Errorf("invalid refresh rate: minimum refresh rate is 1000(ms)")
		}

		if pid != defaultProcPid {
			dataChannel := make(chan *process.Process, 1)

			eg, ctx := errgroup.WithContext(context.Background())

			proc, err := process.NewProcess(pid)
			if err != nil {
				utils.ErrorMsg("pid")
				return fmt.Errorf("invalid pid")
			}

			eg.Go(func() error {
				return process.Serve(proc, dataChannel, ctx, int64(4*procRefreshRate/5))
			})
			eg.Go(func() error {
				return procGraph.ProcVisuals(ctx, dataChannel, procRefreshRate)
			})

			if err := eg.Wait(); err != nil {
				if err != general.ErrCanceledByUser {
					fmt.Printf("Error: %v\n", err)
				}
			}
		} else {
			dataChannel := make(chan []*proc.Process, 1)

			eg, ctx := errgroup.WithContext(context.Background())

			eg.Go(func() error {
				return process.ServeProcs(dataChannel, ctx, int64(4*procRefreshRate/5))
			})
			eg.Go(func() error {
				return procGraph.AllProcVisuals(dataChannel, ctx, procRefreshRate)
			})

			if err := eg.Wait(); err != nil {
				if err != general.ErrCanceledByUser {
					fmt.Printf("Error: %v\n", err)
				}
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(procCmd)

	procCmd.Flags().Uint64P(
		"refresh",
		"r",
		defaultProcRefreshRate,
		"Process information UI refreshes rate in milliseconds greater than 1000",
	)

	procCmd.Flags().Int32P(
		"pid",
		"p",
		defaultProcPid,
		"specify PID of process. Passing PID 0 lists all the processes (same as not using the -p flag).",
	)
}
