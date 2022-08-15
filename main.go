package main

import (
	"chain-vote-api/middlewares"
	"chain-vote-api/repositories"
	"chain-vote-api/security"
	"chain-vote-api/services"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	repositories.ConnectDataBase()

	public := r.Group("/api")
	public.POST("/register", services.Register)
	public.POST("/login", security.Login)

	protected := r.Group("/api")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/me", services.CurrentUser)

	r.Run(":8080")
}
