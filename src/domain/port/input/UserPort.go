package input

import "github.com/InsaneProgZ/user-service/src/domain/model"

type UserPort interface {
	CreateUser(user model.User) error
	FindUser(username string) (model.User, error)
}
