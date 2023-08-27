package middlewares

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewCorsMiddleware),
	fx.Provide(NewAuthMiddleware),
	fx.Provide(NewGZipMiddleware),
	fx.Provide(NewHelmetMiddleware),
	fx.Provide(NewMiddlewares),
)

type IMiddleware interface {
	Setup()
}

type Middlewares []IMiddleware

func NewMiddlewares(cors CorsMiddleware, gzip GZipMiddleware, helmet HelmetMiddleware, auth AuthMiddleware) Middlewares {
	return Middlewares{
		auth,
		cors,
		gzip,
		helmet,
	}
}

func (m Middlewares) Setup() {
	for _, middleware := range m {
		middleware.Setup()
	}
}

func isIgnorePath(path string, prefixes ...string) bool {
	pathLen := len(path)

	for _, p := range prefixes {
		if pl := len(p); pathLen >= pl && path[:pl] == p {
			return true
		}
	}

	return false
}
