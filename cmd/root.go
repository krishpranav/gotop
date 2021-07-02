package cmd

import (
	"context"
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	overallGraph "github.com/krishpranav/gotop/src/display/general"
	"github.com/krishpranav/gotop/src/general"
	info "github.com/krishpranav/gotop/src/general"
	"github.com/krishpranav/gotop/src/utils"
	"golang.org/x/sync/errgroup"
)

const (
	defaultOverallRefreshRate = 1000
	defaultConfigFileLocation = ""
	defaultCPUBehavior        = false
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gotop",
	Short: "gotop is a system and resource monitor written in golang",
	Long: `gotop is a system and resource monitor written in golang.
While using a TUI based command, press ? to get information about key bindings (if any) for that command.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		overallRefreshRate, _ := cmd.Flags().GetUint64("refresh")

		if overallRefreshRate < 1000 {
			return fmt.Errorf("invalid refresh rate: minimum refresh rate is 1000(ms)")
		}

		cpuLoadFlag, _ := cmd.Flags().GetBool("cpuinfo")
		if cpuLoadFlag {
			cpuLoad := info.NewCPULoad()
			dataChannel := make(chan *info.CPULoad, 1)

			eg, ctx := errgroup.WithContext(context.Background())

			eg.Go(func() error {
				return info.GetCPULoad(ctx, cpuLoad, dataChannel, uint64(4*overallRefreshRate/5))
			})

			eg.Go(func() error {
				return overallGraph.RenderCPUinfo(ctx, dataChannel, overallRefreshRate)
			})

			if err := eg.Wait(); err != nil {
				if err != general.ErrCanceledByUser {
					fmt.Printf("Error: %v\n", err)
				}
			}

		} else {
			dataChannel := make(chan utils.DataStats, 1)

			eg, ctx := errgroup.WithContext(context.Background())

			eg.Go(func() error {
				return general.GlobalStats(ctx, dataChannel, uint64(4*overallRefreshRate/5))
			})
			eg.Go(func() error {
				return overallGraph.RenderCharts(ctx, dataChannel, overallRefreshRate)
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		defaultConfigFileLocation,
		"config file (default is $HOME/.gotop.yaml)",
	)

	rootCmd.Flags().Uint64P(
		"refresh",
		"r",
		defaultOverallRefreshRate,
		"Overall stats UI refreshes rate in milliseconds greater than 1000",
	)

	rootCmd.Flags().BoolP(
		"cpuinfo",
		"c",
		defaultCPUBehavior,
		"Info about the CPU Load over all CPUs",
	)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".gotop")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
