package service

// For Ubuntu 16.04 : sudo apt-get install libmagickwand-dev

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/quirkey/magick"
)

type imageMagickConverter struct {
	Converter
}

func (this *imageMagickConverter) Convert(buf []byte) ([]byte, string, error) {
	var image *magick.MagickImage
	var err error
	if image, err = magick.NewFromBlob(buf, "jpg"); err != nil {
		return nil, "", err
	}

	if err = image.Resize("200x100"); err != nil {
		return nil, "", err
	}

	var blob []byte
	if blob, err = image.ToBlob("png"); err != nil {
		return nil, "", err
	}

	hash := sha1.Sum(blob)
	id := hex.EncodeToString(hash[:])

	return blob, id, nil
}

func CreateImageMagickConverter() (Converter, error) {
	return &imageMagickConverter{}, nil
}
