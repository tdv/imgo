package service

type BuildContext interface {
	GetConfig() Config
	GetEntity(id string) interface{}
}
