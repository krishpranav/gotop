package general

import (
	"context"
	"sync"

	"github.com/pesos/grofer/src/utils"
)

type serveFunc func(context.Context, chan utils.DataStats) error
