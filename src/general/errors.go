package general

import (
	"errors"
)

var ErrCanceledByUser = errors.New("canceled by user!!!")

var ErrInvalidContainer = errors.New("container does not exists!!!")
