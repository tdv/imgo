package service

type Builder interface {
	Build() (interface{}, error)
}
