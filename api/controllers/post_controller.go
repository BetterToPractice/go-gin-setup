package controllers

import (
	"github.com/BetterToPractice/go-gin-setup/api/services"
	"github.com/BetterToPractice/go-gin-setup/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostController struct {
	postService services.PostService
}

func NewPostController(postService services.PostService) PostController {
	return PostController{
		postService: postService,
	}
}

func (c PostController) List(ctx *gin.Context) {
	params := new(models.PostQueryParams)
	if err := ctx.Bind(params); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	qr, err := c.postService.Query(params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, qr)
}

func (c PostController) Detail(ctx *gin.Context) {
	qr, err := c.postService.Get(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, qr)
}
