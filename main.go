package main

import (
	"chain-vote-api/middlewares"
	"chain-vote-api/repositories"
	"chain-vote-api/security"
	"chain-vote-api/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	r := gin.Default()

	// Load environment variables if .env file exists
	err := godotenv.Load(".vscode/.env")
	
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	repositories.ConnectDataBase()

	

	public := r.Group("/api")
	public.POST("/register", services.Register)
	public.POST("/login", security.Login)

	protected := r.Group("/api")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/me", services.CurrentUser)
	protected.PUT("/register-address", services.RegisterETHAddress)

	r.Run(":8080")
}
