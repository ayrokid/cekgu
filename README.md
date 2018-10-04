# CEKGU - Restfull API GOLANG with JWT Authentication

Basically this is a starter kit for you to integrate Lumen with [JWT Authentication](https://jwt.io/).

## What's Added

- [Gin Web Framework](https://github.com/gin-gonic/gin).
- [JWT Auth](https://github.com/dgrijalva/jwt-go) for Go.
- [Go MySQL Driver](https://github.com/go-sql-driver/mysql) Go MySQL Driver is a MySQL driver for Go's (golang).
- [GoDotEnv](https://github.com/joho/godotenv) A Go (golang) port of the Ruby dotenv project (which loads env vars from a .env file).
- [Logrus](https://github.com/sirupsen/logrus) Logrus is a structured logger for Go (golang).
- [GORM](https://github.com/jinzhu/gorm) The fantastic ORM library for Golang, aims to be developer friendly.

## Quick Start

- Clone this repo or download it's release archive and extract it somewhere
- You may delete `.git` folder if you get this code via `git clone`
- Copy `.env.example` to `.env`
- Run `go run main.go`


## Deploy Production

1. create windows :
env GOOS=windows GOARCH=amd64 go build main.go

2. create linux:
env GOOS=linux GOARCH=amd64 go build main.go


