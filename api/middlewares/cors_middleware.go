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

func (m CorsMiddleware) Setup() {
	m.handler.Engine.Use(
		cors.New(cors.Config{
			AllowOrigins:  m.config.Cors.AllowOrigins,
			AllowMethods:  m.config.Cors.AllowMethods,
			AllowHeaders:  m.config.Cors.AllowHeaders,
			AllowWildcard: m.config.Cors.AllowWildcard,
		}),
	)
}
