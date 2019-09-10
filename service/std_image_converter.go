package service

// For Ubuntu 16.04 : sudo apt-get install libmagickwand-dev

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"errors"

	"image"
	"image/gif"
	"image/jpeg"
	"image/png"

	_ "strconv"
	"strings"

	"github.com/nfnt/resize"
)

type stdImageConverter struct {
	Converter

	defFormat string

	defWidth  int
	defHeight int

	maxWidth  int
	maxHeight int
}

func (this *stdImageConverter) loadImage(buf []byte, format string) (image.Image, error) {
	if len(format) == 0 {
		return nil, errors.New("Needed not empty format parameter.")
	}

	reader := bytes.NewReader(buf)
	if reader == nil {
		return nil, errors.New("Failed to create reader from the input buffer.")
	}

	var img image.Image
	var err error

	switch f := strings.ToLower(format); f {
	case "gif":
		if img, err = gif.Decode(reader); err != nil {
			return nil, err
		}
	case "jpeg", "jpg":
		if img, err = jpeg.Decode(reader); err != nil {
			return nil, err
		}
	case "png":
		if img, err = png.Decode(reader); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Unsupported input format \"" + f + "\".")
	}

	return img, nil
}

func (this *stdImageConverter) storeImage(img image.Image) ([]byte, error) {
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)
	if writer == nil {
		return nil, errors.New("Failed to create writer.")
	}

	switch this.defFormat {
	case "gif":
		if err := gif.Encode(writer, img, nil); err != nil {
			return nil, err
		}
	case "jpeg", "jpg":
		if err := jpeg.Encode(writer, img, nil); err != nil {
			return nil, err
		}
	case "png":
		if err := png.Encode(writer, img); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Unsupported output format \"" + this.defFormat + "\".")
	}

	return buffer.Bytes(), nil
}

func (this *stdImageConverter) calcLen(len *int, def int, max int) uint {
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
	return uint(l)
}

func (this *stdImageConverter) Convert(buf []byte, format string, width *int, height *int) ([]byte, string, error) {
	var img image.Image
	var err error

	if img, err = this.loadImage(buf, format); err != nil {
		return nil, "", err
	}

	w := this.calcLen(width, this.defWidth, this.maxWidth)
	h := this.calcLen(height, this.defHeight, this.maxHeight)

	img = resize.Resize(w, h, img, resize.Lanczos3)

	var blob []byte

	if blob, err = this.storeImage(img); err != nil {
		return nil, "", err
	}

	hash := sha1.Sum(blob)
	id := hex.EncodeToString(hash[:])

	return blob, id, nil
}

// ImplStdImage - id of standard based implementation of the Converter interface
const ImplStdImage = "std"

var _ = RegisterEntity(
	EntityImageConverter,
	ImplStdImage,
	func(ctx BuildContext) (interface{}, error) {
		config := ctx.GetConfig()
		return &stdImageConverter{
			defFormat: config.GetStrVal("format"),
			defWidth:  config.GetIntVal("size.default.width"),
			defHeight: config.GetIntVal("size.default.height"),
			maxWidth:  config.GetIntVal("size.max.width"),
			maxHeight: config.GetIntVal("size.max.height"),
		}, nil
	},
)
