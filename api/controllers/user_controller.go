package controllers

import (
	"github.com/BetterToPractice/go-gin-setup/api/policies"
	"github.com/BetterToPractice/go-gin-setup/api/services"
	"github.com/BetterToPractice/go-gin-setup/models"
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
func (c UserController) List(ctx *gin.Context) {
	params := new(models.UserQueryParams)
	if err := ctx.ShouldBindQuery(params); err != nil {
		response.Response{Code: http.StatusBadRequest, Message: err}.JSON(ctx)
		return
	}

	qr, err := c.userService.Query(params)
	if err != nil {
		response.Response{Code: http.StatusBadRequest, Message: err}.JSON(ctx)
		return
	}

	response.Response{Code: http.StatusOK, Data: qr}.JSON(ctx)
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
func (c UserController) Detail(ctx *gin.Context) {
	qr, err := c.userService.GetByUsername(ctx.Param("username"))
	if err != nil {
		response.Response{Code: http.StatusNotFound, Message: err}.JSON(ctx)
		return
	}

	response.Response{Code: http.StatusOK, Data: qr}.JSON(ctx)
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
func (c UserController) Destroy(ctx *gin.Context) {
	user, err := c.userService.GetByUsername(ctx.Param("username"))
	if err != nil {
		response.Response{Code: http.StatusNotFound, Message: err}.JSON(ctx)
		return
	}

	loggedInUser, _ := c.authService.Authenticate(ctx)
	if isCan, err := c.userPolicy.CanDelete(loggedInUser, user); !isCan {
		response.Response{Code: http.StatusUnauthorized, Message: err}.JSON(ctx)
		return
	}

	if err := c.userService.Delete(user); err != nil {
		response.Response{Code: http.StatusBadRequest, Message: err}.JSON(ctx)
		return
	}

	response.Response{Code: http.StatusNoContent}.JSON(ctx)
	return
}
