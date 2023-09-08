package controllers

import (
	"github.com/BetterToPractice/go-gin-setup/api/services"
	"github.com/BetterToPractice/go-gin-setup/models"
	"github.com/BetterToPractice/go-gin-setup/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{
		userService: userService,
	}
}

// List godoc
//
//	@Summary		List users
//	@Description	get list users
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Router			/users/ [get]
//	@Success		200	{array}	dto.User
func (c UserController) List(ctx *gin.Context) {
	params := new(models.UserQueryParams)
	if err := ctx.Bind(params); err != nil {
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

func (c UserController) Detail(ctx *gin.Context) {
	qr, err := c.userService.GetByUsername(ctx.Param("username"))
	if err != nil {
		response.Response{Code: http.StatusNotFound, Message: err}.JSON(ctx)
		return
	}

	response.Response{Code: http.StatusOK, Data: qr}.JSON(ctx)
}
