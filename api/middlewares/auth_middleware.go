package middlewares

import (
	"github.com/BetterToPractice/go-gin-setup/api/services"
	"github.com/BetterToPractice/go-gin-setup/constants"
	"github.com/BetterToPractice/go-gin-setup/lib"
	"github.com/gin-gonic/gin"
	"strings"
)

type AuthMiddleware struct {
	config      lib.Config
	handler     lib.HttpHandler
	authService services.AuthService
}

func NewAuthMiddleware(config lib.Config, handler lib.HttpHandler, authService services.AuthService) AuthMiddleware {
	return AuthMiddleware{
		config:      config,
		authService: authService,
		handler:     handler,
	}
}

func (m AuthMiddleware) core() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			auth   = ctx.GetHeader("Authorization")
			prefix = "Bearer "
			token  string
		)

		if auth != "" && strings.HasPrefix(auth, prefix) {
			token = auth[len(prefix):]
		}

		claims, _ := m.authService.ParseToken(token)
		ctx.Set(constants.CurrentUser, claims)
		ctx.Next()
	}
}

func (m AuthMiddleware) Setup() {
	m.handler.Engine.Use(m.core())
}
