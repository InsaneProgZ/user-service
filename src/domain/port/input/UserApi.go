package input

import (
	"github.com/gin-gonic/gin"
)

type UserApi interface {
	CreateUsers(c *gin.Context)
	FindUser(c *gin.Context)
	ValidateUser(c *gin.Context)
}
