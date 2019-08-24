package service

// For Ubuntu 16.04 : sudo apt-get install libmagickwand-dev

import (
	"crypto/sha1"
	"encoding/hex"

	"strconv"

	"github.com/quirkey/magick"
	"github.com/spf13/viper"
)

type imageMagickConverter struct {
	Converter

	defFormat string

	defWidth  int
	defHeight int

	maxWidth  int
	maxHeight int
}

func calcLen(len *int, def int, max int) string {
	var l int
	if len != nil {
		if *len <= 0 {
			l = def
		} else if *len > max {
			l = max
		} else {
			l = *len
		}
	} else {
		l = def
	}
	return strconv.Itoa(l)
}

func (this *imageMagickConverter) Convert(buf []byte, format string, width *int, height *int) ([]byte, string, error) {
	var image *magick.MagickImage
	var err error
	if image, err = magick.NewFromBlob(buf, format); err != nil {
		return nil, "", err
	}

	defer image.Destroy()

	w := calcLen(width, this.defWidth, this.maxWidth)
	h := calcLen(height, this.defHeight, this.maxHeight)

	if err = image.Resize(w + "x" + h); err != nil {
		return nil, "", err
	}

	var blob []byte
	if blob, err = image.ToBlob(this.defFormat); err != nil {
		return nil, "", err
	}

	hash := sha1.Sum(blob)
	id := hex.EncodeToString(hash[:])

	return blob, id, nil
}

func CreateImageMagickConverter(config *viper.Viper) (Converter, error) {
	return &imageMagickConverter{
		defFormat: config.GetString("imageconverter.format"),
		defWidth:  config.GetInt("imageconverter.size.default.width"),
		defHeight: config.GetInt("imageconverter.size.default.height"),
		maxWidth:  config.GetInt("imageconverter.size.max.width"),
		maxHeight: config.GetInt("imageconverter.size.max.height"),
	}, nil
}
