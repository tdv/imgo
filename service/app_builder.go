package service

import (
	"github.com/sarulabs/di"
)

type appBuilder struct {
	Builder
	builder *di.Builder
}

func (this *appBuilder) init() error {
	if builder, err := di.NewBuilder(); err != nil {
		return err
	} else {
		if err := builder.Add(
			[]di.Def{
				{
					Name:  "postgres-storage",
					Scope: di.App,
					Build: func(ctx di.Container) (interface{}, error) {
						return CreatePostgresStorage()
					},
				},
				{
					Name:  "redis-cache",
					Scope: di.App,
					Build: func(ctx di.Container) (interface{}, error) {
						return CreateRedisCache()
					},
				},
				{
					Name:  "image-converter",
					Scope: di.App,
					Build: func(ctx di.Container) (interface{}, error) {
						return CreateImageMagickConverter()
					},
				},
				{
					Name:  "storage",
					Scope: di.App,
					Build: func(ctx di.Container) (interface{}, error) {
						// TODO: select from config
						return ctx.Get("postgres-storage"), nil
					},
				},
				{
					Name:  "cache",
					Scope: di.App,
					Build: func(ctx di.Container) (interface{}, error) {
						// TODO: select from config
						return ctx.Get("redis-cache"), nil
					},
				},
				{
					Name:  "http-server",
					Scope: di.App,
					Build: func(ctx di.Container) (interface{}, error) {
						return CreateHttpServer(
							ctx.Get("image-converter").(Converter),
							ctx.Get("storage").(Storage),
							ctx.Get("cache").(Storage),
						)
					},
				},
			}...,
		); err != nil {
			return err
		} else {
			this.builder = builder
		}
	}

	return nil
}

func (this *appBuilder) Build() (interface{}, error) {
	ctx := this.builder.Build()
	service := ctx.Get("http-server").(Service)
	return service, nil
}

func CreateAppBuilder() (Builder, error) {
	builder := appBuilder{}

	if err := builder.init(); err != nil {
		return nil, err
	}

	return &builder, nil
}
