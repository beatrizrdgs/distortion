package images

import (
	"bytes"
	"image"
	"image/jpeg"
	"os"
)

type Service struct {
}

func NewService() *Service {

	s := &Service{}

	return s

}

func (s *Service) Upload(path string) (image.Image, error) {
	return nil, nil
}

func (s *Service) Resize(img image.Image, width, height uint) image.Image {
	return nil
}

func (s *Service) Save(path string, img image.Image) error {

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

func (s *Service) Download(path string, img image.Image) error {
	return nil
}
