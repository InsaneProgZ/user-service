package main

import (
	"planzin/user/src/adapters/input/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	userHandler := api.NewUserController()

	// Define API routes
	v1 := router.Group("/v1")
	{
		v1.POST("/users", userHandler.CreateUser)
	}
	router.Use(gin.Logger())
	// Run the server
	router.Run("localhost:8080")
}

func appConfig() *api.UserController {
	userHandler := api.NewUserController()
	return userHandler
}
