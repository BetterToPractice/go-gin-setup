package middlewares

import "github.com/BetterToPractice/go-gin-setup/lib"
import "github.com/gin-contrib/cors"

type CorsMiddleware struct {
	handler lib.HttpHandler
	config  lib.Config
}

func NewCorsMiddleware(config lib.Config, handler lib.HttpHandler) CorsMiddleware {
	return CorsMiddleware{
		handler: handler,
		config:  config,
	}
}

func (c CorsMiddleware) Setup() {
	c.handler.Engine.Use(
		cors.New(cors.Config{
			AllowOrigins:  c.config.Cors.AllowOrigins,
			AllowMethods:  c.config.Cors.AllowMethods,
			AllowHeaders:  c.config.Cors.AllowHeaders,
			AllowWildcard: c.config.Cors.AllowWildcard,
		}),
	)
}
