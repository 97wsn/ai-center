package service

import (
	"app/internal/center/service/admin"
	"app/internal/center/service/rpc"

	"github.com/google/wire"
)

type RpcService = rpc.Service
type AdminService = admin.Service

var ProviderSet = wire.NewSet(
	rpc.ProviderSet,
	admin.ProviderSet,
)
