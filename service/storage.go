package service

type Storage interface {
	Put(id string, buf []byte) error
	Get(id string) ([]byte, error)
}
