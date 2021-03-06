package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/client"
	containerGraph "github.com/krishpranav/gotop/src/display/container"
	"github.com/krishpranav/gotop/src/utils"

	"github.com/krishpranav/gotop/src/container"
	"github.com/krishpranav/gotop/src/general"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

const (
	defaultCid                  = ""
	defaultContainerRefreshRate = 1000
)

var containerCmd = &cobra.Command{
	Use:     "container",
	Short:   "container command is used to get information related to docker containers",
	Long:    `container command is used to get information related to docker containers. It provides both overall and per container metrics.`,
	Aliases: []string{"containers", "docker"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return fmt.Errorf("the container command should have no arguments, see gotop container --help for further info")
		}

		cid, _ := cmd.Flags().GetString("container-id")
		containerRefreshRate, _ := cmd.Flags().GetUint64("refresh")

		if containerRefreshRate < 1000 {
			return fmt.Errorf("invalid refresh rate: minimum refresh rate is 1000(ms)")
		}

		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			return err
		}

		eg, ctx := errgroup.WithContext(context.Background())

		if cid != defaultCid {
			dataChannel := make(chan container.PerContainerMetrics, 1)

			eg.Go(func() error {
				return container.ServeContainer(ctx, cli, cid, dataChannel, int64(containerRefreshRate))
			})
			eg.Go(func() error {
				return containerGraph.ContainerVisuals(ctx, dataChannel, containerRefreshRate)
			})

			if err := eg.Wait(); err != nil {
				if err == general.ErrInvalidContainer {
					utils.ErrorMsg("cid")
				}
				if err != general.ErrCanceledByUser {
					log.Fatalf("Error: %v\n", err)
				}
			}
		} else {
			dataChannel := make(chan container.ContainerMetrics)

			allFlag, _ := cmd.Flags().GetBool("all")
			eg.Go(func() error {
				return container.Serve(ctx, cli, allFlag, dataChannel, int64(containerRefreshRate))
			})
			eg.Go(func() error {
				return containerGraph.OverallVisuals(ctx, cli, allFlag, dataChannel, containerRefreshRate)
			})

			if err := eg.Wait(); err != nil {
				if err != general.ErrCanceledByUser {
					log.Fatalf("Error: %v\n", err)
				}
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(containerCmd)

	containerCmd.Flags().StringP(
		"container-id",
		"c",
		"",
		"specify container ID",
	)

	containerCmd.Flags().Uint64P(
		"refresh",
		"r",
		defaultContainerRefreshRate,
		"Container information UI refreshes rate in milliseconds greater than 1000",
	)

	containerCmd.Flags().BoolP(
		"all",
		"a",
		false,
		"Specify to list all containers or only running containers.",
	)
}
