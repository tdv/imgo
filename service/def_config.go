package service

import (
	"errors"
	"strings"

	"github.com/spf13/viper"
)

type config struct {
	Config
	config map[string]interface{}
}

func setDefault(config *viper.Viper) {
	config.SetDefault(ConfigPath(EntityServer, "active"), ImplHttp)

	config.SetDefault(ConfigPath(EntityServer, ImplHttp, "address"), "0.0.0.0:55555")

	config.SetDefault(ConfigPath(EntityImageConverter, "active"), ImplImageMagick)

	config.SetDefault(ConfigPath(EntityImageConverter, ImplImageMagick, "format"), "png")
	config.SetDefault(ConfigPath(EntityImageConverter, ImplImageMagick, "size.default.width"), 300)
	config.SetDefault(ConfigPath(EntityImageConverter, ImplImageMagick, "size.default.height"), 200)
	config.SetDefault(ConfigPath(EntityImageConverter, ImplImageMagick, "size.max.width"), 4096)
	config.SetDefault(ConfigPath(EntityImageConverter, ImplImageMagick, "size.max.height"), 4096)

	config.SetDefault(ConfigPath(EntityImageConverter, ImplStdImage, "format"), "png")
	config.SetDefault(ConfigPath(EntityImageConverter, ImplStdImage, "size.default.width"), 300)
	config.SetDefault(ConfigPath(EntityImageConverter, ImplStdImage, "size.default.height"), 200)
	config.SetDefault(ConfigPath(EntityImageConverter, ImplStdImage, "size.max.width"), 4096)
	config.SetDefault(ConfigPath(EntityImageConverter, ImplStdImage, "size.max.height"), 4096)

	config.SetDefault(ConfigPath(EntityStorage, "active"), ImplPostgres)

	config.SetDefault(ConfigPath(EntityStorage, ImplPostgres, "host"), "localhost")
	config.SetDefault(ConfigPath(EntityStorage, ImplPostgres, "port"), 5432)
	config.SetDefault(ConfigPath(EntityStorage, ImplPostgres, "dbname"), "imgo")
	config.SetDefault(ConfigPath(EntityStorage, ImplPostgres, "sslmode"), "disable")
	config.SetDefault(ConfigPath(EntityStorage, ImplPostgres, "user"), "postgres")
	config.SetDefault(ConfigPath(EntityStorage, ImplPostgres, "password"), "")

	config.SetDefault(ConfigPath(EntityStorage, ImplMySql, "host"), "localhost")
	config.SetDefault(ConfigPath(EntityStorage, ImplMySql, "port"), 3306)
	config.SetDefault(ConfigPath(EntityStorage, ImplMySql, "dbname"), "imgo")
	config.SetDefault(ConfigPath(EntityStorage, ImplMySql, "user"), "postgres")
	config.SetDefault(ConfigPath(EntityStorage, ImplMySql, "password"), "mysql")

	config.SetDefault(ConfigPath(EntityCache, "active"), ImplRedis)

	config.SetDefault(ConfigPath(EntityCache, ImplRedis, "address"), "localhost:6379")
	config.SetDefault(ConfigPath(EntityCache, ImplRedis, "db"), 0)
	config.SetDefault(ConfigPath(EntityCache, ImplRedis, "password"), "")
	config.SetDefault(ConfigPath(EntityCache, ImplRedis, "expiration"), 15)

	config.SetDefault(ConfigPath(EntityCache, ImplMemcached, "nodes"), "localhost:11211")
	config.SetDefault(ConfigPath(EntityCache, ImplMemcached, "expiration"), 15)
}

func (this *config) findNode(path string) interface{} {
	items := strings.Split(path, ".")

	size := len(items)

	if size < 1 {
		panic("Failed to parse path for getting value.")
	}

	var node interface{} = nil

	var partialPath string = ""

	for i, key := range items {
		if i != 0 {
			partialPath = partialPath + "."
		}

		partialPath = partialPath + key

		if node == nil {
			node = this.config[key]
		} else {
			node = node.(map[string]interface{})[key]
		}

		if node == nil {
			panic("Config node \"" + partialPath + "\" not found.")
		}

		if i == size-1 {
			if node != nil {
				return node
			} else {
				panic("Config node \"" + partialPath + "\" not found.")
			}
		}
	}

	panic("Config item not found by path \"" + path + "\".")
}

func (this *config) GetStrVal(path string) string {
	val := this.findNode(path)
	if str, ok := val.(string); ok {
		return str
	} else {
		panic("Required config value from \"" + path + "\" is not string.")
	}
}

func (this *config) GetIntVal(path string) int {
	val := this.findNode(path)
	if i, ok := val.(float64); ok {
		return int(i)
	} else {
		panic("Required config value from \"" + path + "\" is not int.")
	}
}

func (this *config) GetBranch(path string) Config {
	val := this.findNode(path)
	if cfg, ok := val.(map[string]interface{}); ok {
		return &config{config: cfg}
	} else {
		panic("Required config value from \"" + path + "\" is not config branch.")
	}
}

func LoadDefConfig() (Config, error) {
	cfg := viper.New()

	cfg.SetConfigName("config")
	cfg.AddConfigPath("/etc/imgo/")
	cfg.AddConfigPath("$HOME/.imgo")
	cfg.AddConfigPath(".")

	if err := cfg.ReadInConfig(); err != nil {
		return nil, err
	}

	setDefault(cfg)

	config := config{config: cfg.GetStringMap("config")}

	if config.config == nil {
		return nil, errors.New("Failed to load config.")
	}

	return &config, nil
}
