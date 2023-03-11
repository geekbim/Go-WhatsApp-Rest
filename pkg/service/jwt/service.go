package jwt

import (
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type jwtService struct {
	secretKey string
	issuer    string
}

// NewJWTService method is creates a new instance of JWTService
func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")

	if secretKey == "" {
		secretKey = "geekbim"
	}

	return secretKey
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	token = strings.Split(token, "Bearer ")[1]
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}

		return []byte(j.secretKey), nil
	})
}

func (j *jwtService) GetEmailByToken(token string) (string, error) {
	aToken, err := j.ValidateToken(token)
	if err != nil {
		return "", err
	}

	claims := aToken.Claims.(jwt.MapClaims)

	email := fmt.Sprintf("%v", claims["email"])

	return email, nil
}
