package jwt

import (
	"fmt"
	"majoo/pkg/config"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")

	fields := []string{"JWT_SECRET"}

	for _, f := range fields {
		err := config.Required(f)
		if err != nil {
			panic(err.Error())
		}
	}

	return secretKey
}

func (j *jwtService) GenerateToken(UserID string) string {
	claims := &jwtCustomClaim{
		UserID,
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

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	if len(strings.Split(token, "Bearer ")) > 0 {
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
