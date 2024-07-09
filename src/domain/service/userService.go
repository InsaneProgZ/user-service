package service

import (
	"github.com/InsaneProgZ/user-service/src/domain/model"
	"github.com/InsaneProgZ/user-service/src/domain/port/output"
)

type UserService struct {
	UserRepositoryPort output.UserRepositoryPort
}

func NewUserService(userRepositoryPort output.UserRepositoryPort) *UserService {
	return &UserService{UserRepositoryPort: userRepositoryPort}
}

func (u *UserService) CreateUser(user model.User) error {
	error := u.UserRepositoryPort.Save(user)
	if error != nil {
		return error
	}
	return nil
}

func (u *UserService) FindUser(username string) (model.User, error) {
	user := u.UserRepositoryPort.Find(username)
	if user == nil {
		return model.User{}, &model.UserError{Message: "User not found."}
	}
	return *user, nil
}
