package controllers

import (
	"github.com/BetterToPractice/go-gin-setup/api/dto"
	"github.com/BetterToPractice/go-gin-setup/api/policies"
	"github.com/BetterToPractice/go-gin-setup/api/services"
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
//	@Success		200  {object}  response.Response{data=dto.PostPaginationResponse}  "ok"
//	@Failure		400  {object}  response.Response{data=[]response.ValidationErrors}  "bad request"
func (c PostController) List(ctx *gin.Context) {
	params := new(dto.PostQueryParam)
	if err := ctx.ShouldBindQuery(params); err != nil {
		response.BadRequest{Req: dto.PostQueryParam{}, Message: err}.JSON(ctx)
		return
	}

	qr, err := c.postService.Query(params)
	if err != nil {
		response.BadRequest{Message: err}.JSON(ctx)
		return
	}

	response.Response{Data: qr}.JSON(ctx)
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
//	@Success		200  {object}  response.Response{data=dto.PostResponse}  "ok"
//	@Failure		404  {object}  response.Response{}  "not found"
func (c PostController) Detail(ctx *gin.Context) {
	post, err := c.postService.Get(ctx.Param("id"))
	if err != nil {
		response.NotFound{Message: err}.JSON(ctx)
		return
	}

	user, _ := c.authService.Authenticate(ctx)
	if can, err := c.postPolicy.CanViewDetail(user, post); !can {
		response.PolicyResponse{Message: err}.JSON(ctx)
		return
	}

	resp := dto.PostResponse{}
	resp.Serializer(post)
	response.Response{Code: http.StatusOK, Data: resp}.JSON(ctx)
}

// Create godoc
//
//	@Summary		Create a post
//	@Description	Create a post
//	@Tags			post
//	@Accept			application/json
//	@Produce		application/json
//	@Security 		BearerAuth
//	@Param 			data body dto.PostRequest true "Post"
//	@Router			/posts [post]
//	@Success		201  {object}  response.Response{data=dto.PostResponse}  "created"
//	@Failure		400  {object}  response.Response{data=[]response.ValidationErrors}  "bad request"
//	@Failure		404  {object}  response.Response  "not found"
//	@Failure		401  {object}  response.Response  "unauthorized"
//	@Failure		403  {object}  response.Response  "forbidden"
func (c PostController) Create(ctx *gin.Context) {
	user, _ := c.authService.Authenticate(ctx)
	if can, err := c.postPolicy.CanCreate(user); !can {
		response.PolicyResponse{Message: err}.JSON(ctx)
		return
	}

	params := new(dto.PostRequest)
	if err := ctx.ShouldBind(params); err != nil {
		response.BadRequest{Req: dto.PostRequest{}, Message: err}.JSON(ctx)
		return
	}

	resp, err := c.postService.Create(params, user)
	if err != nil {
		response.BadRequest{Message: err}.JSON(ctx)
		return
	}

	response.Response{Code: http.StatusCreated, Data: resp}.JSON(ctx)
}

// Update godoc
//
//	@Summary		Update a post
//	@Description	update a post
//	@Param 			id path string true "post id"
//	@Tags			post
//	@Accept			application/json
//	@Produce		application/json
//	@Security 		BearerAuth
//	@Param 			data body dto.PostUpdateRequest true "Post"
//	@Router			/posts/{id} [patch]
//	@Success		200  {object}  response.Response{data=dto.PostResponse}  "ok"
//	@Failure		400  {object}  response.Response{data=[]response.ValidationErrors}  "bad request"
//	@Failure		404  {object}  response.Response  "not found"
//	@Failure		401  {object}  response.Response  "unauthorized"
//	@Failure		403  {object}  response.Response  "forbidden"
func (c PostController) Update(ctx *gin.Context) {
	post, err := c.postService.Get(ctx.Param("id"))
	if err != nil {
		response.NotFound{Message: err}.JSON(ctx)
		return
	}

	user, _ := c.authService.Authenticate(ctx)
	if can, err := c.postPolicy.CanUpdate(user, post); !can {
		response.PolicyResponse{Message: err}.JSON(ctx)
		return
	}

	params := new(dto.PostUpdateRequest)
	if err := ctx.Bind(params); err != nil {
		response.BadRequest{Req: dto.PostUpdateRequest{}, Message: err}.JSON(ctx)
		return
	}

	resp, err := c.postService.Update(post, params)
	if err != nil {
		response.BadRequest{Message: err}.JSON(ctx)
		return
	}

	response.Response{Data: resp}.JSON(ctx)
}

// Destroy godoc
//
//	@Summary		Delete a post
//	@Description	delete a post
//	@Param 			id path string true "post id"
//	@Tags			post
//	@Accept			application/json
//	@Produce		application/json
//	@Security 		BearerAuth
//	@Router			/posts/{id} [delete]
//	@Success		204  {object}  nil  "no content"
//	@Failure		404  {object}  response.Response  "not found"
//	@Failure		401  {object}  response.Response  "unauthorized"
//	@Failure		403  {object}  response.Response  "forbidden"
func (c PostController) Destroy(ctx *gin.Context) {
	post, err := c.postService.Get(ctx.Param("id"))
	if err != nil {
		response.NotFound{Message: err}.JSON(ctx)
		return
	}

	user, _ := c.authService.Authenticate(ctx)
	if can, err := c.postPolicy.CanDelete(user, post); !can {
		response.PolicyResponse{Message: err}.JSON(ctx)
		return
	}

	if err := c.postService.Delete(post); err != nil {
		response.BadRequest{Message: err}.JSON(ctx)
		return
	}

	response.Response{Code: http.StatusNoContent}.JSON(ctx)
}
