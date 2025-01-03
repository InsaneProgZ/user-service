package controller

import (
	"net/http"
	"time"

	"github.com/InsaneProgZ/user-service/src/adapters/input/controller/request"
	"github.com/InsaneProgZ/user-service/src/domain/model"
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

	if err := uc.userPort.CreateUser(model.User(req)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	glog.Info("end" + time.Now().GoString())

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func (uc *UserController) FindUser(c *gin.Context) {
	glog.Info("start" + time.Now().GoString())

	username := c.Param("username")

	user, err := uc.userPort.FindUser(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	glog.Info("end" + time.Now().GoString())

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) ValidateUser(c *gin.Context) {
	var request request.ValidateUserRequest
	glog.Info("start ValidateUser" + time.Now().GoString())

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isValid, err := uc.userPort.ValidateUser(model.UserValidate(request))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if isValid {
		c.JSON(http.StatusOK, "Valid Passwords")
	} else {
		c.JSON(http.StatusOK, "Invalid Password")
	}

	glog.Info("end ValidateUser" + time.Now().GoString())

}
