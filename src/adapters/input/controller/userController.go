package controller

import (
	"net/http"
	"time"

	"github.com/InsaneProgZ/user-service/src/adapters/input/controller/request"
	"github.com/InsaneProgZ/user-service/src/domain/port/input"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

type UserController struct {
	userPort input.UserPort
}

func NewUserController(userPort input.UserPort) *UserController {
	return &UserController{userPort: userPort}
}

func (uc *UserController) CreateUsers(c *gin.Context) {
	var req request.CreateUserRequest

	glog.Info("start" + time.Now().GoString())

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	glog.Info("end" + time.Now().GoString())

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
