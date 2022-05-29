package jwt

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func (j *jwtService) GetUserIdAndRole(token string) (string, string, error) {
	aToken, err := j.ValidateToken(token)
	if err != nil {
		return "", "", err
	}

	claims := aToken.Claims.(jwt.MapClaims)
	userId := fmt.Sprintf("%v", claims["userId"])
	role := fmt.Sprintf("%v", claims["role"])

	return userId, role, nil
}
