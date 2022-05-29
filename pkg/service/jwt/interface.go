package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(userId, role string) string
	ValidateToken(token string) (*jwt.Token, error)
	GetUserIdAndRole(token string) (string, string, error)
}
