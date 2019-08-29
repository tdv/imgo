package service

type Config interface {
	GetStrVal(path string) string
	GetIntVal(path string) int
	GetBranch(path string) Config
}
