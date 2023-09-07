package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewMainRouter),
	fx.Provide(NewUserRouter),
	fx.Provide(NewAuthRouter),
	fx.Provide(NewPostRouter),
	fx.Provide(NewRoutes),
)

type IRoute interface {
	Setup()
}

type Routes []IRoute

func NewRoutes(
	mainRouter MainRouter,
	userRouter UserRouter,
	authRouter AuthRouter,
	postRouter PostRouter,
) Routes {
	return Routes{
		mainRouter,
		userRouter,
		authRouter,
		postRouter,
	}
}

func (a Routes) Setup() {
	for _, route := range a {
		route.Setup()
	}
}
