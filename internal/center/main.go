package main

import (
	_ "embed"
	"log"
	"os"
	"time"

	"app/internal/center/cmd"

	"github.com/urfave/cli/v2"
)

type App struct {
	serv *cli.App
}

func NewApp() *App {
	return &App{
		serv: cli.NewApp(),
	}
}

func (h *App) Serv() *cli.App {
	return h.serv
}

func (h *App) Register(cmd ...*cli.Command) {
	h.serv.Commands = append(h.serv.Commands, cmd...)
}

func (h *App) Run(args []string) {
	app := h.serv
	app.Name = "app"
	app.Version = "1.0"
	app.Copyright = "97Wsn"
	app.Compiled = time.Now()
	app.Writer = os.Stdout
	cli.ErrWriter = os.Stdout
	if err := app.Run(args); err != nil {
		log.Fatalf("[APP] application run error: %s", err)
	}
}

func main() {
	adminAction, adminFunc := NewAdminAction()
	defer adminFunc()

	rpcAction, rpcFunc := NewRpcAction()
	defer rpcFunc()
	app := NewApp()
	app.Register(cmd.NewAdmin(adminAction))
	app.Register(cmd.NewRPC(rpcAction))

	app.Run(os.Args)

}
