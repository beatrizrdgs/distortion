package images

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

type service struct {
}

func NewService() Handler {

	s := &service{}

	return s

}

func (s *service) Transform(img image.Image, transformation TransformationType) (image.Image, error) {

	switch transformation {

	case Rebaixado:
		img = resize.Resize(uint(img.Bounds().Dx()*5), uint(img.Bounds().Dy()), img, resize.Lanczos3)
	case Pixelado:
		pixelSize := 15
		img = resize.Resize(uint(img.Bounds().Dx()/pixelSize), uint(img.Bounds().Dy()/pixelSize), img, resize.NearestNeighbor)
	default:
		return nil, ErrInvalidTransformationType
	}

	return img, nil

}

func (s *service) EncodeImage(img image.Image, format string, buffer *bytes.Buffer) error {

	switch format {
	case "png":
		png.Encode(buffer, img)
	case "jpeg", "jpg":
		jpeg.Encode(buffer, img, nil)
	default:
		return ErrUnsupportedImageFormat
	}

	return nil

}

func (s *service) Resize(img image.Image, width, height uint) image.Image {
	return nil
}

func (s *service) Save(path string, img image.Image) error {

	b := bytes.NewBuffer(nil)

	err := jpeg.Encode(b, img, nil)
	if err != nil {
		return err
	}

	newImg, err := os.Create(path)
	if err != nil {
		return err
	}
	defer newImg.Close()

	_, err = newImg.Write(b.Bytes())
	if err != nil {
		return err
	}

	return nil

}
