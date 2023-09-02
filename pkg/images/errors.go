package images

import (
	"errors"
	"image-messer/pkg/cerrors"
	"net/http"
)

var (
	ErrInvalidTransformationType = errors.New("invalid transformation type")
	ErrParsingImage              = cerrors.NewCError("error parsing image", http.StatusBadRequest)
	ErrDecodeImage               = cerrors.NewCError("error decoding image", http.StatusInternalServerError)
	ErrUnsupportedImageFormat    = cerrors.NewCError("unsupported image format", http.StatusBadRequest)
)
