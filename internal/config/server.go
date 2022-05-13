package config

import (
	"majoo/pkg/config"
	"os"
	"strconv"
	"time"
)

type ServerConfig struct {
	Port    string
	TimeOut time.Duration
	Debug   bool
}

func convertInt(env string) int {
	v, _ := strconv.Atoi(os.Getenv(env))
	return v
}

func Server() ServerConfig {
	cfg := ServerConfig{
		Port:    os.Getenv("SERVER_PORT"),
		TimeOut: time.Duration(convertInt("APP_TIMEOUT")) * time.Second,
		Debug:   config.ConvertBool("DEBUG"),
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
		"DB_TIMEZONE",
	}

	for _, f := range fields {
		err := config.Required(f)
		if err != nil {
			return err
		}
	}
	return nil
}
