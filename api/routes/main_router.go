package routes

import (
	"github.com/BetterToPractice/go-gin-setup/api/controllers"
	"github.com/BetterToPractice/go-gin-setup/lib"
)

type MainRouter struct {
	swagger        lib.Swagger
	mainController controllers.MainController
	handler        lib.HttpHandler
}

func NewMainRouter(handler lib.HttpHandler, swagger lib.Swagger, mainController controllers.MainController) MainRouter {
	return MainRouter{
		swagger:        swagger,
		mainController: mainController,
		handler:        handler,
	}
}

func (r MainRouter) Setup() {
	r.swagger.Setup()
	r.handler.Engine.GET("/", r.mainController.Index)
}
