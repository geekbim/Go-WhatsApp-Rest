package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type DatabaseConfig struct {
	User              string
	Password          string
	Database          string
	Host              string
	Port              string
	TimeZone          string
	MaxConnectionIdle int
	MaxConnectionOpen int
	Schema            string
	TimeOut           time.Duration
}

func ConvertInt(env string) int {
	v, _ := strconv.Atoi(os.Getenv(env))
	return v
}

func ConvertBool(env string) bool {
	v, _ := strconv.ParseBool(os.Getenv(env))
	return v
}

func DatabasePGSQL() DatabaseConfig {
	schema := "public"
	cfgSchema := os.Getenv("DB_SCHEMA")
	if cfgSchema != "" {
		schema = cfgSchema
	}

	return DatabaseConfig{
		User:              os.Getenv("DB_USERNAME"),
		Password:          os.Getenv("DB_PASSWORD"),
		Database:          os.Getenv("DB_NAME"),
		Host:              os.Getenv("DB_HOST"),
		Port:              os.Getenv("DB_PORT"),
		TimeZone:          os.Getenv("DB_TIMEZONE"),
		MaxConnectionIdle: ConvertInt("DB_MAX_CON_IDLE"),
		MaxConnectionOpen: ConvertInt("DB_MAX_CON_OPEN"),
		Schema:            schema,
		TimeOut:           time.Duration(ConvertInt("APP_TIMEOUT")) * time.Second,
	}
}

func Required(key string) error {
	if os.Getenv(key) == "" {
		errorMsg := fmt.Sprintf("config %s is required", key)
		return errors.New(errorMsg)
	}
	return nil
}
