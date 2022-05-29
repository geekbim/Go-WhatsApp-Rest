package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func (j *jwtService) GenerateToken(userId, role string) string {
	claims := &jwtCustomClaim{
		userId,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(j.secretKey))

	if err != nil {
		panic(err.Error())
	}

	return t
}
