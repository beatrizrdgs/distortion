package images

import (
	"bytes"
	"image"
)

type Handler interface {
	EncodeImage(img image.Image, format string, buffer *bytes.Buffer) error
	Transform(img image.Image, transformation TransformationType) (image.Image, error)
	Resize(img image.Image, width, height uint) image.Image
}
