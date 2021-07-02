package container

import (
	"context"

	"github.com/docker/docker/client"
	"github.com/krishpranav/gotop/src/utils"
)

func Serve(ctx context.Context, cli *client.Client, all bool, dataChannel chan ContainerMetrics, refreshRate int64) error {
	return utils.TickUntilDone(ctx, refreshRate, func() error {
		metrics, err := GetOverallMetrics(ctx, cli, all)
		if err != nil {
			return err
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case dataChannel <- metrics:
		}

		return nil
	})
}

func ServeContainer(ctx context.Context, cli *client.Client, cid string, dataChannel chan PerContainerMetrics, refreshRate int64) error {
	return utils.TickUntilDone(ctx, refreshRate, func() error {
		metrics, err := GetContainerMetrics(ctx, cli, cid)
		if err != nil {
			return err
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case dataChannel <- metrics:
		}

		return nil
	})
}
