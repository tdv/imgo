package service

type imageMagickConverter struct {
	Converter
}

func (this *imageMagickConverter) Convert(buf []byte) ([]byte, string, error) {
	return nil, "", nil
}

func CreateImageMagickConverter() (Converter, error) {
	return &imageMagickConverter{}, nil
}
