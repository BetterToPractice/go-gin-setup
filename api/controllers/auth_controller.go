package controllers

import (
	"github.com/BetterToPractice/go-gin-setup/api/services"
	"github.com/BetterToPractice/go-gin-setup/models/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	authService services.AuthService
	userService services.UserService
}

func NewAuthController(authService services.AuthService, userService services.UserService) AuthController {
	return AuthController{
		authService: authService,
		userService: userService,
	}
}

func (c AuthController) Register(ctx *gin.Context) {
	register := new(dto.Register)
	if err := ctx.ShouldBind(register); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	_, err := c.userService.Register(register.Username, register.Password, register.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, dto.RegisterResponse{
		Username: register.Username,
		Email:    register.Email,
	})
}

func (c AuthController) Login(ctx *gin.Context) {
	login := new(dto.Login)
	if err := ctx.ShouldBind(login); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := c.userService.Verify(login.Username, login.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := c.authService.GenerateToken(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, dto.JwtResponse{Access: token})
}
