package main

import (
	"github.com/InsaneProgZ/user-service/src/adapters/input/controller"
	"github.com/InsaneProgZ/user-service/src/domain/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func main() {
	defer glog.Flush()
	router := gin.Default()

	userController := appConfig()

	// Define API routes
	v1 := router.Group("/v1")
	v1.POST("/users", userController.CreateUsers)
	
	router.Use(gin.Logger())
	// Run the server
	router.Run("localhost:8080")
}

func appConfig() *controller.UserController {
	defer glog.Flush()
	userPort := service.NewUserService()
	userController := controller.NewUserController(userPort)
	return userController
}
