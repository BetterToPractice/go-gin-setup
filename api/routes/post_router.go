package routes

import (
	"github.com/BetterToPractice/go-gin-setup/api/controllers"
	"github.com/BetterToPractice/go-gin-setup/lib"
)

type PostRouter struct {
	postController controllers.PostController
	handler        lib.HttpHandler
}

func NewPostRouter(handler lib.HttpHandler, postController controllers.PostController) PostRouter {
	return PostRouter{
		postController: postController,
		handler:        handler,
	}
}

func (r PostRouter) Setup() {
	r.handler.Engine.GET("/posts", r.postController.List)
	r.handler.Engine.GET("/posts/:id", r.postController.Detail)
}