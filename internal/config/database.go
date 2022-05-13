package config

import (
	"database/sql"
	"majoo/pkg/config"
	"strings"

	_ "github.com/lib/pq"
)

func InitDatabase() *sql.DB {
	cfg := config.DatabasePGSQL()
	configs := []string{
		"host=" + cfg.Host,
		"user=" + cfg.User,
		"password=" + cfg.Password,
		"dbname=" + cfg.Database,
		"port=" + cfg.Port,
		"TimeZone=" + cfg.TimeZone,
		"sslmode=disable",
		"search_path=" + cfg.Schema,
	}
	dsn := strings.Join(configs, " ")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
