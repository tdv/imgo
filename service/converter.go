package service

type Converter interface {
	Convert(buf []byte, format string, width *int, height *int) ([]byte, string, error)
}
