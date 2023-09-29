package routes

import (
	"github.com/BetterToPractice/go-gin-setup/docs"
	"github.com/BetterToPractice/go-gin-setup/lib"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type SwaggerRouter struct {
	config  lib.Config
	handler lib.HttpHandler
}

func NewSwaggerRouter(config lib.Config, handler lib.HttpHandler) SwaggerRouter {
	return SwaggerRouter{
		config:  config,
		handler: handler,
	}
}

func (r SwaggerRouter) Setup() {
	docs.SwaggerInfo.Title = r.config.Swagger.Title
	docs.SwaggerInfo.Description = r.config.Swagger.Description
	docs.SwaggerInfo.Version = r.config.Swagger.Version

	r.handler.Engine.GET(r.config.Swagger.PathUrl, ginSwagger.WrapHandler(swaggerfiles.Handler))
}
