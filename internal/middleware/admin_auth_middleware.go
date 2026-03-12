package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/louisphm091/merchant-platform/internal/config"
	"github.com/louisphm091/merchant-platform/pkg/response"
)

func AdminAuthMiddleware(cfg *config.Config) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, "Unauthorized", "missing authorization header")
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			response.Error(c, http.StatusUnauthorized, "Unauthorized", "invalid authorization header")
			c.Abort()
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			response.Error(c, http.StatusUnauthorized, "Unauthorized", "invalid or expired token")
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			response.Error(c, http.StatusUnauthorized, "Unauthorized", "invalid token claims")
			c.Abort()
			return
		}

		role, ok := claims["role"].(string)

		if !ok || role != "admin" {
			response.Error(c, http.StatusForbidden, "Forbidden", "admin access required")
			c.Abort()
			return
		}

		c.Set("admin_id", claims["admin_id"])
		c.Set("admin_email", claims["email"])
		c.Next()
	}
}
