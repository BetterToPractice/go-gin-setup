package controllers

import (
	"github.com/BetterToPractice/go-gin-setup/api/policies"
	"github.com/BetterToPractice/go-gin-setup/api/services"
	"github.com/BetterToPractice/go-gin-setup/models"
	"github.com/BetterToPractice/go-gin-setup/models/dto"
	"github.com/BetterToPractice/go-gin-setup/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostController struct {
	postService services.PostService
	userService services.UserService
	authService services.AuthService
	postPolicy  policies.PostPolicy
}

func NewPostController(postService services.PostService, userService services.UserService, authService services.AuthService, postPolicy policies.PostPolicy) PostController {
	return PostController{
		postService: postService,
		userService: userService,
		authService: authService,
		postPolicy:  postPolicy,
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

	user, err := c.authService.Authenticate(ctx)
	if canView, err := c.postPolicy.CanViewDetail(user, qr); !canView {
		response.Response{Code: http.StatusUnauthorized, Message: err}.JSON(ctx)
		return
	}

	response.Response{Code: http.StatusOK, Data: qr}.JSON(ctx)
}

// Create godoc
//
//	@Summary		Create a post
//	@Description	Create a post
//	@Tags			post
//	@Accept			application/json
//	@Produce		application/json
//	@Param 			data body dto.PostRequest true "Post"
//	@Router			/posts [post]
func (c PostController) Create(ctx *gin.Context) {
	user, _ := c.authService.Authenticate(ctx)
	if isCan, err := c.postPolicy.CanCreate(user); !isCan {
		response.Response{Code: http.StatusUnauthorized, Message: err}.JSON(ctx)
		return
	}

	params := new(dto.PostRequest)
	if err := ctx.ShouldBind(params); err != nil {
		response.Response{Code: http.StatusBadRequest, Message: err}.JSON(ctx)
		return
	}

	postResponse, err := c.postService.Create(params, user)
	if err != nil {
		response.Response{Code: http.StatusBadRequest, Message: err}.JSON(ctx)
		return
	}

	response.Response{Code: http.StatusOK, Data: postResponse}.JSON(ctx)
}

// Update godoc
//
//	@Summary		Update a post
//	@Description	update a post
//	@Param 			id path string true "post id"
//	@Tags			post
//	@Accept			application/json
//	@Produce		application/json
//	@Param 			data body dto.PostUpdateRequest true "Post"
//	@Router			/posts/{id} [patch]
//	@Success		200  {object}  response.Response{data=dto.PostResponse}  "ok"
func (c PostController) Update(ctx *gin.Context) {
	post, err := c.postService.Get(ctx.Param("id"))
	if err != nil {
		response.Response{Code: http.StatusNotFound, Message: err}.JSON(ctx)
		return
	}

	user, _ := c.authService.Authenticate(ctx)
	if isCan, err := c.postPolicy.CanUpdate(user, post); !isCan {
		response.Response{Code: http.StatusUnauthorized, Message: err}.JSON(ctx)
		return
	}

	params := new(dto.PostUpdateRequest)
	if err := ctx.Bind(params); err != nil {
		response.Response{Code: http.StatusBadRequest, Message: err}.JSON(ctx)
		return
	}

	postResponse, err := c.postService.Update(post, params)
	if err != nil {
		response.Response{Code: http.StatusBadRequest, Message: err}.JSON(ctx)
		return
	}

	response.Response{Code: http.StatusOK, Data: postResponse}.JSON(ctx)
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
	post, err := c.postService.Get(ctx.Param("id"))
	if err != nil {
		response.Response{Code: http.StatusNotFound, Message: err}.JSON(ctx)
		return
	}

	user, _ := c.authService.Authenticate(ctx)
	if isCan, err := c.postPolicy.CanDelete(user, post); !isCan {
		response.Response{Code: http.StatusUnauthorized, Message: err}.JSON(ctx)
		return
	}

	if err := c.postService.Delete(post); err != nil {
		response.Response{Code: http.StatusBadRequest, Message: err}.JSON(ctx)
		return
	}

	response.Response{Code: http.StatusNoContent}.JSON(ctx)
}
