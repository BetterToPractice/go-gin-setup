package controllers

import (
	"github.com/BetterToPractice/go-gin-setup/api/dto"
	"github.com/BetterToPractice/go-gin-setup/api/services"
	"github.com/BetterToPractice/go-gin-setup/pkg/response"
	"github.com/gin-gonic/gin"
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
	register := new(dto.RegisterRequest)
	if err := ctx.ShouldBind(register); err != nil {
		response.BadRequest{Req: dto.RegisterRequest{}, Message: err}.JSON(ctx)
		return
	}

	resp, err := c.authService.Register(register)
	if err != nil {
		response.BadRequest{Message: err}.JSON(ctx)
		return
	}

	response.Response{Data: resp}.JSON(ctx)
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
	login := new(dto.LoginRequest)
	if err := ctx.ShouldBind(login); err != nil {
		response.BadRequest{Req: dto.LoginRequest{}, Message: err}.JSON(ctx)
		return
	}

	token, err := c.authService.Login(login)
	if err != nil {
		response.PolicyResponse{Message: err}.JSON(ctx)
		return
	}

	response.Response{Data: token}.JSON(ctx)
}
