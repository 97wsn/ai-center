package config

import (
	_ "embed"
)

type AppConf struct {
	Env string `yaml:"env"`
}
