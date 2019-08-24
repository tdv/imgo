package service

import (
	"github.com/sarulabs/di"
	"github.com/spf13/viper"
)

type appBuilder struct {
	Builder
	builder *di.Builder
}

func (this *appBuilder) init(config *viper.Viper) error {
	if builder, err := di.NewBuilder(); err != nil {
		return err
	} else {
		if err := builder.Add(
			[]di.Def{
				{
					Name:  "postgres",
					Scope: di.App,
					Build: func(ctx di.Container) (interface{}, error) {
						return CreatePostgresStorage(config)
					},
				},
				{
					Name:  "redis",
					Scope: di.App,
					Build: func(ctx di.Container) (interface{}, error) {
						return CreateRedisCache(config)
					},
				},
				{
					Name:  "imageconverter",
					Scope: di.App,
					Build: func(ctx di.Container) (interface{}, error) {
						return CreateImageMagickConverter(config)
					},
				},
				{
					Name:  "http-server",
					Scope: di.App,
					Build: func(ctx di.Container) (interface{}, error) {
						return CreateHttpServer(
							config,
							ctx.Get("imageconverter").(Converter),
							ctx.Get(config.GetString("storage.active")).(Storage),
							ctx.Get(config.GetString("cache.active")).(Storage),
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

func CreateAppBuilder(config *viper.Viper) (Builder, error) {
	builder := appBuilder{}

	if err := builder.init(config); err != nil {
		return nil, err
	}

	return &builder, nil
}
