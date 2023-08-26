package middlewares

import (
	"github.com/BetterToPractice/go-gin-setup/lib"
	helmet "github.com/danielkov/gin-helmet"
)

type HelmetMiddleware struct {
	handler lib.HttpHandler
}

func NewHelmetMiddleware(handler lib.HttpHandler) HelmetMiddleware {
	return HelmetMiddleware{
		handler: handler,
	}
}

func (m HelmetMiddleware) Setup() {
	m.handler.Engine.Use(helmet.Default())
}
