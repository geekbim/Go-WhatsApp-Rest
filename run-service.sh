#!/bin/sh
export $(cat .env | xargs)
go run main.go