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

// Register godoc
//
//	@Summary		Register a new User
//	@Description	register a new user
//	@Tags			auth
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/register [post]
//	@Success		200  {object}  response.Response{data=dto.RegisterResponse}  "ok"
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

// Login godoc
//
//	@Summary		Login a User
//	@Description	Login a user's application
//	@Tags			auth
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/login [post]
//	@Success		200  {object}  response.Response{data=dto.JwtResponse}  "ok"
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
