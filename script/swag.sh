#!/usr/bin/env bash
# http://localhost:8080/swagger/index.html
swag fmt
swag init -g ../internal/controller/http/v1/route.go -parseInternal
