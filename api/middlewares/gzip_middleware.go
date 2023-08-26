package middlewares

import (
	"github.com/BetterToPractice/go-gin-setup/lib"
	"github.com/gin-contrib/gzip"
)

type GZipMiddleware struct {
	handler lib.HttpHandler
}

func NewGZipMiddleware(handler lib.HttpHandler) GZipMiddleware {
	return GZipMiddleware{
		handler: handler,
	}
}

func (m GZipMiddleware) Setup() {
	m.handler.Engine.Use(gzip.Gzip(gzip.DefaultCompression))
}
