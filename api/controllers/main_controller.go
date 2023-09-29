package controllers

import (
	"github.com/BetterToPractice/go-gin-setup/lib"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MainController struct {
	config lib.Config
}

func NewMainController(config lib.Config) MainController {
	return MainController{
		config: config,
	}
}

func (c MainController) Index(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, c.config.Swagger.DocUrl)
}
