package service

type Converter interface {
	Convert(buf []byte) ([]byte, string, error)
}
