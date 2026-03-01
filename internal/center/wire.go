//go:build wireinject
// +build wireinject

package main

import (
	"app/internal/center/cmd"
	"app/internal/center/config"
	"app/internal/center/server"
	"app/internal/center/service"

	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	// config
	config.ProviderSet,

	// http server
	server.NewAdminServer,

	// rpc server
	server.NewGRPCServer,

	// service
	service.ProviderSet,
)

func NewAdminAction() (*cmd.AdminAction, func()) {
	panic(wire.Build(providerSet, wire.Struct(new(cmd.AdminAction), "*")))
}

func NewRpcAction() (*cmd.RPCAction, func()) {
	panic(wire.Build(providerSet, wire.Struct(new(cmd.RPCAction), "*")))
}
