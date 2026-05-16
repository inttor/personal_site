package middleware

import (
	"PersonalSite/backend/config"
	jwtpkg "PersonalSite/backend/pkg/jwt"
	"PersonalSite/backend/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, 401, "Missing authorization header")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Error(c, 401, "Invalid authorization format")
			c.Abort()
			return
		}

		claims, err := jwtpkg.ParseToken(parts[1], cfg.JWTSecret)
		if err != nil {
			response.Error(c, 401, "Invalid or expired token")
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
