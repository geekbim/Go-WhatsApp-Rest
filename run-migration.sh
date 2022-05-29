#!/bin/sh
export DATABASE_URL="postgres://postgres:hwhwhwlol@localhost:5432/gokomodo_db?sslmode=disable" \
export MIGRATION_PATH="/home/geekbim/go/src/gokomodo/migration/pgsql" \
# go run migration/main.go migration:status
# go run migration/main.go migration:down
go run migration/main.go migration:up
# go run migration/main.go migration:create create_table_orders --table=orders