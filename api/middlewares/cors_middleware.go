package middlewares

import "github.com/BetterToPractice/go-gin-setup/lib"
import "github.com/gin-contrib/cors"

type CorsMiddleware struct {
	handler lib.HttpHandler
}

func NewCorsMiddleware(handler lib.HttpHandler) CorsMiddleware {
	return CorsMiddleware{
		handler: handler,
	}
}

func (c CorsMiddleware) Setup() {
	c.handler.Engine.Use(
		cors.New(cors.Config{
			AllowOrigins:  []string{"*"},
			AllowMethods:  []string{"*"},
			AllowHeaders:  []string{"*"},
			AllowWildcard: true,
		}),
	)
}
