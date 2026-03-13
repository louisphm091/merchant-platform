package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AdminClaims struct {
	AdminID string `json:"admin_id"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateAdminJWT(adminID string, email string, secret string) (string, error) {
	claims := AdminClaims{
		AdminID: adminID,
		Email:   email,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   adminID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
