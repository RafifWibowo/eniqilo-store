package main

import (
	"eniqilo_store/db"
	"eniqilo_store/routes"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db.Connect()
	r := routes.Init()
	r.Run(":8080")
}