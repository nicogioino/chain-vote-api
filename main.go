package main

import (
	"chain-vote-api/middlewares"
	"chain-vote-api/repositories"
	"chain-vote-api/router"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	r := gin.Default()
	r.Use(middlewares.AllowAll())
	router.InitializeRouter(r)

	// Load environment variables if .env file exists
	err := godotenv.Load(".vscode/.env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	repositories.ConnectDataBase()

	log.Fatalln(r.Run(":8080"))
}
