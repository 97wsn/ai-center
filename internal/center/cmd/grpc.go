package cmd

import (
	"fmt"

	"app/internal/center/server"

	"github.com/urfave/cli/v2"
)

type RPCAction struct {
	Server *server.GRPCServer
}

func (a *RPCAction) Handle(c *cli.Context) error {
	return a.Server.Run(c.String("addr"))
}

func NewRPC(action *RPCAction) *cli.Command {
	return &cli.Command{
		Name:  "grpc",
		Usage: "grpc command eg: ./app grpc --addr=:3000",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "addr",
				Usage:    "--addr=:3000",
				Required: true,
			},
		},
		Before: func(ctx *cli.Context) error {
			fmt.Println("初始化配置信息...")
			//初始化nacos
			return nil
		},
		Action: func(ctx *cli.Context) error {
			return action.Handle(ctx)
		},
	}
}
