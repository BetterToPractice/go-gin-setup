package controllers

import (
	"github.com/BetterToPractice/go-gin-setup/api/services"
	"github.com/BetterToPractice/go-gin-setup/models"
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
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	qr, err := c.userService.Query(params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(200, qr)
}

func (c UserController) Detail(ctx *gin.Context) {
	user, err := c.userService.GetByUsername(ctx.Param("username"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, user)
}
