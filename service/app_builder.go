package service

import (
	"github.com/sarulabs/di"
)

type appBuilder struct {
	Builder
	severName string
	builder   *di.Builder
}

func (this *appBuilder) init(config Config) error {
	this.severName = config.GetStrVal(ConfigPath(EntityServer, "active"))

	if builder, err := di.NewBuilder(); err != nil {
		return err
	} else {
		if err := builder.Add(
			[]di.Def{
				{
					Name:  ImplPostgres,
					Scope: di.App,
					Build: func(ctx di.Container) (interface{}, error) {
						cfg := config.GetBranch(ConfigPath(EntityStorage, ImplPostgres))
						return CreatePostgresStorage(cfg)
					},
				},
				{
					Name:  ImplMySql,
					Scope: di.App,
					Build: func(ctx di.Container) (interface{}, error) {
						cfg := config.GetBranch(ConfigPath(EntityStorage, ImplMySql))
						return CreateMySqlStorage(cfg)
					},
				},
				{
					Name:  ImplRedis,
					Scope: di.App,
					Build: func(ctx di.Container) (interface{}, error) {
						cfg := config.GetBranch(ConfigPath(EntityCache, ImplRedis))
						return CreateRedisCache(cfg)
					},
				},
				{
					Name:  ImplMemcached,
					Scope: di.App,
					Build: func(ctx di.Container) (interface{}, error) {
						cfg := config.GetBranch(ConfigPath(EntityCache, ImplMemcached))
						return CreateMemcachedCache(cfg)
					},
				},
				{
					Name:  ImplImageMagick,
					Scope: di.App,
					Build: func(ctx di.Container) (interface{}, error) {
						cfg := config.GetBranch(ConfigPath(EntityImageConverter, ImplImageMagick))
						return CreateImageMagickConverter(cfg)
					},
				},
				{
					Name:  ImplStdImage,
					Scope: di.App,
					Build: func(ctx di.Container) (interface{}, error) {
						cfg := config.GetBranch(ConfigPath(EntityImageConverter, ImplStdImage))
						return CreateStdImageConverter(cfg)
					},
				},
				{
					Name:  ImplHttp,
					Scope: di.App,
					Build: func(ctx di.Container) (interface{}, error) {
						return CreateHttpServer(
							config.GetBranch(ConfigPath(EntityServer, ImplHttp)),
							ctx.Get(config.GetStrVal(ConfigPath(EntityImageConverter, "active"))).(Converter),
							ctx.Get(config.GetStrVal(ConfigPath(EntityStorage, "active"))).(Storage),
							ctx.Get(config.GetStrVal(ConfigPath(EntityCache, "active"))).(Storage),
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
	service := ctx.Get(this.severName).(Service)
	return service, nil
}

func CreateAppBuilder(config Config) (Builder, error) {
	builder := appBuilder{}

	if err := builder.init(config); err != nil {
		return nil, err
	}

	return &builder, nil
}
