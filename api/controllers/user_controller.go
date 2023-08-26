package controllers

import "github.com/gin-gonic/gin"

type UserController struct {
}

func NewUserController() UserController {
	return UserController{}
}

func (c UserController) List(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "user list",
	})
}

func (c UserController) Detail(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "user detail",
	})
}
