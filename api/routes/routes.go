package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewUserRouter),
	fx.Provide(NewRoutes),
)

type IRoute interface {
	Setup()
}

type Routes []IRoute

func NewRoutes(
	userRouter UserRouter,
) Routes {
	return Routes{
		userRouter,
	}
}

func (a Routes) Setup() {
	for _, route := range a {
		route.Setup()
	}
}
