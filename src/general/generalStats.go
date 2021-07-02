package general

import (
	"context"
	"sync"

	"github.com/krishpranav/gotop/src/utils"
)

type serveFunc func(context.Context, chan utils.DataStats) error

func GlobalStats(ctx context.Context, dataChannel chan utils.DataStats, refreshRate uint64) error {

	serveFuncs := []serveFunc{
		ServeCPURates,
		ServeMemRates,
		ServeDiskRates,
		ServeNetRates,
	}

	return utils.TickUntilDone(ctx, int64(refreshRate), func() error {
		var wg sync.WaitGroup

		errCh := make(chan error, len(serveFuncs))

		for _, sf := range serveFuncs {
			wg.Add(1)
			go func(sf serveFunc, dc chan utils.DataStats) {
				defer wg.Done()
				errCh <- sf(ctx, dc)
			}(sf, dataChannel)
		}

		wg.Wait()
		close(errCh)
		for err := range errCh {
			if err != nil {
				return err
			}
		}

		return nil
	})
}
