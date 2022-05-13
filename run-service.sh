#!/bin/sh
export $(cat .env | xargs)
go run cmd/main.go