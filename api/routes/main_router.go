package routes

import (
	"github.com/BetterToPractice/go-gin-setup/api/controllers"
	"github.com/BetterToPractice/go-gin-setup/lib"
)

type MainRouter struct {
	swagger        lib.Swagger
	mainController controllers.MainController
}

func NewMainRouter(swagger lib.Swagger, mainController controllers.MainController) MainRouter {
	return MainRouter{
		swagger:        swagger,
		mainController: mainController,
	}
}

func (r MainRouter) Setup() {
	r.swagger.Setup()
}
