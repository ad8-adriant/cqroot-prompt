package prompt

import (
	"errors"

	"github.com/ad8-adriant/cqroot-prompt/constants"
)

var (
	ErrModelConversion = errors.New("model conversion failed")
	ErrUserQuit        = constants.ErrUserQuit
)
