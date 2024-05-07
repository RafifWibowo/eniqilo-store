package main

import (
	"eniqilo_store/db"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.Default()

	db.Connect()

	r.Run(":8080")
}