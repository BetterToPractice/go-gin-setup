package controllers

import (
	"fmt"
	"github.com/BetterToPractice/go-gin-setup/api/services"
	"github.com/BetterToPractice/go-gin-setup/models"
	"github.com/BetterToPractice/go-gin-setup/pkg/response"
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

// List godoc
//
//	@Summary		List of posts
//	@Description	get list of posts
//	@Tags			post
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/posts [get]
func (c PostController) List(ctx *gin.Context) {
	params := new(models.PostQueryParams)

	if err := ctx.ShouldBindQuery(params); err != nil {
		fmt.Println(err.Error())
		response.Response{Code: http.StatusBadRequest, Message: err}.JSON(ctx)
		return
	}

	qr, err := c.postService.Query(params)
	if err != nil {
		response.Response{Code: http.StatusBadRequest, Message: err}.JSON(ctx)
		return
	}

	response.Response{Code: http.StatusOK, Data: qr}.JSON(ctx)
}

// Detail godoc
//
//	@Summary		Detail a post
//	@Description	get detail a post
//	@Param 			id path string true "post id"
//	@Tags			post
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/posts/{id} [get]
func (c PostController) Detail(ctx *gin.Context) {
	qr, err := c.postService.Get(ctx.Param("id"))
	if err != nil {
		response.Response{Code: http.StatusNotFound, Message: err}.JSON(ctx)
		return
	}

	response.Response{Code: http.StatusOK, Data: qr}.JSON(ctx)
}

// Destroy godoc
//
//	@Summary		Delete a post
//	@Description	delete a post
//	@Param 			id path string true "post id"
//	@Tags			post
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/posts/{id} [delete]
func (c PostController) Destroy(ctx *gin.Context) {
	err := c.postService.Delete(ctx.Param("id"))
	if err != nil {
		response.Response{Code: http.StatusBadRequest, Message: err}.JSON(ctx)
		return
	}
	response.Response{Code: http.StatusNoContent}.JSON(ctx)
}
