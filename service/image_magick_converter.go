package service

// For Ubuntu 16.04 : sudo apt-get install libmagickwand-dev

import (
	"crypto/sha1"
	"encoding/hex"

	"strconv"

	"github.com/quirkey/magick"
)

type imageMagickConverter struct {
	Converter

	defFormat string

	defWidth  int
	defHeight int

	maxWidth  int
	maxHeight int
}

func (this *imageMagickConverter) calcLen(len *int, def int, max int) string {
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

	w := this.calcLen(width, this.defWidth, this.maxWidth)
	h := this.calcLen(height, this.defHeight, this.maxHeight)

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

// ImplImageMagick - id of ImageMagick based implementation of the Converter interface
const ImplImageMagick = "imagemagick"

var _ = RegisterEntity(
	EntityImageConverter,
	ImplImageMagick,
	func(ctx BuildContext) (interface{}, error) {
		config := ctx.GetConfig()
		return &imageMagickConverter{
			defFormat: config.GetStrVal("format"),
			defWidth:  config.GetIntVal("size.default.width"),
			defHeight: config.GetIntVal("size.default.height"),
			maxWidth:  config.GetIntVal("size.max.width"),
			maxHeight: config.GetIntVal("size.max.height"),
		}, nil
	},
)
