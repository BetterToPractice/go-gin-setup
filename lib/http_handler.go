package lib

import "github.com/gin-gonic/gin"

type HttpHandler struct {
	Engine *gin.Engine
}

func NewHttpHandler() HttpHandler {
	engine := gin.Default()
	return HttpHandler{
		Engine: engine,
	}
}
