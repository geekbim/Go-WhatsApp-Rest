package jwt

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	if len(strings.Split(token, "Bearer ")) <= 1 {
		return nil, fmt.Errorf("token not found")
	}
	token = strings.Split(token, "Bearer ")[1]
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}

		return []byte(j.secretKey), nil
	})
}
