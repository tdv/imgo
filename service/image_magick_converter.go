package service

import (
	"github.com/quirkey/magick"
)

type imageMagickConverter struct {
	Converter
}

func (this *imageMagickConverter) Convert(buf []byte) ([]byte, string, error) {
	if image, err := magick.NewFromBlob(buf, ""); err != nil {
		return err
	} else if err = image.Resize("200x100"); err != nil {
		return err
	}
	return nil, "", nil
}

func CreateImageMagickConverter() (Converter, error) {
	return &imageMagickConverter{}, nil
}
