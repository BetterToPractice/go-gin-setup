package middlewares

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewCorsMiddleware),
	fx.Provide(NewJWTMiddleware),
	fx.Provide(NewGZipMiddleware),
	fx.Provide(NewHelmetMiddleware),
	fx.Provide(NewMiddlewares),
)

type IMiddleware interface {
	Setup()
}

type Middlewares []IMiddleware

func NewMiddlewares(cors CorsMiddleware) Middlewares {
	return Middlewares{
		cors,
	}
}

func (m Middlewares) Setup() {
	for _, middleware := range m {
		middleware.Setup()
	}
}
