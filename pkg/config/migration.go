package config

import (
	"errors"
	"os"
)

type Migration struct {
	Path     string
	Database string
}

func Migrations() (*Migration, error) {
	mg := Migration{
		Path:     os.Getenv("MIGRATION_PATH"),
		Database: os.Getenv("DATABASE_URL"),
	}
	if mg.Path == "" {
		return nil, errors.New("path required")
	}
	if mg.Database == "" {
		return nil, errors.New("database required")
	}
	return &mg, nil
}
