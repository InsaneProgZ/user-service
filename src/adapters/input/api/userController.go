package api

import (
	"net/http"

	"github.com/InsaneProgZ/user-service/src/adapters/input/api/request"
	"github.com/InsaneProgZ/user-service/src/port/input"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserPort input.UserPort
}

func NewUserController(userPort input.UserPort) *UserController {
	return &UserController{UserPort: userPort}
}

func (uc *UserController) CreateUsers(c *gin.Context) {
	var req request.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
