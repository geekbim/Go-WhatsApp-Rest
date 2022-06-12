package config

import (
	"go-wa-rest/pkg/config"
	"os"
	"strconv"
	"time"
)

type ServerConfig struct {
	Port    string
	TimeOut time.Duration
}

func convertInt(env string) int {
	v, _ := strconv.Atoi(os.Getenv(env))
	return v
}

func Server() ServerConfig {
	cfg := ServerConfig{
		Port:    os.Getenv("SERVER_PORT"),
		TimeOut: time.Duration(convertInt("APP_TIMEOUT")) * time.Second,
	}
	err := cfg.Validate()
	if err != nil {
		panic(err)
	}
	return cfg
}

func (c *ServerConfig) Validate() error {
	fields := []string{
		"SERVER_PORT",
	}

	for _, f := range fields {
		err := config.Required(f)
		if err != nil {
			return err
		}
	}
	return nil
}
