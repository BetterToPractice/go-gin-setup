package controllers

import (
	"github.com/BetterToPractice/go-gin-setup/api/dto"
	"github.com/BetterToPractice/go-gin-setup/api/policies"
	"github.com/BetterToPractice/go-gin-setup/api/services"
	"github.com/BetterToPractice/go-gin-setup/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userService services.UserService
	authService services.AuthService
	userPolicy  policies.UserPolicy
}

func NewUserController(userService services.UserService, authService services.AuthService, userPolicy policies.UserPolicy) UserController {
	return UserController{
		userService: userService,
		authService: authService,
		userPolicy:  userPolicy,
	}
}

// List godoc
//
//	@Summary		List several users
//	@Description	get list several users
//	@Tags			user
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/users [get]
//	@Success		200  {object}  response.Response{data=dto.UserPaginationResponse}  "ok"
//	@Failure		400  {object}  response.Response{data=[]response.ValidationErrors}  "bad request"
func (c UserController) List(ctx *gin.Context) {
	params := new(dto.UserQueryParam)
	if err := ctx.ShouldBindQuery(params); err != nil {
		response.BadRequest{Req: dto.UserQueryParam{}, Message: err}.JSON(ctx)
		return
	}

	qr, err := c.userService.Query(params)
	if err != nil {
		response.BadRequest{Message: err}.JSON(ctx)
		return
	}

	response.Response{Data: qr}.JSON(ctx)
}

// Detail godoc
//
//	@Summary		Get a User
//	@Description	get a user by username
//	@Tags			user
//	@Accept			application/json
//	@Produce		application/json
//	@Param			username  path  string  true  "Username"
//	@Router			/users/{username} [get]
//	@Success		200  {object}  response.Response{data=dto.UserResponse}  "ok"
//	@Failure		404  {object}  response.Response  "not found"
func (c UserController) Detail(ctx *gin.Context) {
	user, err := c.userService.GetByUsername(ctx.Param("username"))
	if err != nil {
		response.Response{Code: http.StatusNotFound, Message: err}.JSON(ctx)
		return
	}

	resp := dto.UserResponse{}
	resp.Serializer(user)

	response.Response{Code: http.StatusOK, Data: resp}.JSON(ctx)
}

// Destroy godoc
//
//	@Summary		Delete a User
//	@Description	perform delete a user by username
//	@Tags			user
//	@Accept			application/json
//	@Produce		application/json
//	@Security 		BearerAuth
//	@Param			username  path  string  true  "Username"
//	@Router			/users/{username} [delete]
//	@Success		204  {object}  nil  "no content"
//	@Failure		404  {object}  response.Response  "not found"
//	@Failure		401  {object}  response.Response  "unauthorized"
//	@Failure		403  {object}  response.Response  "forbidden"
func (c UserController) Destroy(ctx *gin.Context) {
	user, err := c.userService.GetByUsername(ctx.Param("username"))
	if err != nil {
		response.NotFound{Message: err}.JSON(ctx)
		return
	}

	loggedInUser, _ := c.authService.Authenticate(ctx)
	if isCan, err := c.userPolicy.CanDelete(loggedInUser, user); !isCan {
		response.PolicyResponse{Message: err}.JSON(ctx)
		return
	}

	if err := c.userService.Delete(user); err != nil {
		response.BadRequest{Message: err}.JSON(ctx)
		return
	}

	response.Response{Code: http.StatusNoContent}.JSON(ctx)
	return
}
