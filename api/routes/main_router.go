package routes

import (
	"github.com/BetterToPractice/go-gin-setup/api/controllers"
	"github.com/BetterToPractice/go-gin-setup/lib"
)

type MainRouter struct {
	mainController controllers.MainController
	handler        lib.HttpHandler
}

func NewMainRouter(handler lib.HttpHandler, mainController controllers.MainController) MainRouter {
	return MainRouter{
		mainController: mainController,
		handler:        handler,
	}
}

func (r MainRouter) Setup() {
	r.handler.Engine.GET("/", r.mainController.Index)
}
