package server

import (
	"log"
	"os"
	"time"

	"app/internal/center/service"

	center "github.com/97wsn/ai-center/api/center/rpc/v1"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

type GRPCServer struct {
	reg         registry.Registrar
	middlewares *middleware.Middleware
	service     *service.RpcService
}

func NewGRPCServer(reg registry.Registrar, service *service.RpcService) *GRPCServer {
	return &GRPCServer{reg: reg, service: service}
}

func (a *GRPCServer) Run(addr string) error {
	// gprc server
	grpcServ := grpc.NewServer(
		grpc.Address(addr),
		grpc.Timeout(time.Second*5),
		grpc.Middleware(
			// trace
			tracing.Server(),
			// metadata
			metadata.Server(metadata.WithPropagatedPrefix("x-very-")),
			// nolint validate
			validate.Validator(),
		),
	)
	// 服务注册
	center.RegisterUserServer(grpcServ, a.service.User)

	// create app
	app := kratos.New(
		kratos.Name(os.Getenv("APP_NAME")), // 应用名
		kratos.Version("v1.0.0"),           // 应用版本
		kratos.Registrar(a.reg),
		kratos.Server(grpcServ),
		kratos.Metadata(map[string]string{
			"start_time": time.Now().String(),
			"describe":   "ai应用中心RPC服务",
		}),
	)

	if err := app.Run(); err != nil {
		log.Printf("[APP] listen error: %s\n", err)
		return err
	}
	return nil
}
