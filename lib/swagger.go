package lib

import (
	"github.com/BetterToPractice/go-gin-setup/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Swagger struct {
	config  Config
	handler HttpHandler
}

func NewSwagger(config Config, handler HttpHandler) Swagger {
	return Swagger{
		config:  config,
		handler: handler,
	}
}

func (l Swagger) SetUrl() {
	l.handler.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func (l Swagger) Setup() {
	docs.SwaggerInfo.Title = l.config.Swagger.Title
	docs.SwaggerInfo.Description = l.config.Swagger.Description
	docs.SwaggerInfo.Version = l.config.Swagger.Version

	l.SetUrl()
}
