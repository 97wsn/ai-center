package server

import (
	"context"
	_ "embed"
	"net/http"

	svc "app/internal/center/service/admin"

	center "github.com/97wsn/ai-center/api/center/admin/v1"
	"github.com/gin-gonic/gin"

	"github.com/go-kratos/kratos/v2"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

type AdminServer struct {
	service *svc.Service
}

func NewAdminServer(service *svc.Service) *AdminServer {
	return &AdminServer{service: service}
}

func (r *AdminServer) Run(ctx context.Context, addr string) error {
	srv := khttp.NewServer(khttp.Address(addr))
	srv.HandlePrefix("/", r.router())
	center.RegisterUserHTTPServer(srv, r.service.User)

	app := kratos.New(
		kratos.Name("center-admin"),
		kratos.Server(srv),
	)
	return app.Run()
}

func (r *AdminServer) router() http.Handler {
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, "hello kratos")
	})

	return router
}
