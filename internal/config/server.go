package config

import (
	"go_wa_rest/pkg/config"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Env     string
	Port    string
	TimeOut time.Duration
}

func convertInt(env string) int {
	v, _ := strconv.Atoi(os.Getenv(env))
	return v
}

func Server() ServerConfig {
	err := godotenv.Load("./.env")
	if err != nil {
		panic("Error loading .env file")
	}

	cfg := ServerConfig{
		Env:     os.Getenv("ENV"),
		Port:    os.Getenv("SERVER_PORT"),
		TimeOut: time.Duration(convertInt("APP_TIMEOUT")) * time.Second,
	}
	err = cfg.Validate()
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
