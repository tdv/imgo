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

	config := config{config: cfg.GetStringMap("config")}

	if config.config == nil {
		return nil, errors.New("Failed to load config.")
	}

	return &config, nil
}
