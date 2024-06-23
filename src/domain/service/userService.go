package service

import (
	"fmt"

	"github.com/InsaneProgZ/user-service/src/domain"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) CreateUser(user domain.CreateUser) {
	fmt.Println("test")
}
