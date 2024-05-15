package distortion

import (
	"image"
	"strings"
)

type imageSize struct {
	bound         image.Rectangle
	width, height int
}

type transformationType string

const (
	stretch  transformationType = "stretch"
	pixelate transformationType = "pixelate"
	jokerize transformationType = "jokerize"
	chuuify  transformationType = "chuuify"

	jokerPath = "assets/transform/joker.png"
	chuuPath  = "assets/transform/chuu.png"
)

func newTransformationType(tt string) transformationType {
	switch strings.ToLower(tt) {
	case "stretch":
		return stretch
	case "pixelate":
		return pixelate
	case "jokerize":
		return jokerize
	case "chuuify":
		return chuuify
	default:
		return ""
	}
}
