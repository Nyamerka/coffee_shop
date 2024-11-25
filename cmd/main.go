package main

import (
	"github.com/Nyamerka/coffee_shop/db"
	"github.com/Nyamerka/coffee_shop/internal/auth"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

var jwtSecret []byte

func initJWT() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET is not set")
	}
	jwtSecret = []byte(secret)
}

func main() {
	db.InitDB()
	r := gin.Default()
	auth.InitCasbin()
	initJWT()
	auth.RegisterAuthRoutes(r)

	r.Run(":8080")
}
