package images

import (
	"image"
	"strings"
)

type ImageSize struct {
	Bound         image.Rectangle
	Width, Height int
}

type TransformationType string

const (
	Rebaixado TransformationType = "rebaixado"
	Pixelado  TransformationType = "pixelado"
)

func NewTransformationType(tt string) TransformationType {

	switch strings.ToLower(tt) {
	case "rebaixado":
		return Rebaixado
	case "pixelado":
		return Pixelado
	default:
		return ""
	}

}
