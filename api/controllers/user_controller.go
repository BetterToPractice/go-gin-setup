package controllers

import (
	"github.com/BetterToPractice/go-gin-setup/api/services"
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
	users, _ := c.userService.Query()
	ctx.JSON(200, users)
}

func (c UserController) Detail(ctx *gin.Context) {
	user, err := c.userService.GetByUsername(ctx.Param("username"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
