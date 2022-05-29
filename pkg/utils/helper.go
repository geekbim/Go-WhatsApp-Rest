package utils

import (
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func StringToInt(value string) (conv int) {
	if value != "" {
		conv, _ = strconv.Atoi(value)
	}
	if value == "" {
		conv = 0
	}
	return
}
