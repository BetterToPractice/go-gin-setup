package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct{}

type Config struct {
	Engine *gin.Engine
}

func NewHandler(c *Config) {
	h := &Handler{}
	c.Engine.GET("/users", h.UserList)
	c.Engine.GET("/users/:id", h.UserDetail)
}

func (h Handler) UserList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"results": "users",
	})
}

func (h Handler) UserDetail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "user detail",
	})
}
