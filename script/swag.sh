#!/usr/bin/env bash
# http://localhost:8080/swagger/index.html
swag fmt
swag init -g ../cmd/aio/aio.go -parseInternal
