package main

import (
	"github.com/InsaneProgZ/user-service/src/adapters/input/api"
	"github.com/InsaneProgZ/user-service/src/application/service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	userHandler := appConfig()

	// Define API routes
	v1 := router.Group("/v1")
	{
		v1.POST("/users", userHandler.CreateUsers)
	}
	router.Use(gin.Logger())
	// Run the server
	router.Run("localhost:8080")
}

func appConfig() *api.UserController {
	userPort := &service.UserService{}
	userHandler := api.NewUserController(userPort)
	return userHandler
}
