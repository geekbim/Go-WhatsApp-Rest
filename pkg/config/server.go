package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ConvertInt(env string) int {
	v, _ := strconv.Atoi(os.Getenv(env))
	return v
}

func Required(key string) error {
	if os.Getenv(key) == "" {
		errorMsg := fmt.Sprintf("config %s is required", key)
		return errors.New(errorMsg)
	}
	return nil
}
