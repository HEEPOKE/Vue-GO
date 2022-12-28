package main

import (
	"Server/Api/database"
	"Server/Api/routes"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	database.ConnectDB()
	routes.Router()
}
