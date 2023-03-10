package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

// JWTService is a contract of what jwtService can do
type JWTService interface {
	ValidateToken(token string) (*jwt.Token, error)
	GetEmailByToken(token string) (string, error)
}
