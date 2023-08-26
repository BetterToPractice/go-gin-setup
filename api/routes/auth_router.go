package routes

import (
	"github.com/BetterToPractice/go-gin-setup/api/controllers"
	"github.com/BetterToPractice/go-gin-setup/lib"
)

type AuthRouter struct {
	handler        lib.HttpHandler
	authController controllers.AuthController
}

func NewAuthRouter(handler lib.HttpHandler, authController controllers.AuthController) AuthRouter {
	return AuthRouter{
		handler:        handler,
		authController: authController,
	}
}

func (r AuthRouter) Setup() {
	r.handler.Engine.POST("/login", r.authController.Login)
	r.handler.Engine.POST("/register", r.authController.Register)
}
