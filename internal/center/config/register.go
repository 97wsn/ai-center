package config

import (
	"app/internal/pkg/register"

	registry "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
)

func NewEtcdRegister(conf *Config) *registry.Registry {
	return register.NewEtcdRegistry(&register.EtcdConf{
		Env:       conf.App.Env,
		Endpoints: conf.Etcd.Endpoints,
		TlsCert:   "",
		TlsKey:    "",
		TlsCa:     "",
		Token:     "",
	})
}
