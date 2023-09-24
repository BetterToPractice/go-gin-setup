package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MainController struct{}

func NewMainController() MainController {
	return MainController{}
}

func (c MainController) Index(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, "/swagger/index.html")
}
