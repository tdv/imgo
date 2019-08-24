package service

import (
	"github.com/spf13/viper"
)

func setDefault(config *viper.Viper) {
	config.SetDefault("server.address", "0.0.0.0:55555")

	config.SetDefault("imageconverter.format", "png")
	config.SetDefault("imageconverter.size.default.width", 300)
	config.SetDefault("imageconverter.size.default.height", 200)
	config.SetDefault("imageconverter.size.max.width", 4096)
	config.SetDefault("imageconverter.size.max.height", 4096)

	config.SetDefault("storage.active", "postgres")

	config.SetDefault("storage.postgres.host", "localhost")
	config.SetDefault("storage.postgres.port", 5432)
	config.SetDefault("storage.postgres.dbname", "imgo")
	config.SetDefault("storage.postgres.sslmode", "disable")
	config.SetDefault("storage.postgres.user", "postgres")
	config.SetDefault("storage.postgres.password", "")

	config.SetDefault("cache.active", "redis")

	config.SetDefault("cache.redis.address", "localhost:6379")
	config.SetDefault("cache.redis.db", 0)
	config.SetDefault("cache.redis.password", "")
	config.SetDefault("cache.redis.expiration", 15)
}

func InitConfig() (*viper.Viper, error) {
	config := viper.New()

	config.SetConfigName("config")
	config.AddConfigPath("/etc/imgo/")
	config.AddConfigPath("$HOME/.imgo")
	config.AddConfigPath(".")

	if err := config.ReadInConfig(); err != nil {
		return nil, err
	}

	setDefault(config)

	return config, nil
}
