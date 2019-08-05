package service

type Service interface {
	Start()
	Stop() error
	Started() bool
}
