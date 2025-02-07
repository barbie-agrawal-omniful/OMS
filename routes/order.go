package routes

import (
	"context"

	"oms/controllers"

	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/http"
	"github.com/omniful/go_commons/log"
)

func PublicRoutes(ctx context.Context, server *http.Server) error {
	OrderGroup := server.Engine.Group("/api/v1", log.RequestLogMiddleware(log.MiddlewareOptions{
		Format:      config.GetString(ctx, "log.format"),
		Level:       config.GetString(ctx, "log.level"),
		LogRequest:  config.GetBool(ctx, "log.request"),
		LogResponse: config.GetBool(ctx, "log.response"),
	}))

	OrderGroup.POST("", controllers.CreateOrder)
	OrderGroup.GET("/orders", controllers.ViewOrders)

	return nil
}
