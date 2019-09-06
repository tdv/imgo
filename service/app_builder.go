package service

import (
	"errors"

	"github.com/sarulabs/di"
)

type appBuilder struct {
	Builder
	config  Config
	builder *di.Builder
}

func (this *appBuilder) Build() (interface{}, error) {
	ctx := this.builder.Build()
	service := ctx.Get(this.config.GetStrVal(configPath(EntityServer, "active"))).(Service)
	return service, nil
}

var mainBuilderInst *appBuilder = nil

type buildContext struct {
	BuildContext

	entity    string
	impl      string
	config    Config
	container di.Container
}

func (this *buildContext) GetConfig() Config {
	return this.config.GetBranch(configPath(this.entity, this.impl))
}

func (this *buildContext) GetEntity(id string) interface{} {
	return this.container.Get(this.config.GetStrVal(configPath(id, "active")))
}

func RegisterEntity(entity, impl string, creator func(BuildContext) (interface{}, error)) error {
	if mainBuilderInst == nil {
		mainBuilderInst = &appBuilder{}
	}

	if mainBuilderInst.builder == nil {
		if builder, err := di.NewBuilder(); err != nil {
			panic("Failed to create application builder. Error: " + err.Error())
		} else {
			mainBuilderInst.builder = builder
		}
	}

	mainBuilderInst.builder.Add(di.Def{
		Name:  impl,
		Scope: di.App,
		Build: func(ctx di.Container) (interface{}, error) {
			if mainBuilderInst.config == nil {
				return nil, errors.New("Application builder was not initialized.")
			}
			return creator(&buildContext{
				entity:    entity,
				impl:      impl,
				config:    mainBuilderInst.config,
				container: ctx,
			})
		},
	})

	return nil
}

func InitAppBuilder(config Config) (Builder, error) {
	if mainBuilderInst.config != nil {
		return nil, errors.New("You can't twice initialize the application builder.")
	}

	mainBuilderInst.config = config

	return mainBuilderInst, nil
}
