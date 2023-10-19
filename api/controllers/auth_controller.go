package controllers

import (
	"github.com/BetterToPractice/go-gin-setup/api/services"
	"github.com/BetterToPractice/go-gin-setup/models/dto"
	"github.com/BetterToPractice/go-gin-setup/pkg/response"
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
//	@Param 			data body dto.Register true "Post"
//	@Router			/register [post]
//	@Success		200  {object}  response.Response{data=dto.RegisterResponse}  "ok"
//	@Failure		400  {object}  response.Response{data=[]response.ValidationErrors}  "bad request"
func (c AuthController) Register(ctx *gin.Context) {
	register := new(dto.Register)
	if err := ctx.ShouldBind(register); err != nil {
		response.BadRequest{Req: dto.Register{}, Message: err}.JSON(ctx)
		return
	}

	_, err := c.authService.Register(register.Username, register.Password, register.Email)
	if err != nil {
		response.BadRequest{Message: err}.JSON(ctx)
		return
	}

	response.Response{
		Code: http.StatusCreated,
		Data: dto.RegisterResponse{
			Username: register.Username,
			Email:    register.Email,
		},
	}.JSON(ctx)
}

// Login godoc
//
//	@Summary		Login a User
//	@Description	Login a user's application
//	@Tags			auth
//	@Accept			application/json
//	@Produce		application/json
//	@Param 			data body dto.Login true "Post"
//	@Router			/login [post]
//	@Success		200  {object}  response.Response{data=dto.LoginResponse}  "ok"
//	@Failure		400  {object}  response.Response{data=[]response.ValidationErrors}  "bad request"
func (c AuthController) Login(ctx *gin.Context) {
	login := new(dto.Login)
	if err := ctx.ShouldBind(login); err != nil {
		response.BadRequest{Req: dto.Login{}, Message: err}.JSON(ctx)
		return
	}

	token, err := c.authService.Login(login)
	if err != nil {
		response.BadRequest{Message: err}.JSON(ctx)
		return
	}

	ctx.JSON(http.StatusOK, token)
}
