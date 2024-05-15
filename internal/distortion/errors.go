package distortion

import (
	"errors"
)

var (
	errInvalidTransformationType = errors.New("invalid transformation type")
	errInvalidOrientation        = errors.New("invalid orientation")
	errUnsupportedImageFormat    = errors.New("unsupported image format")
)
