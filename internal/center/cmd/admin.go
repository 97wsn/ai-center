package cmd

import (
	"fmt"

	"app/internal/center/server"

	"github.com/urfave/cli/v2"
)

type AdminAction struct {
	Server *server.AdminServer
}

func (a *AdminAction) Handle(c *cli.Context) error {
	return a.Server.Run(c.Context, c.String("addr"))
}

func NewAdmin(action *AdminAction) *cli.Command {
	return &cli.Command{
		Name:  "admin",
		Usage: "admin command eg: ./app admin --addr=:8080",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "addr",
				Usage:    "--addr=:8080",
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
