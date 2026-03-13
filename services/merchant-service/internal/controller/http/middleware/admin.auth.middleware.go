package middleware

import (
	"merchant-platform/merchant-service/internal/infrastructure/config"
	"merchant-platform/merchant-service/pkg/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AdminAuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, "unauthorized",
				response.NewError("MISSING_AUTH_HEADER", "authorization header is required"))
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			response.Error(c, http.StatusUnauthorized, "unauthorized",
				response.NewError("INVALID_AUTH_HEADER", "authorization header must be Bearer <token>"))
			c.Abort()
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			response.Error(c, http.StatusUnauthorized, "unauthorized",
				response.NewError("INVALID_TOKEN", "token is invalid or expired"))
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			response.Error(c, http.StatusUnauthorized, "unauthorized",
				response.NewError("INVALID_CLAIMS", "token claims are invalid"))
			c.Abort()
			return
		}

		role, ok := claims["role"].(string)
		if !ok || role != "admin" {
			response.Error(c, http.StatusForbidden, "forbidden",
				response.NewError("ADMIN_ACCESS_REQUIRED", "admin role is required"))
			c.Abort()
			return
		}

		c.Set("admin_id", claims["admin_id"])
		c.Set("admin_email", claims["email"])
		c.Next()
	}
}
