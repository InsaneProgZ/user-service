package service

import (
	"fmt"

	"github.com/InsaneProgZ/user-service/src/application/domain"
)

type UserService struct{}

func (u *UserService) CreateUser(user domain.CreateUser) {
	fmt.Println("test")
}
