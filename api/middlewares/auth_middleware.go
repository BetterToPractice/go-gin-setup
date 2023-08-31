package middlewares

import (
	"github.com/BetterToPractice/go-gin-setup/api/services"
	"github.com/BetterToPractice/go-gin-setup/lib"
	"github.com/gin-gonic/gin"
	"net/http"
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
	prefixes := m.config.Auth.IgnorePathPrefixes
	return func(ctx *gin.Context) {
		if isIgnorePath(ctx.Request.URL.String(), prefixes...) {
			ctx.Next()
			return
		}

		var (
			auth   = ctx.GetHeader("Authorization")
			prefix = "Bearer "
			token  string
		)

		if auth != "" && strings.HasPrefix(auth, prefix) {
			token = auth[len(prefix):]
		}

		claims, err := m.authService.ParseToken(token)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("currentUser", claims)
		ctx.Next()
	}
}

func (m AuthMiddleware) Setup() {
	m.handler.Engine.Use(m.core())
}
