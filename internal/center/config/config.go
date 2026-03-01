package config

import (
	_ "embed"
	"io"
	"log"
	"strings"

	registry "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	kregistry "github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"gopkg.in/yaml.v3"
)

//go:embed config.yaml
var bsc string

type Config struct {
	App   AppConf   `yaml:"app"`
	DB    DBConf    `yaml:"database"`
	Redis RedisConf `yaml:"redis"`
	Etcd  EtcdConf  `yaml:"etcd"`
}

func NewConfig() *Config {
	conf := &Config{}
	confContent, err := io.ReadAll(strings.NewReader(bsc))
	if err != nil {
		log.Fatalf("[config] failed bootstrap.yaml file not found:%s", err.Error())
	}

	err = yaml.Unmarshal(confContent, conf)
	if err != nil {
		log.Fatalf("[config] failed nacos config unmarshal error %s", err.Error())
	}
	return conf
}

var ProviderSet = wire.NewSet(
	NewConfig,
	NewEtcdRegister,
	wire.Bind(new(kregistry.Discovery), new(*registry.Registry)),
	wire.Bind(new(kregistry.Registrar), new(*registry.Registry)),
	NewDBConf,
)
